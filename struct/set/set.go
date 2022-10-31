package set

import "sync"

// Set 集合
type Set struct {
	m  map[interface{}]struct{} // 字典实现
	rw sync.RWMutex             // 锁
}

func New() *Set {
	return &Set{
		m:  make(map[interface{}]struct{}),
		rw: sync.RWMutex{},
	}
}

func NewSet(cap int) *Set {
	var temp map[interface{}]struct{}
	if cap <= 0 {
		temp = make(map[interface{}]struct{})
	} else {
		temp = make(map[interface{}]struct{}, cap)
	}
	return &Set{
		m:  temp,
		rw: sync.RWMutex{},
	}
}

func (s *Set) Add(item interface{}) {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.m[item] = struct{}{}
}

func (s *Set) Remove(item int) {
	s.rw.Lock()
	defer s.rw.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item int) bool {
	s.rw.RLock()
	defer s.rw.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set) Clear() {
	s.rw.Lock()
	defer s.rw.Unlock()
	s.m = map[interface{}]struct{}{}
}

func (s *Set) List() (list []interface{}) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
