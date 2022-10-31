package ssort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort1(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	BubbleSort(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, num)
}

func TestSelectSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	SelectSort(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, num)
}

func TestInsertSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	InsertSort(num, func(i, j int) bool {
		return i < j
	})
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, num)
}

func TestShellSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	ShellSort(num, func(i, j int) bool {
		return i < j
	})
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, num)
}

func TestMergeSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	MergeSort(num, 0, len(num))
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, num)
}

func TestMergeSort2(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	MergeSort2(num, 0, len(num))
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, num)
}

func TestQSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	QSort(num, 0, len(num)-1, func(i, j int) bool {
		return i > j
	})
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, num)
}

func TestQSortState(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	QSortStable(num, true)
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, num)
}

func TestSortList(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	head := &ListNode{}
	tail := head
	for i := range num {
		tail.Next = &ListNode{Val: num[i]}
		tail = tail.Next
	}
	head = SortList(head.Next, nil)
	ret := make([]int, 0, len(num))
	for ; head != nil; head = head.Next {
		ret = append(ret, head.Val)
	}
	target := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	require.EqualValues(t, target, ret)
}
