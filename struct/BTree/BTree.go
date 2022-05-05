package BTree

type Node struct {
	val   rune
	left  *Node
	right *Node
	isOK  int
}

type Tree struct {
	head  *Node
	count struct {
		zero   []*Node
		one    []*Node
		two    []*Node
		height int
	}
}

type Stack struct {
	stack []*Node
}
type Queue Stack

func NewQueue() *Queue {
	return &Queue{}
}
func (q *Queue) Push(node *Node) {
	q.stack = append(q.stack, node)
}
func (q *Queue) Front() (node *Node) {
	node = q.stack[0]
	q.stack = q.stack[1:]
	return
}
func (q *Queue) IsEmpty() bool {
	return len(q.stack) == 0
}
func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(node *Node) {
	s.stack = append(s.stack, node)
}
func (s *Stack) Pop() (ret *Node) {
	ret = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
func (s *Stack) Top() *Node {
	return s.stack[len(s.stack)-1]
}
func NewProTree(str string, sep rune) (tree *Tree) {
	tree = &Tree{head: nil}
	var newProTree func(node **Node)
	index := 0
	newProTree = func(node **Node) {
		c := rune(str[index])
		index++
		if c == sep {
			*node = nil
		} else {
			*node = &Node{val: c}
			newProTree(&((*node).left))
			newProTree(&((*node).right))
		}
	}
	newProTree(&tree.head)
	return
}

func (t *Tree) Pro() (str string) {
	return pro(t.head)
}

func pro(node *Node) (str string) {
	if node == nil {
		return
	}
	return string(node.val) + pro(node.left) + pro(node.right)
}

func (t *Tree) Mid() (str string) {
	return mid(t.head)
}

func mid(node *Node) (str string) {
	if node == nil {
		return
	}
	return mid(node.left) + string(node.val) + mid(node.right)
}

func (t *Tree) Back() (str string) {
	return back(t.head)
}

func back(node *Node) (str string) {
	if node == nil {
		return
	}
	return back(node.left) + back(node.right) + string(node.val)
}
func (node *Node) Print(flag int) string {
	node.isOK = flag
	return string(node.val)
}
func (node *Node) IsVisit(flag int) bool {
	return node == nil || node.isOK == flag
}
func (t *Tree) NoPro(flag int) (ret string) {
	stack := NewStack()
	stack.Push(t.head)
	for !stack.IsEmpty() {
		p := stack.Pop()
		ret += p.Print(flag)
		if !p.right.IsVisit(flag) {
			stack.Push(p.right)
		}
		if !p.left.IsVisit(flag) {
			stack.Push(p.left)
		}
	}
	return
}
func (t *Tree) NoMid(flag int) (ret string) {
	stack := NewStack()
	stack.Push(t.head)
	for !stack.IsEmpty() {
		p := stack.Top()
		if !p.left.IsVisit(flag) {
			stack.Push(p.left)
		} else {
			stack.Pop()
			ret += p.Print(flag)
			if !p.right.IsVisit(flag) {
				stack.Push(p.right)
			}
		}
	}
	return
}
func (t *Tree) NoBack(flag int) (ret string) {
	stack := NewStack()
	stack.Push(t.head)
	for !stack.IsEmpty() {
		p := stack.Top()
		if !p.left.IsVisit(flag) {
			stack.Push(p.left)
		} else if !p.right.IsVisit(flag) {
			stack.Push(p.right)
		} else {
			stack.Pop()
			ret += p.Print(flag)
		}
	}
	return
}

func (t *Tree) Level() (ret string) {
	queue := NewQueue()
	queue.Push(t.head)
	for !queue.IsEmpty() {
		p := queue.Front()
		ret += string(p.val)
		if p.left != nil {
			queue.Push(p.left)
		}
		if p.right != nil {
			queue.Push(p.right)
		}
	}
	return
}
func (t *Tree) Count() {
	var count func(node *Node) int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	count = func(node *Node) int {
		if node == nil {
			return 0
		}
		if node.left == nil && node.right == nil {
			t.count.zero = append(t.count.zero, node)
		} else if node.left != nil && node.right != nil {
			t.count.two = append(t.count.two, node)
		} else {
			t.count.one = append(t.count.one, node)
		}
		return max(count(node.left), count(node.right)) + 1
	}
	t.count.height = count(t.head)
}
