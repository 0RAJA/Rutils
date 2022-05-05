package RCache

import (
	"fmt"
	"github.com/0RAJA/Rutils/RCache/RErr"
	"github.com/0RAJA/Rutils/RSync/singleflight"
	"log"
	"net/http"
	"sync"
)

//Cache的表层

// Getter 自行实现调用数据库函数
type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

/*
一个 Group 可以认为是一个缓存的命名空间，每个 Group 拥有一个唯一的名称 name。比如可以创建三个 Group，缓存学生的成绩命名为 scores，缓存学生信息的命名为 info，缓存学生课程的命名为 courses。
第二个属性是 getter Getter，即缓存未命中时获取源数据的回调(callback)。
第三个属性是 mainCache cache，即一开始实现的并发缓存。
构建函数 NewGroup 用来实例化 Group，并且将 groups 存储在全局变量 groups 中。
GetGroup 用来特定名称的 Group，这里使用了只读锁 RLock()，因为不涉及任何冲突变量的写操作。
*/

// Group 分布式缓存组
type Group struct {
	name      string              //命名
	getter    Getter              //即缓存未命中时获取源数据的回调(callback)
	mainCache syncCache           //缓存
	peers     PeerPicker          //其他节点
	loader    *singleflight.Group //保证一个请求同一时刻只会执行一次(缓存击破)
}

var (
	mu     sync.RWMutex //保证groups并发安全
	groups = make(map[string]*Group)
)

// NewGroup 缓存,name:命名空间,cacheBytes:最大缓存字节数,getter:回调函数
func NewGroup(name string, cacheBytes int, getter Getter) *Group {
	if getter == nil {
		panic(RErr.GetterIsNil)
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: syncCache{cacheBytes: cacheBytes},
		loader:    singleflight.NewGroup(),
	}
	groups[name] = g
	return g
}

// GetGroup 获取指定name的缓存组
func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()
	return groups[name]
}

// Get 获取缓存
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf(RErr.KeyIsNil)
	}
	//从缓存中获取
	if v, ok := g.mainCache.get(key); ok {
		return v, nil
	}
	//从指定源获取
	return g.load(key)
}

func (g *Group) Set(key string, value ByteView) {
	g.mainCache.add(key, value)
}

/*
修改 load 方法，使用 PickPeer() 方法选择节点，
	若非本机节点，则调用 getFromPeer() 从远程获取。
	若是本机节点或失败，则回退到 getLocally()。
*/

//从指定地点指定源获取值
func (g *Group) load(key string) (value ByteView, err error) {
	//防止缓存穿透
	view, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				if value, err = g.getFromPeer(peer, key); err != nil {
					return value, nil
				}
				log.Println("[RCache] Failed to get from peer", err)
			}
		}
		return g.getLocally(key)
	})
	return view.(ByteView), err
}

//从本地指定源获取值
func (g *Group) getLocally(key string) (value ByteView, err error) {
	var bytes []byte
	if bytes, err = g.getter.Get(key); err == nil {
		value = ByteView{cloneBytes(bytes)}
		g.populateCache(key, value)
	}
	return value, err
}

//getFromPeer 方法，使用实现了 PeerGetter 接口的 httpGetter 从访问远程节点，获取缓存值。
func (g *Group) getFromPeer(peer PeerGetter, key string) (value ByteView, err error) {
	var bytes []byte
	if bytes, err = peer.Get(g.name, key); err == nil {
		value = ByteView{cloneBytes(bytes)}
	}
	return
}

//将缓存保存到cache
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}

//RegisterPeers 方法，将 实现了 PeerPicker 接口的 HTTPPool 注入到 Group 中。
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic(RErr.RegisterPeersErr)
	}
	g.peers = peers
}

// StartServer StartAPIServer 启动辅助服务
func (g *Group) StartServer(apiAddr, prefix, key string) {
	http.Handle(prefix, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get(key)
			view, err := g.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write(view.ByteSlice())
		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr, nil))
}

// StartCacheServer 开启分布式缓存 addr 当前主机ip及端口,address 所有主机ip及端口
func (g *Group) StartCacheServer(addr string, addressMap map[int]string) {
	//初始化
	pool := NewHTTPPool(addr)
	//加载其他节点
	address := make([]string, len(addressMap))
	for _, v := range addressMap {
		address = append(address, v)
	}
	pool.Set(address...)
	//注入group
	g.RegisterPeers(pool)
	log.Println("RCache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, pool))
}
