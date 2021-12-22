package RHash

import (
	"math"
)

type KV struct {
	k, v string
}

type Map interface {
	new(cap int) Map
	operation(k string, f func(kv *KV) bool) (int, *KV, bool) //遍历下标
	setValue(idx int, value *KV)                              //设置值
	delValue(idx int, k string)                               //删除值
	rehash() Map
}

type M struct {
	Map
}

func NewM(m Map, cap int) *M {
	if cap == 0 {
		cap = 1
	}
	return &M{Map: m.new(cap)}
}

func (m *M) Set(k, v string) {
	idx, _, ok := m.operation(k, func(kv *KV) bool {
		return kv == nil
	})
	if !ok {
		m.Map = m.rehash()
		m.Set(k, v)
	} else {
		m.setValue(idx, &KV{k: k, v: v})
	}
}

func (m *M) Get(k string) (string, bool) {
	if _, value, ok := m.operation(k, func(kv *KV) bool { return kv != nil && kv.k == k }); ok {
		return value.v, true
	}
	return "", false
}

func (m *M) Del(k string) bool {
	idx, _, ok := m.operation(k, func(kv *KV) bool {
		return kv != nil && kv.k == k
	})
	if ok {
		m.delValue(idx, k)
	}
	return ok
}

type MyMap1 []*KV

func (m MyMap1) new(cap int) Map {
	p := make(MyMap1, cap)
	return p
}

func (m MyMap1) setValue(idx int, value *KV) {
	m[idx] = value
}

func (m MyMap1) delValue(idx int, k string) {
	m[idx] = nil
}

func (m MyMap1) operation(k string, f func(kv *KV) bool) (int, *KV, bool) {
	length := len(m)
	index := toHash(k) % length
	for i := 0; i < length; i++ {
		if x := (index + i) % length; f(m[x]) {
			return x, m[x], true
		}
	}
	return 0, nil, false
}

func (m MyMap1) rehash() Map {
	length := len(m)
	m = append(m, make(MyMap1, 2*length)...)
	for i := 0; i < length; i++ {
		if kv := m[i]; kv != nil {
			m.delValue(i, kv.k)
			m.setValue(i, kv)
		}
	}
	return m
}

type MyMap2 []*KV

func (m MyMap2) new(cap int) Map {
	p := make(MyMap2, cap)
	return p
}

func (m MyMap2) setValue(idx int, value *KV) {
	m[idx] = value
}

func (m MyMap2) delValue(idx int, k string) {
	m[idx] = nil
}

func (m MyMap2) operation(k string, f func(idx *KV) bool) (int, *KV, bool) {
	length := len(m)
	index := toHash(k) % length
	var t int
	for x := 0; x <= length/2; x++ {
		if t = (index + x*x) % length; f(m[t]) {
			return t, m[t], true
		}
		t = (index - x*x) % length
		if t < 0 {
			t += length
		}
		if f(m[t]) {
			return t, m[t], true
		}
	}
	return 0, nil, false
}

func (m MyMap2) rehash() Map {
	length := len(m)
	m = append(m, make(MyMap2, 2*length)...)
	for i := 0; i < length; i++ {
		if kv := m[i]; kv != nil {
			m.delValue(i, kv.k)
			m.setValue(i, kv)
		}
	}
	return m
}

type Node struct {
	kv   *KV
	next *Node
}

type MyMap3 []*Node

func (m MyMap3) new(cap int) Map {
	p := make(MyMap3, cap)
	return p
}

func (m MyMap3) operation(k string, f func(kv *KV) bool) (int, *KV, bool) {
	length := len(m)
	index := toHash(k) % length
	if f(nil) {
		return index, nil, true
	}
	for p := m[index]; p != nil; p = p.next {
		if f(p.kv) {
			return index, p.kv, true
		}
	}
	return 0, nil, false
}

func (m MyMap3) setValue(idx int, value *KV) {
	if m[idx] == nil {
		m[idx] = &Node{kv: value}
	} else {
		m[idx].next = &Node{
			kv:   value,
			next: m[idx].next,
		}
	}
}

func (m MyMap3) delValue(idx int, k string) {
	if p := m[idx]; p.kv.k == k {
		m[idx] = m[idx].next
	} else {
		for ; p.next != nil; p = p.next {
			if p.next.kv.k == k {
				p.next = p.next.next
				break
			}
		}
	}
}

func toHash(s string) (sum int) {
	for i := range s {
		sum += int(s[i])
		sum %= math.MaxInt64 / 2
	}
	return
}

func (m MyMap3) rehash() Map {
	length := len(m)
	m = append(m, make(MyMap3, 2*length)...)
	for i := 0; i < length; i++ {
		for p := m[i]; p != nil; p = p.next {
			m.delValue(i, p.kv.k)
			m.setValue(i, p.kv)
		}
	}
	return m
}
