package Monotonous

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
