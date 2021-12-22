package avl

// AvlTree AVL树
type AvlTree struct {
	Root *AvlTreeNode // 树根节点
}

// KV 节点
type KV struct {
	k, v string
}

// AvlTreeNode AVL节点
type AvlTreeNode struct {
	kv     *KV
	Height int          // 该节点作为树根节点，树的高度，方便计算平衡因子
	Left   *AvlTreeNode // 左子树
	Right  *AvlTreeNode // 右字树
}

// NewAVLTree 初始化一个AVL树
func NewAVLTree() *AvlTree {
	return new(AvlTree)
}

// UpdateHeight 更新节点的树高度
func (node *AvlTreeNode) UpdateHeight() {
	if node == nil {
		return
	}

	var leftHeight, rightHeight int = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	// 哪个子树高算哪棵的
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	// 高度加上自己那一层
	node.Height = maxHeight + 1
}

// BalanceFactor 计算平衡因子
func (node *AvlTreeNode) BalanceFactor() int {
	var leftHeight, rightHeight int = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}

/*
在右子树上插上右儿子导致失衡，左旋，转一次。
在左子树上插上左儿子导致失衡，右旋，转一次。
在左子树上插上右儿子导致失衡，先左后右旋，转两次。
在右子树上插上左儿子导致失衡，先右后左旋，转两次。
旋转规律记忆法：单旋和双旋，单旋反方向，双旋同方向。
*/

// RightRotation 单右旋操作
func RightRotation(Root *AvlTreeNode) *AvlTreeNode {
	// 只有Pivot和B，Root位置变了//Root的左孩子变成pivot的右孩子
	Pivot := Root.Left
	B := Pivot.Right
	Pivot.Right = Root
	Root.Left = B

	// 只有Root和Pivot变化了高度
	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

// LeftRotation 单左旋操作
func LeftRotation(Root *AvlTreeNode) *AvlTreeNode {
	// 只有Pivot和B，Root位置变了
	Pivot := Root.Right
	B := Pivot.Left
	Pivot.Left = Root
	Root.Right = B

	// 只有Root和Pivot变化了高度
	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

// LeftRightRotation 先左后右旋操作
func LeftRightRotation(node *AvlTreeNode) *AvlTreeNode {
	node.Left = LeftRotation(node.Left)
	return RightRotation(node)
}

// RightLeftRotation 先右后左旋操作
func RightLeftRotation(node *AvlTreeNode) *AvlTreeNode {
	node.Right = RightRotation(node.Right)
	return LeftRotation(node)
}

func (tree *AvlTree) Add(k, v string) {
	tree.Root = tree.Root.add(&KV{k: k, v: v})
}

func (node *AvlTreeNode) add(kv *KV) *AvlTreeNode {
	if node == nil {
		return &AvlTreeNode{
			kv:     kv,
			Height: 1,
		}
	}
	if node.kv.k == kv.k {
		node.kv = kv
		return node
	}
	var newTreeNode *AvlTreeNode
	if kv.k > node.kv.k {
		node.Right = node.Right.add(kv)
		factor := node.BalanceFactor()
		if factor <= -2 { //右边变高了
			if kv.k < node.Right.kv.k { //在左边
				newTreeNode = RightLeftRotation(node)
			} else {
				newTreeNode = LeftRotation(node)
			}
		}
	} else {
		node.Left = node.Left.add(kv)
		factor := node.BalanceFactor()
		if factor >= 2 { //右边变高了
			if kv.k > node.Left.kv.k { //在右边
				newTreeNode = LeftRightRotation(node)
			} else {
				newTreeNode = RightRotation(node)
			}
		}
	}
	if newTreeNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newTreeNode.UpdateHeight()
		return newTreeNode
	}
}

// FindMinValue 找出最小值的节点
func (tree *AvlTree) FindMinValue() *AvlTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.findMinValue()
}

func (node *AvlTreeNode) findMinValue() *AvlTreeNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}

	// 一直左子树递归
	return node.Left.findMinValue()
}

// FindMaxValue 找出最大值的节点
func (tree *AvlTree) FindMaxValue() *AvlTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.findMaxValue()
}

func (node *AvlTreeNode) findMaxValue() *AvlTreeNode {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}

	// 一直右子树递归
	return node.Right.findMaxValue()
}

// Find 查找指定节点
func (tree *AvlTree) Find(k string) (string, bool) {
	if tree.Root == nil {
		// 如果是空树，返回空
		return "", false
	}

	return tree.Root.find(k)
}

func (node *AvlTreeNode) find(k string) (string, bool) {
	if node.kv.k == k {
		// 如果该节点刚刚等于该值，那么返回该节点
		return node.kv.v, true
	} else if k < node.kv.k && node.Left != nil { //如果查找的值小于节点值，从节点的左子树开始找,左子树为空，表示找不到该值了，返回nil
		return node.Left.find(k)
	} else if node.Right != nil { //如果查找的值大于节点值，从节点的右子树开始找,右子树为空，表示找不到该值了，返回nil
		return node.Right.find(k)
	}
	return "", false
}

// MidOrder 中序遍历
func (tree *AvlTree) MidOrder() {
	tree.Root.midOrder()
}

func (node *AvlTreeNode) midOrder() {
	if node == nil {
		return
	}
	node.Left.midOrder()  // 先打印左子树
	node.Right.midOrder() // 打印右子树
}

/*
删除
1. 删除的节点是叶子节点，没有儿子，直接删除后看离它最近的父亲节点是否失衡，做调整操作。
2. 删除的节点下有两个子树，选择高度更高的子树下的节点来替换被删除的节点，如果左子树更
	高，选择左子树中最大的节点，也就是左子树最右边的叶子节点，如果右子树更高，选择右子树
	中最小的节点，也就是右子树最左边的叶子节点。最后，删除这个叶子节点，也就是变成情况1。
3. 删除的节点只有左子树，可以知道左子树其实就只有一个节点，被删除节点本身（假设左子树多
	于2个节点，那么高度差就等于2了，不符合AVL树定义），将左节点替换被删除的节点，最后删
	除这个左节点，变成情况1。
4. 删除的节点只有右子树，可以知道右子树其实就只有一个节点，被删除节点本身（假设右子树多
	于2个节点，那么高度差就等于2了，不符合AVL树定义），将右节点替换被删除的节点，最后删
	除这个右节点，变成情况1。
	后面三种情况最后都变成 情况1 ，就是将删除的节点变成叶子节点，然后可以直接删除该叶子节
	点，然后看其最近的父亲节点是否失衡，失衡时对树进行平衡。
*/

func (tree *AvlTree) Del(k string) {
	tree.Root = tree.Root.delete(k)
}

func (node *AvlTreeNode) delete(k string) *AvlTreeNode {
	if node == nil {
		return nil
	}
	if k < node.kv.k {
		node.Left = node.Left.delete(k)
		node.Left.UpdateHeight()
	} else if k > node.kv.k {
		node.Right = node.Right.delete(k)
		node.Right.UpdateHeight()
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		}
		//该节点有两棵子树，选择更高的哪个来替换
		//第二种情况，删除的节点下有两个子树，选择高度更高的子树下的节点来替换被删除的节点，
		//如果左子树更高，选择左子树中最大的节点，也就是左子树最右边的叶子节点，如果右子树更高，选择
		//右子树中最小的节点，也就是右子树最左边的叶子节点。最后，删除这个叶子节点。
		if node.Left != nil && node.Right != nil {
			if node.Left.Height > node.Right.Height {
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}
				//替换节点
				node.kv = maxNode.kv
				//删除节点
				node.Left = node.Left.delete(maxNode.kv.k)
				//更新高度
				node.Left.UpdateHeight()
			} else {
				minNode := node.Right
				for minNode.Left != nil {
					minNode = minNode.Left
				}
				node.kv = minNode.kv
				node.Right = node.Right.delete(minNode.kv.k)
				node.Right.UpdateHeight()
			}
		} else { // 只有左子树或只有右子树
			if node.Left != nil {
				//只有左子树,则该子树只有一个节点
				node.kv = node.Left.kv
				node.Height = 1
				node.Left = nil
			} else if node.Right != nil {
				node.kv = node.Right.kv
				node.Height = 1
				node.Right = nil
			}
		}
		return node
	}
	// 左右子树递归删除节点后需要平衡
	var newNode *AvlTreeNode
	if node.BalanceFactor() >= 2 { //左高
		if node.Left.BalanceFactor() >= 0 { //左左
			newNode = RightRotation(node)
		} else { //左右
			newNode = LeftRightRotation(node)
		}
	} else if node.BalanceFactor() <= -2 { //右边变高了
		if node.Right.BalanceFactor() <= 0 { //在左边
			newNode = RightLeftRotation(node)
		} else {
			newNode = LeftRotation(node)
		}
	}
	if newNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newNode.UpdateHeight()
		return newNode
	}
}
