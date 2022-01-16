package Monotonous

import (
	"fmt"
	"testing"
)

type M int

func (m M) CompareTo(v Comparable) int {
	return int(m - v.(M))
}

func TestNewStructure(t *testing.T) {
	queue := NewStructure(new(Queue))
	queue.Push(M(1))
	queue.Push(M(2))
	fmt.Println(queue.Max())
	fmt.Println(queue.Min())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Top())
	fmt.Println(queue.Max())
	fmt.Println(queue.Min())
}
