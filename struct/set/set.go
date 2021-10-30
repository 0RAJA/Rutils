package set

import "sync"

// Set 集合
type Set struct {
	m            map[int]struct{} //字典实现
	len          int              //集合大小
	sync.RWMutex                  //锁
}

func NewSet(cap int64) *Set {
	var temp map[int]struct{}
	if cap <= 0 {
		temp = make(map[int]struct{})
	} else {
		temp = make(map[int]struct{}, cap)
	}
	return &Set{
		m:       temp,
		len:     0,
		RWMutex: sync.RWMutex{},
	}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
	s.len = len(s.m)
}

func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	if s.len == 0 {
		return
	}
	delete(s.m, item)
	s.len = len(s.m)
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return s.len
}

func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{}
	s.len = 0
}

func (s *Set) List() (list []int) {
	s.RLock()
	defer s.RUnlock()
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
