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
		K Comparable
		V interface{}
	}

	Element struct {
		next []*Element
		kv   *KV
	}

	SkipList struct {
		node          Element     //第一个为头结点,不存值
		maxLevel      int         //最大层数
		len           int         //跳表长度
		randSource    rand.Source //生成随机数
		probTable     []float64   //用于查询每一层生成索引的概率
		prevNodeCache []*Element  // 用于保存查询一个值时经过每一层时的最后一个节点
	}
)

func NewKV(key Comparable, value interface{}) *KV {
	return &KV{
		K: key,
		V: value,
	}
}
func NewSkipList() *SkipList {
	return &SkipList{
		node:          Element{next: make([]*Element, MaxLevel)},
		maxLevel:      MaxLevel,
		randSource:    rand.NewSource(time.Now().UnixNano()),
		probTable:     probabilityTable(Probability, MaxLevel),
		prevNodeCache: make([]*Element, MaxLevel),
	}
}

// Len 获取长度
func (sList *SkipList) Len() int {
	return sList.len
}

// Get 获取数据 如果没有相等的数据，返回它的下一个元素(也可能是nil),bool仍是false
func (sList *SkipList) Get(k Comparable) (*KV, bool) {
	prev := &sList.node //重要 从头结点开始遍历
	var node *Element
	for now := sList.maxLevel - 1; now >= 0; now-- {
		node = prev.next[now]
		for node != nil && k.CompareTo(node.kv.K) > 0 {
			prev = node
			node = node.next[now]
		}
	}
	//指针已经到达第一层
	if node != nil {
		return node.kv, k.CompareTo(node.kv.K) == 0
	}
	return nil, false
}

// PrevNodeCache 每一层下来的节点所组成的切片
func (sList *SkipList) PrevNodeCache(k Comparable) []*Element {
	prev := &sList.node
	var node *Element
	for now := sList.maxLevel - 1; now >= 0; now-- { //在每一层进行搜索合适的位置然后向下一层
		node = prev.next[now]
		for node != nil && k.CompareTo(node.kv.K) > 0 { //发现没到就继续往右走
			prev = node //保留着向下的通道
			node = node.next[now]
		}
		sList.prevNodeCache[now] = prev //存放每层搜索的最后一个节点
	}
	return sList.prevNodeCache
}

// Delete 删除数据
func (sList *SkipList) Delete(key Comparable) {
	prev := sList.PrevNodeCache(key)
	if ele := prev[0].next[0]; ele != nil && key.CompareTo(ele.kv.K) == 0 { //如果找到了那个k
		for k, v := range ele.next { //把连接key的节点的next跳过key
			prev[k].next[k] = v
		}
		sList.len--
	}
}

func newElement(kv *KV, level int) *Element {
	return &Element{
		kv:   kv,
		next: make([]*Element, level),
	}
}

// Put 插入数据
func (sList *SkipList) Put(kv *KV) {
	prev := sList.PrevNodeCache(kv.K)                                        //保存着每层最后一个小于目标值得节点
	if ele := prev[0].next[0]; ele != nil && kv.K.CompareTo(ele.kv.K) == 0 { //找到就替换值
		ele.kv.V = kv.V
		return
	}
	//没找到就新创建一个
	element := newElement(kv, sList.randomLevel()) //随机分配一个层数
	//加入一个索引层
	for k := range element.next { //把k层的element都插入
		element.next[k] = prev[k].next[k]
		prev[k].next[k] = element
	}
	sList.len++
}

// GetMax 获取最后一个元素
func (sList *SkipList) GetMax() (kv *KV) {
	node := &sList.node
	for now := sList.maxLevel - 1; now >= 0; now-- {
		for node.next[now] != nil {
			node = node.next[now]
		}
	}
	return node.kv
}

// GetMin 获取第一个元素
func (sList *SkipList) GetMin() (kv *KV) {
	return sList.node.next[0].kv
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
	for level < sList.maxLevel && n < sList.probTable[level-1] {
		level++
	}
	return level
}
