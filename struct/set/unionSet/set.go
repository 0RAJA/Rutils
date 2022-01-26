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

func (s *Set) find(x int) int {
	if s.parent[x] != x {
		s.parent[x] = s.find(s.parent[x])
	}
	return s.parent[x]
}

func (s *Set) union(x, y int) {
	xf := s.find(x)
	yf := s.find(y)
	if xf == yf {
		return
	}
	if s.size[xf] > s.size[yf] {
		xf, yf = yf, xf
	}
	s.size[xf] += s.size[yf]
	s.parent[yf] = xf
}

func (s *Set) InSameSet(x, y int) bool {
	return s.find(x) == s.find(y)
}
