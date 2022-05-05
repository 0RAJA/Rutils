package RCache

import (
	"fmt"
	"github.com/0RAJA/Rutils/RCache/RErr"
	"github.com/0RAJA/Rutils/RCache/consistenthash"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

//提供被其他节点访问的能力(基于http)

const (
	// DefBasePath 默认路由前缀
	DefBasePath = "/RCache/"
	// DefReplicas 默认虚拟节点倍率
	DefReplicas = 50
	// DefHTTP 默认前缀
	DefHTTP = "http://"
)

// HttpPool 节点间 HTTP 通信的核心数据结构
type HttpPool struct {
	self        string                 //self 用来记录自己的地址，包括 主机名/ip 和 端口。eg:"http://localhost:9999"
	basePath    string                 //basePath，作为节点间通讯地址的前缀
	mu          sync.RWMutex           //并发安全
	peersMap    *consistenthash.Map    //来根据具体的 key 选择节点。
	httpGetters map[string]*httpGetter //映射远程节点与对应的 httpGetter。每一个远程节点对应一个 httpGetter，因为 httpGetter 与远程节点的地址 baseURL 有关
}

//客户端
type httpGetter struct {
	baseURL string
}

// Get 给远程节点发起请求
func (hg *httpGetter) Get(group string, key string) ([]byte, error) {
	u := fmt.Sprintf("%v%v/%v", hg.baseURL, url.QueryEscape(group), url.QueryEscape(key))
	//发起请求
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	//判断状态码
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}
	//读取信息
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}
	return bytes, nil
}

func NewHTTPPool(self string) *HttpPool {
	return &HttpPool{self: self, basePath: DefBasePath}
}

func (hp *HttpPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", hp.self, fmt.Sprintf(format, v...))
}

//回应远程请求
func (hp *HttpPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, hp.basePath) {
		http.Error(w, RErr.BadRequest, http.StatusNotFound)
		return
		//panic(RErr.UrlPathErr + r.URL.LocalPath)
	}
	hp.Log("%s %s", r.Method, r.URL.Path)
	// /<basepath>/<groupname>/<key> required 解析url
	parts := strings.SplitN(r.URL.Path[len(hp.basePath):], "/", 2)
	if len(parts) != 2 {
		//返回错误
		http.Error(w, RErr.BadRequest, http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	key := parts[1]
	//获取缓存组
	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, RErr.NoGroup, http.StatusNotFound)
		return
	}
	view, err := group.Get(key)
	if err != nil {
		//服务器内部错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//设置响应头中的内容类型为二进制流数据
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}

// Set 存储其他节点
func (hp *HttpPool) Set(peers ...string) {
	hp.mu.Lock()
	defer hp.mu.Unlock()
	//将节点存入hash环
	hp.peersMap = consistenthash.New(DefReplicas, nil)
	hp.peersMap.Add(peers...)
	hp.httpGetters = make(map[string]*httpGetter, len(peers))
	//为每一个节点创建了一个 HTTP 客户端 httpGetter
	for _, peer := range peers {
		hp.httpGetters[peer] = &httpGetter{baseURL: DefHTTP + peer + hp.basePath}
	}
}

// PickPeer 为 key 选择一个节点
func (hp *HttpPool) PickPeer(key string) (PeerGetter, bool) {
	hp.mu.Lock()
	defer hp.mu.Unlock()
	if peer := hp.peersMap.Get(key); peer != "" && peer != hp.self {
		hp.Log("Pick peer %s", peer)
		return hp.httpGetters[peer], true
	}
	return nil, false
}
