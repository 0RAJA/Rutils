package main

import (
	"Rutils/RCache"
	"flag"
	"fmt"
	"log"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

//创建缓存组
func createGroup() *RCache.Group {
	return RCache.NewGroup("scores", 2<<10, RCache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))
}

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
