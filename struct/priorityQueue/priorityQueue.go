package main

import (
	"container/heap"
	"fmt"
)

// 最小堆

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].priority > (*pq)[j].priority
}
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index, (*pq)[j].index = i, j
}

func (pq *PriorityQueue) Pop() interface{} {
	t := (*pq)[len(*pq)-1]
	t.index = -1 // 标记被弹出
	*pq = (*pq)[:len(*pq)-1]
	return t
}

// Push 添加节点
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

// Update 更新节点
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	// 创建节点并设计他们的优先级
	items := map[string]int{"二毛": 5, "张三": 3, "狗蛋": 9}
	i := 0
	pq := make(PriorityQueue, len(items)) // 创建优先级队列，并初始化
	for k, v := range items {             // 将节点放到优先级队列中
		pq[i] = &Item{
			value:    k,
			priority: v,
			index:    i}
		i++
	}
	heap.Init(&pq) // 初始化堆
	item := &Item{ // 创建一个item
		value:    "李四",
		priority: 1,
	}
	heap.Push(&pq, item)           // 入优先级队列
	pq.Update(item, item.value, 6) // 更新item的优先级
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s index:%.2d\n", item.priority, item.value, item.index)
	}
}
