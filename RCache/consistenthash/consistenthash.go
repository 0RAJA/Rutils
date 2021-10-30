package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

/*一致性哈希算法的实现
一致性哈希算法将 key 映射到 2^32 的空间中，将这个数字首尾相连，形成一个环.
	1. 计算节点/机器(通常使用节点的名称、编号和 IP 地址)的哈希值，放置在环上。
	2. 计算 key 的哈希值，放置在环上，顺时针寻找到的第一个节点，就是应选取的节点/机器。
优点:
	一致性哈希算法，在新增/删除节点时，只需要重新定位该节点附近的一小部分数据，而不需要重新定位所有的节点
问题:数据倾斜
解决:使用虚拟节点扩充节点数
	1. 计算虚拟节点的 Hash 值，放置在环上。
	2. 计算 key 的 Hash 值，在环上顺时针寻找到应选取的虚拟节点，例如是 peer2-1，那么就对应真实节点 peer2。
*/

type Hash func(data []byte) uint32

type Map struct {
	mu       sync.RWMutex   //并发安全
	hash     Hash           //Hash函数
	replicas int            //虚拟节点倍数
	keys     []int          //哈希环 存着节点的hash值
	hashMap  map[int]string //虚拟节点与真实节点的映射表,键是虚拟节点的哈希值,值是真实节点的名称
}

// New 构造函数 允许自定义hash函数和节点倍数
func New(replicas int, hash Hash) *Map {
	m := &Map{hash: hash, replicas: replicas, hashMap: map[int]string{}}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add 增加节点
func (m *Map) Add(keys ...string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			//加入哈希环
			m.keys = append(m.keys, hash)
			//建立映射
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get 获取对应节点
func (m *Map) Get(key string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	//通过hash值找到在key之后的第一个节点,因为可能i为n,所以取余,使其称为环
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

// Remove 移除节点
func (m *Map) Remove(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := 0; i < m.replicas; i++ {
		hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
		idx := sort.SearchInts(m.keys, hash)
		if idx == len(m.keys) {
			continue
		}
		m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
		delete(m.hashMap, hash)
	}
}
