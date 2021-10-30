package RCache

import (
	"Rutils/RCache/cache"
	"sync"
)

//提供并发保护

type syncCache struct {
	mu         sync.RWMutex
	cache      *cache.Cache
	cacheBytes int
}

func (sc *syncCache) add(key string, value ByteView) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	if sc.cache == nil {
		//默认回调函数为nil
		sc.cache = cache.NewWithBytes(sc.cacheBytes, nil)
	}
	sc.cache.Add(key, value)
}

func (sc *syncCache) get(key string) (value ByteView, ok bool) {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	if sc.cache == nil {
		return
	}
	//v是Value接口类型,value是Value接口类型,v的实例实现了Size方法,即实现了Value类型
	if v, ok := sc.cache.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
