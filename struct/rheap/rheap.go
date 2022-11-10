package rheap

import (
	"sort"
)

type HeapInterface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

// 向下调整
func down(h HeapInterface, idx, length int) bool {
	t := idx
	for l := t*2 + 1; l < length && l > 0; {
		if h.Less(l, idx) {
			t = l
		}
		r := l + 1
		if r < length && r > 0 && h.Less(r, t) {
			t = r
		}
		if t == idx {
			break
		} else {
			h.Swap(t, idx)
			idx = t
		}
	}
	return t > idx
}

// Init 构建堆
func Init(h HeapInterface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func up(h HeapInterface, j int) {
	for {
		i := (j - 1) / 2
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func Push(h HeapInterface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h HeapInterface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Remove(h HeapInterface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}
