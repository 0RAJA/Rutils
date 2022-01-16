package Monotonous

import "container/list"

type Comparable interface {
	CompareTo(v Comparable) int
}

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

func NewStructure(monotonous monotonous) *Structure {
	return &Structure{monotonous: monotonous}
}

type Assists struct {
	nums      list.List
	assistMax list.List
	assistMin list.List
}

type Queue Assists

func (q *Queue) Push(v Comparable) {
	q.nums.PushBack(v)
	q.PushAssist(v)
}

func (q *Queue) Pop() interface{} {
	if q.nums.Len() <= 0 {
		return nil
	}
	t := q.nums.Front()
	q.PopAssist(t.Value.(Comparable))
	q.nums.Remove(t)
	return t.Value
}
func (q *Queue) Top() interface{} {
	if q.nums.Len() <= 0 {
		return nil
	}
	return q.nums.Front().Value
}
func (q *Queue) PopAssist(v Comparable) {
	if v.CompareTo(q.assistMin.Front().Value.(Comparable)) == 0 {
		q.assistMin.Remove(q.assistMin.Front())
	}
	if v.CompareTo(q.assistMax.Front().Value.(Comparable)) == 0 {
		q.assistMax.Remove(q.assistMax.Front())
	}
}

func (q *Queue) PushAssist(v Comparable) {
	for q.assistMin.Len() > 0 && v.CompareTo(q.assistMin.Back().Value.(Comparable)) < 0 {
		q.assistMin.Remove(q.assistMin.Back())
	}
	q.assistMin.PushBack(v)
	for q.assistMax.Len() > 0 && v.CompareTo(q.assistMax.Back().Value.(Comparable)) > 0 {
		q.assistMax.Remove(q.assistMax.Back())
	}
	q.assistMax.PushBack(v)
}

func (q *Queue) Min() interface{} {
	return q.assistMin.Front().Value
}

func (q *Queue) Max() interface{} {
	return q.assistMax.Front().Value
}

type Stack Assists

func (s *Stack) Push(v Comparable) {
	s.nums.PushBack(v)
	s.PushAssist(v)
}

func (s *Stack) Pop() interface{} {
	if s.nums.Len() <= 0 {
		return nil
	}
	t := s.nums.Back()
	s.PopAssist(t.Value.(Comparable))
	s.nums.Remove(t)
	return t.Value
}
func (s *Stack) Top() interface{} {
	if s.nums.Len() <= 0 {
		return nil
	}
	return s.nums.Back().Value
}
func (s *Stack) PopAssist(v Comparable) {
	if v.CompareTo(s.assistMin.Back().Value.(Comparable)) == 0 {
		s.assistMin.Remove(s.assistMin.Back())
	}
	if v.CompareTo(s.assistMax.Back().Value.(Comparable)) == 0 {
		s.assistMax.Remove(s.assistMax.Back())
	}
}

func (s *Stack) PushAssist(v Comparable) {
	if v.CompareTo(s.assistMin.Back().Value.(Comparable)) < 0 {
		s.assistMin.PushBack(v)
	}
	if v.CompareTo(s.assistMax.Back().Value.(Comparable)) > 0 {
		s.assistMax.PushBack(v)
	}
}

func (s *Stack) Min() interface{} {
	return s.assistMin.Back().Value
}

func (s *Stack) Max() interface{} {
	return s.assistMax.Back().Value
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
