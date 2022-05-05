package BSTree

type BSTree struct {
	root *BSNode
}

func NewBSTree() *BSTree {
	return new(BSTree)
}

// Comparable 比较接口小于,返回负值,大于返回正值,等于返回0
type Comparable interface {
	CompareTo(Comparable) int
}

type KV struct {
	key  Comparable
	data interface{}
}
type BSNode struct {
	kv          KV
	size        int //子树个数
	left, right *BSNode
}

func newBSNode(kv KV) *BSNode {
	return &BSNode{kv: kv, size: 1}
}

func (root *BSNode) getSize() int {
	if root == nil {
		return 0
	}
	return root.size
}

func (tree *BSTree) Put(key Comparable, data interface{}) interface{} {
	var ret interface{}
	tree.root, ret = tree.root.put(KV{key: key, data: data})
	return ret
}

func (root *BSNode) put(kv KV) (*BSNode, interface{}) {
	if root == nil {
		return newBSNode(kv), nil
	}
	var oldData interface{}
	cmp := kv.key.CompareTo(root.kv.key)
	if cmp == 0 {
		if root.kv.data != kv.data {
			oldData = root.kv.data
			root.kv.data = kv.data
			return root, oldData
		} else {
			return root, nil
		}
	} else if cmp < 0 {
		root.left, oldData = root.left.put(kv)
	} else {
		root.right, oldData = root.right.put(kv)
	}
	root.size = root.left.getSize() + root.right.getSize() + 1
	return root, oldData
}

func (tree *BSTree) Get(key Comparable) (interface{}, bool) {
	return tree.root.get(key)
}

func (root *BSNode) get(key Comparable) (interface{}, bool) {
	if root == nil {
		return nil, false
	}
	cmp := key.CompareTo(root.kv.key)
	if cmp == 0 {
		return root.kv.data, true
	} else if cmp < 0 {
		return root.left.get(key)
	} else {
		return root.right.get(key)
	}
}

func (tree *BSTree) Delete(key Comparable) {
	tree.root = tree.root.delete(key)
}
func (root *BSNode) getMax() *BSNode {
	if p := root; p != nil {
		for p.right != nil {
			p = p.right
		}
		return p
	}
	return nil
}

func (root *BSNode) getMin() *BSNode {
	if p := root; p != nil {
		for p.left != nil {
			p = p.left
		}
		return p
	}
	return nil
}

func (root *BSNode) delMin() *BSNode {
	if root == nil {
		return nil
	}
	if root.left == nil {
		return root.right
	}
	root.left = root.left.delMin()
	root.size = root.left.getSize() + root.right.getSize() + 1
	return root
}

func (root *BSNode) delMax() *BSNode {
	if root == nil {
		return nil
	}
	if root.right == nil {
		return root.left
	}
	root.right = root.right.delMin()
	root.size = root.left.getSize() + root.right.getSize() + 1
	return root
}

func (root *BSNode) delete(key Comparable) *BSNode {
	if root == nil {
		return nil
	}
	cmp := key.CompareTo(root.kv.key)
	if cmp == 0 {
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		}
		t := root.left.getMax()
		root.kv = t.kv
		root.left = root.left.delMax()
	} else if cmp < 0 {
		root.left = root.left.delete(key)
	} else {
		root.right = root.right.delete(key)
	}
	root.size = root.left.getSize() + root.right.getSize() + 1
	return root
}
