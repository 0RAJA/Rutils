package no_lock_queue

import (
	"sync/atomic"
	"unsafe"
)

// LKQueue 无锁队列
type LKQueue struct {
	count int64          // 元素个数
	head  unsafe.Pointer // 第一个
	tail  unsafe.Pointer // 最后一个
}

// 节点
type node struct {
	value interface{}    // 当前 value
	next  unsafe.Pointer // next
}

// NewLKQueue returns an empty queue.
func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{head: n, tail: n}
}

// 原子读
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// cas
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}

// Enqueue 写入元素到队列
func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{value: v}
	defer func() { atomic.AddInt64(&q.count, 1) }()
	for {
		tail := load(&q.tail)      // 当前最后一个
		next := load(&tail.next)   // 最后一个的下一个
		if tail == load(&q.tail) { // 判断是否被修改
			if next == nil { // 这是最后一个
				if cas(&tail.next, next, n) {
					cas(&q.tail, tail, n) // 尝试将当前队列末尾指向这个节点，失败说明有其他的操作了
					return
				}
			} else { // 不是最后一个 试着把tail摆动到下一个节点
				cas(&q.tail, tail, next)
			}
		}
	}
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *LKQueue) Dequeue() interface{} {
	if atomic.LoadInt64(&q.count) == 0 {
		return nil
	}
	defer func() { atomic.AddInt64(&q.count, -1) }()
	for {
		head := load(&q.head)      // 头
		tail := load(&q.tail)      // 尾
		next := load(&head.next)   // 头的下一个
		if head == load(&q.head) { // are head, tail, and next consistent?
			if head == tail { // is queue empty or tail falling behind?
				if next == nil { // is queue empty?
					return nil
				}
				// tail is falling behind.  try to advance it
				cas(&q.tail, tail, next)
			} else {
				// read value before CAS otherwise another dequeue might free the next node
				v := next.value
				if cas(&q.head, head, next) {
					return v // Dequeue is done.  return
				}
			}
		}
	}
}

// Count 计数
func (q *LKQueue) Count() int64 {
	return atomic.LoadInt64(&q.count)
}
