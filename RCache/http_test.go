package RCache

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHttpPool_ServeHTTP(t *testing.T) {
	var db = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}
	NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}

/*
测试:
http://localhost:9999/RCache/scores/Tom
	630
http://localhost:9999/RCache/scores/kkk
	kkk not exist
http://localhost:9999/RCache/scores/Tom
	630
*/
