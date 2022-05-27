package link

import "fmt"

//循环链表

type Ring struct {
	next, prev *Ring
	data       interface{}
}

func InitRing() *Ring {
	var ring Ring
	ring.prev = &ring
	ring.next = &ring
	return &ring
}

// InsertByHead 头插法
func (r *Ring) InsertByHead(data ...interface{}) {
	for i := 0; i < len(data); i++ {
		p := &Ring{
			next: r.next,
			prev: r,
			data: data[i],
		}
		r.next.prev = p
		r.next = p
	}
}

// InsertByTail 尾插法
func (r *Ring) InsertByTail(data ...interface{}) {
	for i := 0; i < len(data); i++ {
		p := &Ring{
			next: r,
			prev: r.prev,
			data: data[i],
		}
		r.prev.next = p
		r.prev = p
	}
}

func (r *Ring) Print() {
	for p := r.next; p != r; p = p.next {
		fmt.Println(p.data)
	}
}

func (r *Ring) PrevNode() *Ring {
	if r.prev == nil {
		return InitRing()
	}
	return r.prev
}

func (r *Ring) NextNode() *Ring {
	if r.next == nil {
		return InitRing()
	}
	return r.next
}

func (r *Ring) NodeData() interface{} {
	return r.data
}

func (r *Ring) Move(n int) *Ring {
	if r.next == nil || r.prev == nil {
		return InitRing()
	}
	p := r
	switch n > 0 {
	case true:
		for i := 0; i < n; i++ {
			p = p.next
		}
	default:
		for i := 0; i < n; i++ {
			p = p.prev
		}
	}
	return p
}

// Link 往节点A，链接一个节点，并且返回之前节点A的后驱节点
func (r *Ring) Link(s *Ring) *Ring {
	n := r.next
	if s != nil {
		p := s.prev
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}
