package cache

import (
	"container/list"
)

//基础缓存

const (
	MB             = 1024 * 1024
	DefMaxBytes    = 32 * MB
	DefStringBytes = 16
)

// Value 所存储值的类型
type Value interface {
	// Size 用于计算此值的类型所占用的内存
	Size() int
	//防止有些像[]int{1,2,3}这种类型,需要手动计算大小
}

// Cache 缓存
type Cache struct {
	//最大内存
	maxBytes int
	//已使用内存
	usedBytes int
	//缓存的value为指向双向链表的节点的指针
	cache map[string]*list.Element
	//缓存所构成双向链表,用于缓存淘汰
	ll *list.List
	//回调函数,当缓存中的某条数据要被清理时,会调用此函数,回调函数可以为nil
	onEvicted func(key string, value Value)
}

//双向链表节点
type entry struct {
	//此key用于在淘汰缓存后删除map中响应的映射
	key   string
	value Value
}

// New 实例化
func New(onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes:  DefMaxBytes,
		cache:     make(map[string]*list.Element),
		ll:        list.New(),
		onEvicted: onEvicted,
	}
}

// NewWithBytes 指定最大内存实例化
func NewWithBytes(maxBytes int, onEvicted func(key string, value Value)) *Cache {
	//修订非法输入
	if maxBytes <= 0 || maxBytes > DefMaxBytes {
		maxBytes = DefMaxBytes
	}
	return &Cache{
		maxBytes:  maxBytes,
		cache:     make(map[string]*list.Element),
		ll:        list.New(),
		onEvicted: onEvicted,
	}
}

// Get 获取
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		//提升优先度
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

// ReMoveOldest 缓存淘汰
func (c *Cache) ReMoveOldest() {
	//拿到优先级最低的元素
	ele := c.ll.Back()
	if ele != nil {
		//从双向链表中删除
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		//从map中删除
		delete(c.cache, kv.key)
		//更新已使用内存
		c.usedBytes -= DefStringBytes + kv.value.Size()
		//调用回调函数
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
}

// Del 删除
func (c *Cache) Del(key string) {
	ele, ok := c.cache[key]
	//不存在直接返回
	if !ok {
		return
	}
	//从链表中删除
	c.ll.Remove(ele)
	kv := ele.Value.(*entry)
	//从map中删除
	delete(c.cache, kv.key)
	//更新已使用内存
	c.usedBytes -= DefStringBytes + kv.value.Size()
	//调用回调函数
	if c.onEvicted != nil {
		c.onEvicted(kv.key, kv.value)
	}
	return
}

// Add 增加
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		//提高优先度
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		//更新已使用内存
		c.usedBytes += value.Size() - kv.value.Size()
		//更新内容
		kv.value = value
	} else {
		//直接插入
		ele = c.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		//保存到map中
		c.cache[key] = ele
		//更新已使用内存
		c.usedBytes += DefStringBytes + value.Size()
	}
	//淘汰缓存
	for c.usedBytes >= c.maxBytes {
		c.ReMoveOldest()
	}
}

// Update 更新
func (c *Cache) Update(key string, value Value) {
	c.Add(key, value)
}
