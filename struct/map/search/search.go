package search

type KV struct {
	k, v string
}

type BSTree struct {
	root *BSNode
}

type BSNode struct {
	kv           *KV
	lNode, rNode *BSNode
}

func NewBSTree() *BSTree {
	return new(BSTree)
}

func (tree *BSTree) Add(k, v string) {
	tree.root = tree.root.add(&KV{k: k, v: v})
}

func (root *BSNode) add(kv *KV) *BSNode {
	if root == nil {
		return &BSNode{kv: kv}
	} else if root.kv.k == kv.k {
		root.kv = kv
	} else if kv.k < root.kv.k {
		root.lNode = root.lNode.add(kv)
	} else {
		root.rNode = root.rNode.add(kv)
	}
	return root
}
func (tree *BSTree) Find(k string) (string, bool) {
	return tree.root.find(k)
}
func (root *BSNode) find(k string) (string, bool) {
	if root == nil {
		return "", false
	}
	if root.kv.k == k {
		return root.kv.v, true
	} else if root.kv.k > k {
		return root.lNode.find(k)
	} else {
		return root.rNode.find(k)
	}
}

func (tree *BSTree) Del(k string) {
	tree.root = tree.root.del(k)
}

func (root *BSNode) del(k string) *BSNode {
	if root != nil {
		if k < root.kv.k {
			root.lNode = root.lNode.del(k)
		} else if k > root.kv.k {
			root.rNode = root.rNode.del(k)
		} else {
			if root.lNode != nil && root.rNode != nil { //两个孩子都存在的情况
				t := root.rNode.findMin()
				root.kv = t.kv
				root.rNode = root.rNode.del(t.kv.k)
			} else { //叶子节点或者只有一个孩子的情况
				if root.lNode != nil {
					root = root.lNode
				} else if root.rNode != nil {
					root = root.rNode
				} else {
					root = nil
				}
			}
		}
	}
	return root
}

func (root *BSNode) findMin() (ret *BSNode) {
	if ret = root; ret != nil {
		for ret.lNode != nil {
			ret = ret.lNode
		}
	}
	return ret
}
func (root *BSNode) findMax() (ret *BSNode) {
	if ret = root; ret != nil {
		for ret.rNode != nil {
			ret = ret.rNode
		}
	}
	return ret
}
