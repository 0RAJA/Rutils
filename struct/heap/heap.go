package heap

import (
	"math"
)

const MaxSize = 1000
const MaxData = math.MaxInt64

// Heap 一个最大堆，一棵完全二叉树
// 最大堆要求节点元素都不小于其左右孩子
type Heap struct {
	// 堆的大小
	Len int
	Cap int
	// 使用内部的数组来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i，右儿子下标为 2i+1
	Array []int
}

// NewHeap 初始化一个空堆,创建哨兵
func NewHeap() *Heap {
	return &Heap{
		Len:   0,
		Cap:   MaxSize,
		Array: append([]int{MaxData}, make([]int, MaxSize)...),
	}
}

// Push 最大堆插入元素
func (h *Heap) Push(x int) {
	// 堆没有元素时，使元素成为顶点后退出
	h.Len++
	if h.Len > h.Cap {
		h.Array = append(h.Array, make([]int, h.Len)...)
		h.Cap += h.Len
	}
	// i 是要插入节点的下标
	i := h.Len
	h.Array[i] = x
	//如果是第一个元素直接插入就完成任务了
	if h.Len == 1 {
		return
	}
	for i > 0 {
		// parent为该元素父亲节点的下标
		parent := i / 2

		// 如果插入的值小于等于父亲节点，那么可以直接退出循环，因为父亲仍然是最大的
		if x <= h.Array[parent] { //有个哨兵
			break
		}

		// 否则将父亲节点与该节点互换，然后向上翻转，将最大的元素一直往上推
		h.Array[i] = h.Array[parent]
		i = parent
	}
}

// Pop 最大堆移除根节点元素，也就是最大的元素
func (h *Heap) Pop() int {
	// 没有元素，返回-1
	if h.Len == 0 {
		return -1
	}

	// 取出根节点
	ret := h.Array[1]

	// 因为根节点要被删除了，将最后一个节点放到根节点的位置上
	h.Array[1] = h.Array[h.Len] // 将最后一个元素的值先拿出来
	h.Len--

	// 对根节点进行向下翻转，小的值 x 一直下沉，维持最大堆的特征
	h.siftDown(1, h.Len)
	return ret
}

//向下调整
func (h *Heap) siftDown(num, size int) {
	t := 0
	for num*2 <= size {
		if h.Array[num] >= h.Array[num*2] {
			t = num
		} else {
			t = num * 2
		}
		if num*2+1 <= size {
			if h.Array[t] < h.Array[num*2+1] {
				t = num*2 + 1
			}
		}
		if t == num {
			break
		} else {
			h.Array[t], h.Array[num] = h.Array[num], h.Array[t]
			num = t
		}
	}
}

func (h *Heap) MaxHeap(size int) {
	h.siftDown(1, size)
}

// HeapSort 堆排
func (h *Heap) HeapSort() {
	h.MaxHeap(h.Len)
	for i := h.Len; i > 1; i-- {
		h.Array[1], h.Array[i] = h.Array[i], h.Array[1]
		h.siftDown(1, i-1)
	}
}
