缓存流程图:
```text
1. 本地:
                            是
接收 key --> 检查是否被缓存 -----> 返回缓存值 ⑴
                |  否                          是
                |-----> 是否应当从远程节点获取 -----> 与远程节点交互 --> 返回缓存值 ⑵
                                |  否
                                |-----> 调用`回调函数`，获取值并添加到缓存 --> 返回缓存值 ⑶

2. 分布式:

使用一致性哈希选择节点        是                                    是
    |-----> 是否是远程节点 -----> HTTP 客户端访问远程节点 --> 成功？-----> 服务端返回返回值
                    |  否                                    ↓  否
                    |----------------------------> 回退到本地节点处理。
```
缓存实现过程
1. cache
       最基础的缓存(非并发安全),采用LRU(双向链表优先级判断)缓存淘汰机制
       缓存值类型需要实现Size()方法返回其内存大小
2. syncCache byteView
       封装cache,使其并发安全
       返回缓存类型为 ByteView 实现Size()方法,ByteSlice()返回缓存值的二进制副本,String()将缓存值转换为字符串 
3. http
       赋予缓存值远程访问的能力
```go
       主要依赖于 HttpPool 结构{
             self        string                 //self 用来记录自己的地址，包括 主机名/ip 和 端口。eg:"http://localhost:9999"
             basePath    string                 //basePath，作为节点间通讯地址的前缀
             mu          sync.RWMutex           //并发安全
             peersMap    *consistenthash.Map    //来根据具体的 key 选择节点。
             httpGetters map[string]*httpGetter //映射远程节点与对应的 httpGetter客户端。每一个远程节点对应一个 httpGetter，因为 httpGetter 与远程节点的地址 baseURL 有关
       }
       //客户端
       type httpGetter struct {
               baseURL string //所对应远程断点的ip地址,端口号和固定前缀,只需要缓存组名和缓存名即可进行通信
       }
```
4. consistenthash
   一致性哈希算法可以通过key值获取最适合的节点值
5. peers
```go
    // PeerPicker 对应 HttpPool,其 PickPeer() 方法用于根据传入的 key 选择相应节点 PeerGetter。
   type PeerPicker interface {
        PickPeer(key string) (peer PeerGetter, ok bool)
   }

   // PeerGetter 接口对应 HTTP 客户端 PeerGetter,其 Get() 方法用于从对应 group 查找缓存值
   type PeerGetter interface {
           Get(group string, key string) ([]byte, error)
   }
```
6. singleflight
   保证同一个查询在同一时刻最多执行一次.防止缓存穿透
7. RCache
   分布式缓存组,对上述内容的最高级封装
```go
type Group struct {
    name      string              //命名
    getter    Getter              //即缓存未命中时获取源数据的回调(callback)
    mainCache syncCache           //缓存
    peers     PeerPicker          //用于与其他节点通信的HttpPool结构
    loader    *singleflight.Group //保证一个请求同一时刻只会执行一次(缓存击破)
}
```

使用方法:
```go
    func createGroup() *RCache.Group {
        return RCache.NewGroup("scores", 2<<10, RCache.GetterFunc(func(key string) ([]byte, error) {
            //log.Println("[SlowDB] search key", key)
            //if v, ok := db[key]; ok {
            //	return []byte(v), nil
            //}
            ...
            return nil, fmt.Errorf("%s not exist", key)
        }))
    }
    1. 本地使用
    group := createGroup()
        view,err := group.Get("Tom")
        if err != nil {
        return
    }
    fmt.Println(view.String())
    2. 在线使用
    func main() {
        var port int
        var api bool
        flag.IntVar(&port, "port", 8001, "RCache server port")
        flag.BoolVar(&api, "api", false, "Start a api server?")
        flag.Parse()

    apiAddr := "localhost:9999"
    addrMap := map[int]string{
        8001: "localhost:8001",
        8002: "localhost:8002",
        8003: "localhost:8003",
    }
    //创建缓存组
    group := createGroup()
    //开启辅助端口
    if api {
        go group.StartServer(apiAddr, "/api", "key")
    }
    //开启分布式缓存
        group.StartCacheServer(addrMap[port], addrMap)
    }
```
        测试:
        $ go build -o server
        $ ./server.exe -port=8001 &
          ./server.exe -port=8002 &
          ./server.exe -port=8003 -api=1

        $ curl http://localhost:9999/api?key=Jack
          // 589
        $ curl http://localhost:8001/RCache/scores/Jack
          // 589
