package LLRBTree

import "fmt"

const (
	RED   = true
	BLACK = false
)

// LLRBTree 左倾红黑树
type LLRBTree struct {
	Root *LLRBTNode
}

// LLRBTNode 左倾红黑树节点
type LLRBTNode struct {
	Value int64      // 值
	Times int64      // 值出现的次数
	Left  *LLRBTNode // 左子树
	Right *LLRBTNode // 右子树
	Color bool       // 父亲指向该节点的链接颜色
}

func NewLLRBTree() *LLRBTree {
	return &LLRBTree{}
}

// IsRed 节点颜色
func (node *LLRBTNode) IsRed() bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}

// RotateLeft 左旋
func (node *LLRBTNode) RotateLeft() *LLRBTNode {
	if node == nil {
		return nil
	}
	r := node.Right
	node.Right = r.Left
	r.Left = node
	r.Color = node.Color
	node.Color = RED
	return r
}

// RotateRight 右旋
func (node *LLRBTNode) RotateRight() *LLRBTNode {
	if node == nil {
		return nil
	}
	l := node.Left
	node.Left = l.Right
	l.Right = node
	l.Color = node.Color
	node.Color = RED
	return l
}

// ColorChange 调整颜色
func (node *LLRBTNode) ColorChange() {
	if node == nil {
		return
	}
	node.Color = !node.Color
	node.Left.Color = !node.Left.Color
	node.Right.Color = !node.Right.Color
}

func (tree *LLRBTree) Add(value int64) {
	tree.Root = tree.Root.add(value)
	//根节点连接永远为黑
	tree.Root.Color = BLACK
}

//添加
func (node *LLRBTNode) add(value int64) *LLRBTNode {
	if node == nil {
		return &LLRBTNode{
			Value: value,
			Times: 1,
			Color: RED,
		}
	}
	if value == node.Value {
		node.Times++
	} else if value > node.Value {
		node.Right = node.Right.add(value)
	} else {
		node.Left = node.Left.add(value)
	}
	nowNode := node
	if node.Right.Color == RED && !node.Left.Color {
		nowNode = node.RotateLeft()
	} else {
		if node.Left.IsRed() && node.Left.Left.IsRed() {
			nowNode = node.RotateRight()
		}
		if node.Left.IsRed() && node.Right.IsRed() {
			node.ColorChange()
		}
	}
	return nowNode
}

// FindMinValue 找出最小值的节点
func (tree *LLRBTree) FindMinValue() *LLRBTNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.FindMinValue()
}

func (node *LLRBTNode) FindMinValue() *LLRBTNode {
	// 左子树为空，表面已经是最左的节点了，该值就是最小值
	if node.Left == nil {
		return node
	}

	// 一直左子树递归
	return node.Left.FindMinValue()
}

// FindMaxValue 找出最大值的节点
func (tree *LLRBTree) FindMaxValue() *LLRBTNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.FindMaxValue()
}

func (node *LLRBTNode) FindMaxValue() *LLRBTNode {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}

	// 一直右子树递归
	return node.Right.FindMaxValue()
}

// Find 查找指定节点
func (tree *LLRBTree) Find(value int64) *LLRBTNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}

	return tree.Root.Find(value)
}

func (node *LLRBTNode) Find(value int64) *LLRBTNode {
	if value == node.Value {
		// 如果该节点刚刚等于该值，那么返回该节点
		return node
	} else if value < node.Value {
		// 如果查找的值小于节点值，从节点的左子树开始找
		if node.Left == nil {
			// 左子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Left.Find(value)
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		if node.Right == nil {
			// 右子树为空，表示找不到该值了，返回nil
			return nil
		}
		return node.Right.Find(value)
	}
}

// MidOrder 中序遍历
func (tree *LLRBTree) MidOrder() {
	tree.Root.MidOrder()
}

func (node *LLRBTNode) MidOrder() {
	if node == nil {
		return
	}

	// 先打印左子树
	node.Left.MidOrder()

	// 按照次数打印根节点
	for i := 1; i <= int(node.Times); i++ {
		fmt.Println(node.Value)
	}
	// 打印右子树
	node.Right.MidOrder()
}

/*
删除操作就复杂得多了。对照一下 2-3 树。
1. 情况1：如果删除的是非叶子节点，找到其最小后驱节点，也就是在其右子树中一直向左找，
	找到的该叶子节点替换被删除的节点，然后删除该叶子节点，变成情况2。
2. 情况2：如果删除的是叶子节点，如果它是红节点，也就是父亲指向它的链接为红色，那么直接删除即可。
	否则，我们需要进行调整，使它变为红节点，再删除。
在这里，为了使得删除叶子节点时可以直接删除，叶子节点必须变为红节点。（在 2-3 树中，也
就是2节点要变成3节点，我们知道要不和父亲合并再递归向上，要不向兄弟借值然后重新分布）
我们创造两种操作，如果删除的节点在左子树中，可能需要进行红色左移，如果删除的节点在右子树
中，可能需要进行红色右移。
*/

// MoveRedLeft 红色左移
// 节点 h 是红节点，其左儿子和左儿子的左儿子都为黑节点，左移后使得其左儿子或左儿子的左儿子有
//一个是红色节点
func MoveRedLeft(h *LLRBTNode) *LLRBTNode {
	// 应该确保 isRed(h) && !isRed(h.left) && !isRed(h.left.left)
	if h.IsRed() && !h.Left.IsRed() && !h.Left.Left.IsRed() {
		h.ColorChange()
		if h.Right.Left.IsRed() {
			//先对h.right右旋
			h.Right = h.Right.RotateRight()
			//对h左旋
			h = h.RotateLeft()
			h.ColorChange()
		}
	}
	return h
}

// MoveRedRight /*
func MoveRedRight(h *LLRBTNode) *LLRBTNode {
	// 应该确保 isRed(h) && !isRed(h.right) && !isRed(h.right.left);
	if h.IsRed() && !h.Right.IsRed() && !h.Right.Left.IsRed() {
		h.ColorChange()
		//左儿子有左红链接
		if h.Left.Left.IsRed() {
			h = h.RotateRight()
			h.ColorChange()
		}
	}
	return h
}

// FixUp 恢复左倾红黑树特征
func (node *LLRBTNode) FixUp() *LLRBTNode {
	nowNode := node
	//红色在右边,左旋
	if nowNode.Right.IsRed() {
		nowNode = nowNode.RotateLeft()
	}
	//连续两个左链接都是红色需要右旋
	if nowNode.Left.IsRed() && nowNode.Left.Left.IsRed() {
		nowNode = nowNode.RotateRight()
	}
	//左右都是红色,需要变色
	if nowNode.Left.IsRed() && nowNode.Right.IsRed() {
		nowNode.ColorChange()
	}
	return nowNode
}

// Delete 左倾红黑树删除元素
func (tree *LLRBTree) Delete(value int64) {
	// 当找不到值时直接返回
	if tree.Find(value) == nil {
		return
	}

	if !tree.Root.Left.IsRed() && !tree.Root.Right.IsRed() {
		// 左右子树都是黑节点，那么先将根节点变为红节点，方便后面的红色左移或右移
		tree.Root.Color = RED
	}

	tree.Root = tree.Root.Delete(value)

	// 最后，如果根节点非空，永远都要为黑节点，赋值黑色
	if tree.Root != nil {
		tree.Root.Color = BLACK
	}
}

// Delete 对该节点所在的子树删除元素
func (node *LLRBTNode) Delete(value int64) *LLRBTNode {
	// 辅助变量
	nowNode := node
	// 删除的元素比子树根节点小，需要从左子树删除
	if value < nowNode.Value {
		// 因为从左子树删除，所以要判断是否需要红色左移
		if !nowNode.Left.IsRed() && !nowNode.Left.Left.IsRed() {
			// 左儿子和左儿子的左儿子都不是红色节点，那么没法递归下去，先红色左移
			nowNode = MoveRedLeft(nowNode)
		}

		// 现在可以从左子树中删除了
		nowNode.Left = nowNode.Left.Delete(value)
	} else {
		// 删除的元素等于或大于树根节点

		// 左节点为红色，那么需要右旋，方便后面可以红色右移
		if nowNode.Left.IsRed() {
			nowNode = nowNode.RotateRight()
		}

		// 值相等，且没有右孩子节点，那么该节点一定是要被删除的叶子节点，直接删除
		// 为什么呢，反证，它没有右儿子，但有左儿子，因为左倾红黑树的特征，那么左儿子一定是红色，但是前面的语句已经把红色左儿子右旋到右边，不应该出现右儿子为空。
		if value == nowNode.Value && nowNode.Right == nil {
			return nil
		}

		// 因为从右子树删除，所以要判断是否需要红色右移
		if !nowNode.Right.IsRed() && !nowNode.Right.Left.IsRed() {
			// 右儿子和右儿子的左儿子都不是红色节点，那么没法递归下去，先红色右移
			nowNode = MoveRedRight(nowNode)
		}

		// 删除的节点找到了，它是中间节点，需要用最小后驱节点来替换它，然后删除最小后驱节点
		if value == nowNode.Value {
			minNode := nowNode.Right.FindMinValue()
			nowNode.Value = minNode.Value
			nowNode.Times = minNode.Times

			// 删除其最小后驱节点
			nowNode.Right = nowNode.Right.DeleteMin()
		} else {
			// 删除的元素比子树根节点大，需要从右子树删除
			nowNode.Right = nowNode.Right.Delete(value)
		}
	}

	// 最后，删除叶子节点后，需要恢复左倾红黑树特征
	return nowNode.FixUp()
}

// DeleteMin 对该节点所在的子树删除最小元素
func (node *LLRBTNode) DeleteMin() *LLRBTNode {
	// 辅助变量
	nowNode := node

	// 没有左子树，那么删除它自己
	if nowNode.Left == nil {
		return nil
	}

	// 判断是否需要红色左移，因为最小元素在左子树中
	if nowNode.Left.IsRed() && nowNode.Left.Left.IsRed() {
		nowNode = MoveRedLeft(nowNode)
	}

	// 递归从左子树删除
	nowNode.Left = nowNode.Left.DeleteMin()

	// 修复左倾红黑树特征
	return nowNode.FixUp()
}
