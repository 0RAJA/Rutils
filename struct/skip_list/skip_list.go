package skip_list

import (
	"math"
	"math/rand"
	"time"
)

const (
	MaxLevel    = 18
	Probability = 1 / math.E
)

type (
	// Comparable 大于返回正数,小于返回负数,等于返回0
	Comparable interface {
		CompareTo(Comparable) int
	}
	KV struct {
		k Comparable
		v interface{}
	}

	Element struct {
		next []*Element
		kv   KV
	}

	SkipList struct {
		Node          Element     //第一个为头结点,不存值
		MaxLevel      int         //最大层数
		Len           int         //跳表长度
		randSource    rand.Source //生成随机数
		probTable     []float64   //用于查询每一层生成索引的概率
		prevNodeCache []*Element  // 用于保存查询一个值时经过每一层时的最后一个节点
	}
)

func NewSkipList() *SkipList {
	return &SkipList{
		Node:          Element{next: make([]*Element, MaxLevel)},
		MaxLevel:      MaxLevel,
		randSource:    rand.NewSource(time.Now().UnixNano()),
		probTable:     probabilityTable(Probability, MaxLevel),
		prevNodeCache: make([]*Element, MaxLevel),
	}
}

func (sList *SkipList) Get(k Comparable) (interface{}, bool) {
	prev := &sList.Node //重要 从头结点开始遍历
	var node *Element
	for now := sList.MaxLevel - 1; now >= 0; now-- {
		node = prev.next[now]
		for node != nil && k.CompareTo(node.kv.k) > 0 {
			prev = node
			node = node.next[now]
		}
	}
	//指针已经到达第一层
	if node != nil && k.CompareTo(node.kv.k) == 0 {
		return node.kv.v, true
	}
	return nil, false
}

func (sList *SkipList) PrevNodeCache(k Comparable) []*Element {
	prev := &sList.Node
	var node *Element
	for now := sList.MaxLevel - 1; now >= 0; now-- {
		node = prev.next[now]
		for node != nil && k.CompareTo(node.kv.k) > 0 {
			prev = node
			node = node.next[now]
		}
		sList.prevNodeCache[now] = prev //存放每层搜索的最后一个节点
	}
	return sList.prevNodeCache
}

func (sList *SkipList) Delete(key Comparable) {
	prev := sList.PrevNodeCache(key)
	if ele := prev[0].next[0]; ele != nil && key.CompareTo(ele.kv.k) == 0 { //如果找到了那个k
		for k, v := range ele.next { //把连接key的节点的next跳过key
			prev[k].next[k] = v
		}
		sList.Len--
	}
}

func newElement(kv KV, level int) *Element {
	return &Element{
		kv:   kv,
		next: make([]*Element, level),
	}
}

func (sList *SkipList) Put(key Comparable, value interface{}) {
	prev := sList.PrevNodeCache(key)
	if ele := prev[0].next[0]; ele != nil && key.CompareTo(ele.kv.k) == 0 {
		ele.kv.v = value
		return
	}
	element := newElement(KV{key, value}, sList.randomLevel()) //随机分配一个层数
	//加入一个索引层
	for k := range element.next {
		element.next[k] = prev[k].next[k]
		prev[k].next[k] = element
	}
	sList.Len++
}

//计算每一层生成索引的概率-> 生成一张概率表
func probabilityTable(probability float64, level int) (ret []float64) {
	for i := 1; i <= level; i++ {
		ret = append(ret, math.Pow(probability, float64(i-1)))
	}
	return ret
}

//对每一层进行概率评估
func (sList *SkipList) randomLevel() int {
	n := float64(sList.randSource.Int63()) / math.MaxInt64 //随机概率
	level := 1                                             //第1层始终有索引
	for level < sList.MaxLevel && n < sList.probTable[level-1] {
		level++
	}
	return level
}
