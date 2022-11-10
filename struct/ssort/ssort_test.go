package ssort

import (
	"sort"
	"testing"

	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestBubbleSort1(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		sort.Ints(nums)
		BubbleSort(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
		require.EqualValues(t, nums, nums2)
	}
}

func TestSelectSort(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		sort.Ints(nums)
		SelectSort(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
		require.EqualValues(t, nums, nums2)
	}
}

func TestInsertSort(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		sort.Ints(nums)
		InsertSort(nums2, func(num1, num2 int) bool { return num1 < num2 })
		require.EqualValues(t, nums, nums2)
	}
}

func TestShellSort(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		sort.Ints(nums)
		ShellSort(nums2, func(num1, num2 int) bool { return num1 < num2 })
		require.EqualValues(t, nums, nums2)
	}
}

func TestMergeSort(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		MergeSort2(nums, 0, len(nums))
		MergeSort(nums2, 0, len(nums2))
		require.EqualValues(t, nums, nums2)
	}
}

func TestMergeSort2(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		MergeSort(nums, 0, len(nums))
		MergeSort2(nums2, 0, len(nums2))
		require.EqualValues(t, nums, nums2)
	}
}

func TestQSort(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		sort.Ints(nums)
		QSort(nums2, func(num1, num2 int) bool { return num1 < num2 })
		require.EqualValues(t, nums, nums2)
	}
}

func TestQSortState(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		MergeSort(nums, 0, len(nums))
		QSortStable(nums2, true)
		require.EqualValues(t, nums, nums2)
	}
}

func TestSortList(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		head := &ListNode{}
		tail := head
		for i := range nums {
			tail.Next = &ListNode{Val: nums[i]}
			tail = tail.Next
		}
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		sort.Ints(nums2)
		head = SortList(head.Next, nil)
		ret := make([]int, 0, len(nums))
		for ; head != nil; head = head.Next {
			ret = append(ret, head.Val)
		}
		require.EqualValues(t, nums2, ret)
	}
}

func TestSearchMiddle(t *testing.T) {
	for i := 0; i < int(utils.RandomInt(5, 10)); i++ {
		cnt := utils.RandomInt(10, 100)
		if cnt%2 == 0 {
			cnt++
		}
		nums := make([]int, cnt)
		for i := range nums {
			nums[i] = int(utils.RandomInt(-100, 100))
		}
		result := SearchMiddle(nums, 0, len(nums)-1)
		sort.Ints(nums)
		target := nums[len(nums)/2]
		require.Equal(t, target, result, nums)
	}
}
