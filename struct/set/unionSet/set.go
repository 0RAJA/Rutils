package unionSet

type Set struct {
	parent, size []int
}

func NewSet(n int) *Set {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range size {
		parent[i] = i
		size[i] = 1
	}
	return &Set{parent: parent, size: size}
}

// Find 查找x所在的集合
func (s *Set) Find(x int) int {
	if s.parent[x] != x {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

// Union 联合
func (s *Set) Union(x, y int) {
	xf := s.Find(x)
	yf := s.Find(y)
	if xf == yf {
		return
	}
	if s.size[xf] > s.size[yf] {
		xf, yf = yf, xf
	}
	s.size[xf] += s.size[yf]
	s.parent[yf] = xf
}

// InSameSet 判断xy是否在同一个集合
func (s *Set) InSameSet(x, y int) bool {
	return s.Find(x) == s.Find(y)
}
