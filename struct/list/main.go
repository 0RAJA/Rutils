package list

import (
	"sync"
)

const (
	OverLenErr = "index out"
)

/*
列表 List ：存放数据，数据按顺序排列，可以依次入队和出队，有序号关系，可以取出某序号的数据。
先进先出的 队列 (queue) 和先进后出的 栈（stack） 都是列表。
大家也经常听说一种叫 线性表 的数据结构，表示具有相同特性的数据元素的有限序列，实际上就是 列表 的同义词。

双端列表，也可以叫双端队列
我们会用双向链表来实现这个数据结构：
*/

type DoubleList struct {
	head *DoubleNode // 头
	tail *DoubleNode // 尾
	len  int         // 长度
	lock sync.Mutex  // 并发pop
}

/*
链表的第一个元素的前驱节点为 nil ，最后一
个元素的后驱节点也为 nil
*/

type DoubleNode struct {
	pre   *DoubleNode // 前驱
	next  *DoubleNode // 后驱
	value string      // 值
}

func (node *DoubleNode) GetValue() string {
	if node == nil {
		return ""
	}
	return node.value
}

func (node *DoubleNode) GetPre() *DoubleNode {
	if node == nil {
		return nil
	}
	return node.pre
}

func (node *DoubleNode) GetNext() *DoubleNode {
	if node == nil {
		return nil
	}
	return node.next
}

func (node *DoubleNode) HashNext() bool {
	if node == nil {
		return false
	}
	return node.next != nil
}

func (node *DoubleNode) HashPre() bool {
	if node == nil {
		return false
	}
	return node.pre != nil
}

func (node *DoubleNode) IsNil() bool {
	return node == nil
}

func DoubleListInit() *DoubleList {
	return &DoubleList{
		head: nil,
		tail: nil,
		len:  0,
		lock: sync.Mutex{},
	}
}

// AddNodeFromHead 添加节点到链表头部的第N个元素之前，N=0表示新节点成为新的头部
func (list *DoubleList) AddNodeFromHead(n int, v string) {
	// 加锁
	list.lock.Lock()
	defer list.lock.Unlock()
	// 索引超过列表长度,panic
	if n > list.len {
		panic(OverLenErr)
	}
	// 找到头
	node := list.head
	// 遍历找第n个元素
	for i := 1; i < n; i++ {
		node = node.next
	}
	newNode := new(DoubleNode)
	newNode.value = v
	// 说明列表为空
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点的前一个
		pre := node.pre
		// node节点是头,需要换个头
		if pre.IsNil() {
			newNode.next = node
			node.pre = newNode
			// 成为新头
			list.head = newNode
		} else {
			pre.next = newNode
			newNode.pre = pre
			node.pre = newNode
			newNode.next = node
		}
	}
	// 列表长度++
	list.len++
}

// AddNodeFromTail 添加节点到链表尾部的第N个元素之前，N=0表示新节点成为新的尾部
func (list *DoubleList) AddNodeFromTail(n int, v string) {
	// 加锁
	list.lock.Lock()
	defer list.lock.Unlock()
	// 索引超过列表长度,panic
	if n > list.len {
		panic(OverLenErr)
	}
	// 找到尾部
	node := list.tail
	// 遍历找第n个元素
	for i := 1; i < n; i++ {
		node = node.pre
	}
	newNode := new(DoubleNode)
	newNode.value = v
	// 说明列表为空
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		// 定位到的节点的前一个
		pre := node.next
		// node节点是头,需要换个头
		if pre.IsNil() {
			newNode.pre = node
			node.next = newNode
			// 成为新头
			list.tail = newNode
		} else {
			pre.pre = newNode
			newNode.next = pre
			node.next = newNode
			newNode.pre = node
		}
	}
	// 列表长度++
	list.len++
}

// IndexFromHead 从头部开始某个位置获取列表节点,索引从0开始。
func (list *DoubleList) IndexFromHead(n int) *DoubleNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}
	node := list.head
	for i := 0; i < n; i++ {
		node = node.next
	}
	return node
}

// IndexFromTail 从头部开始某个位置获取列表节点,索引从0开始。
func (list *DoubleList) IndexFromTail(n int) *DoubleNode {
	// 索引超过或等于列表长度，一定找不到，返回空指针
	if n >= list.len {
		return nil
	}
	node := list.tail
	for i := 0; i < n; i++ {
		node = node.pre
	}
	return node
}

// PopFromHead 从头部开始移除并返回第n个的节点
func (list *DoubleList) PopFromHead(n int) *DoubleNode {
	// 加锁
	list.lock.Lock()
	defer list.lock.Unlock()
	// 索引超过长度
	if n > list.len || n <= 0 {
		return nil
	}
	node := list.head
	// 找到第n个元素
	for i := 1; i < n; i++ {
		node = node.next
	}
	if !node.pre.IsNil() {
		node.pre.next = node.next
	} else {
		list.head = node.next
	}
	if !node.next.IsNil() {
		node.next.pre = node.pre
	} else {
		list.tail = node.pre
	}
	list.len--
	node.pre = nil
	node.next = nil
	return node
}

// PopFromTail 从尾部开始移除并返回第n个节点
func (list *DoubleList) PopFromTail(n int) *DoubleNode {
	// 加锁
	list.lock.Lock()
	defer list.lock.Unlock()
	// 索引超过长度
	if n > list.len || n <= 0 {
		return nil
	}
	node := list.tail
	// 找到第n个元素
	for i := 1; i < n; i++ {
		node = node.pre
	}
	if !node.pre.IsNil() {
		node.pre.next = node.next
	} else {
		list.head = node.next
	}
	if !node.next.IsNil() {
		node.next.pre = node.pre
	} else {
		list.tail = node.pre
	}
	list.len--
	node.pre = nil
	node.next = nil
	return node
}

func (list *DoubleList) PrintHead() (slice []string) {
	for i := list.head; i != nil; i = i.next {
		slice = append(slice, i.value)
	}
	return
}

func (list *DoubleList) PrintTail() (slice []string) {
	for i := list.tail; i != nil; i = i.pre {
		slice = append(slice, i.value)
	}
	return
}
