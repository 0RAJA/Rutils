package PrefixTree

type PrefixTree struct {
	suffix  map[byte]*PrefixTree
	results []string
}

// NewPrefixTree 生成前缀树
func NewPrefixTree() *PrefixTree {
	return &PrefixTree{suffix: map[byte]*PrefixTree{}}
}

func (t *PrefixTree) Insert(k, v string) {
	root := t
	for i := 0; i < len(k); i++ {
		if root.suffix[k[i]] == nil {
			root.suffix[k[i]] = &PrefixTree{suffix: make(map[byte]*PrefixTree)}
		}
		root = root.suffix[k[i]]
	}
	root.results = append(root.results, v)
}

func (t *PrefixTree) Search(k string) []string {
	root := t
	for i := 0; i < len(k); i++ {
		if root.suffix[k[i]] != nil {
			root = root.suffix[k[i]]
		} else {
			return []string{}
		}
	}
	return root.results
}
