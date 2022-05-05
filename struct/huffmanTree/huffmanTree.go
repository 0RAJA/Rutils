package huffmanTree

import (
	"math"
	"sort"
)

type HFMNode struct {
	weight                 int
	parent, lChild, rChild int
	c                      rune
}

type HFMTree struct {
	tree   []*HFMNode
	code   map[rune]string
	weight map[rune]int
	rs     []rune
}

func searchMin(list []*HFMNode, length int) (int, int) {
	min1, min2 := math.MaxInt32, math.MaxInt32
	var index1, index2 int
	for i := 1; i <= length; i++ {
		if list[i].parent != 0 {
			continue
		}
		if min2 > list[i].weight {
			t1, t2 := list[i].weight, i
			if min1 > list[i].weight {
				t1, t2 = min1, index1
				min1, index1 = list[i].weight, i
			}
			min2, index2 = t1, t2
		}
	}
	return index1, index2
}

func (t *HFMTree) searchCode(c rune) (ret string) {
	index := 0
	for i := 1; i < len(t.tree); i++ {
		v := t.tree[i]
		if v.c == c {
			index = i
			break
		}
	}
	for t.tree[index].parent != 0 {
		p := t.tree[index].parent
		if t.tree[p].lChild == index {
			ret = "0" + ret
		} else {
			ret = "1" + ret
		}
		index = p
	}
	return
}

func NewHFMTree(str string) (tree *HFMTree) {
	weight := make(map[rune]int)
	rs := make([]rune, 0)
	for _, v := range str {
		if weight[v] == 0 {
			rs = append(rs, v)
		}
		weight[v]++
	}
	return NewHFMTreeWithWright(weight, rs)
}

func NewHFMTreeWithWright(weight map[rune]int, rs []rune) (tree *HFMTree) {
	tree = &HFMTree{code: map[rune]string{}, weight: weight, rs: rs}
	sort.Slice(tree.rs, func(i, j int) bool {
		return tree.rs[i] < tree.rs[j]
	})
	tree.tree = make([]*HFMNode, 2*len(tree.weight))
	m := 1
	for _, c := range tree.rs {
		tree.tree[m] = &HFMNode{weight: tree.weight[c], c: c}
		m++
	}
	for i := len(tree.weight) + 1; i < len(tree.tree); i++ {
		min1, min2 := searchMin(tree.tree, i-1)
		if tree.tree[min1].weight == tree.tree[min2].weight {
			if min2 < min1 {
				min2, min1 = min1, min2
			}
		}
		tree.tree[i] = &HFMNode{
			weight: tree.tree[min1].weight + tree.tree[min2].weight,
			lChild: min1,
			rChild: min2,
		}
		tree.tree[min1].parent = i
		tree.tree[min2].parent = i
	}
	for c := range tree.weight {
		code := tree.searchCode(c)
		tree.code[c] = code
	}
	return tree
}

func (t *HFMTree) ToCode(str string) (ret string) {
	for _, v := range str {
		ret += t.code[v]
	}
	return
}

func (t *HFMTree) ReCode(str string) (ret string) {
	p := t.tree[len(t.tree)-1]
	for _, v := range str {
		if v == '0' {
			p = t.tree[p.lChild]
		} else {
			p = t.tree[p.rChild]
		}
		if p.lChild == 0 && p.rChild == 0 {
			ret += string(p.c)
			p = t.tree[len(t.tree)-1]
		}
	}
	return
}
