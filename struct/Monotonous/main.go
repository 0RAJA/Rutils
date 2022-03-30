package Monotonous

import "container/list"

//单调队列
type monotonous interface {
	Pop() interface{}
	Push(v Comparable)
	Top() interface{}
	Min() interface{}
	Max() interface{}
}

type Structure struct {
	monotonous
}

type Comparable interface {
	CompareTo(v Comparable) int
}

func NewStructure(monotonous monotonous) *Structure {
	return &Structure{monotonous: monotonous}
}

type Assists struct {
	nums      list.List
	assistMax list.List
	assistMin list.List
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
