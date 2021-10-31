package set

import "sync"

// set 集合
type set struct {
	m            map[interface{}]struct{} //字典实现
	sync.RWMutex                          //锁
}

func New() *set {
	return &set{
		m:       make(map[interface{}]struct{}),
		RWMutex: sync.RWMutex{},
	}
}

func NewSet(cap int) *set {
	var temp map[interface{}]struct{}
	if cap <= 0 {
		temp = make(map[interface{}]struct{})
	} else {
		temp = make(map[interface{}]struct{}, cap)
	}
	return &set{
		m:       temp,
		RWMutex: sync.RWMutex{},
	}
}

func (s *set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
}

func (s *set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *set) Len() int {
	return len(s.m)
}

func (s *set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]struct{}{}
}

func (s *set) List() (list []interface{}) {
	s.RLock()
	defer s.RUnlock()
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
