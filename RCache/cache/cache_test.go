package cache

import (
	"Rutils/RMemory"
	"fmt"
	"testing"
)

type Mints []int

func (m Mints) Size() int {
	return len(m)*RMemory.SizeOfMemoryInt(int(1)) + RMemory.SizeOfMemoryInt(Mints{})
}

func TestCache_Get(t *testing.T) {
	MCache := New(nil)
	key1, value1 := "key1", Mints{1, 2, 3}
	MCache.Add(key1, value1)
	if v, ok := MCache.Get(key1); ok {
		fmt.Println("OK", v)
	} else {
		fmt.Println("NO")
	}
}

func TestCache_ReMoveOldest(t *testing.T) {
	MCache := NewWithBytes(20, nil)
	key1, value1 := "key1", Mints{1, 2, 3}
	MCache.Add(key1, value1)
	if v, ok := MCache.Get(key1); ok {
		fmt.Println("OK", v)
	} else {
		fmt.Println("NO")
	}
}

func TestCache_OnEvicted(t *testing.T) {
	MCache := New(func(key string, value Value) {
		fmt.Println(key, value.(Mints), "已经被删除")
	})
	key1, value1 := "key1", Mints{1, 2, 3}
	MCache.Add(key1, value1)
	MCache.Del(key1)
}
