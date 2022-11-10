package treenums_test

import (
	"testing"

	"github.com/0RAJA/Rutils/struct/treenums"
	"github.com/stretchr/testify/require"
)

func TestTreenums(t *testing.T) {
	numArray := Constructor([]int{7, 2, 7, 2, 0})
	require.Equal(t, numArray.SumRange(0, 4), 18)
	numArray.Update(4, 6) // 7, 2, 7, 2, 6
	numArray.Update(0, 2) // 2, 2, 7, 2, 6
	numArray.Update(0, 9) // 9, 2, 7, 2, 6
	require.Equal(t, numArray.SumRange(4, 4), 6)
	numArray.Update(3, 8) // 9, 2, 7, 8, 6
	require.Equal(t, numArray.SumRange(0, 4), 32)
	numArray.Update(4, 1)
	require.Equal(t, numArray.SumRange(0, 3), 26)
	require.Equal(t, numArray.SumRange(0, 4), 27)
}

type NumArray struct {
	*treenums.TreeNum
	nums []int
}

/*
给你一个数组 nums ，请你完成两类查询。

其中一类查询要求 更新 数组num下标对应的值
另一类查询要求返回数组nums中索引left和索引right之间（包含）的nums元素的和，其中left <= right
*/
func Constructor(nums []int) NumArray {
	treeNum := treenums.NewTreeNum(len(nums))
	for i := range nums {
		treeNum.Add(i, nums[i])
	}
	return NumArray{
		TreeNum: treeNum,
		nums:    nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	a := this.nums[index]
	this.Add(index, val-a)
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Ask(right) - this.Ask(left-1)
}
