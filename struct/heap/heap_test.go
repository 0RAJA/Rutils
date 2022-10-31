package heap

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewHeap(t *testing.T) {
	heap := NewHeap()
	nums := make([]int, 20)
	for i := range nums {
		time.Sleep(1000 * time.Microsecond)
		nums[i] = randNum()
	}
	for i := range nums {
		heap.Push(nums[i])
	}
	heap.HeapSort()
	sort.Ints(nums)
	require.EqualValues(t, nums, heap.Array[1:heap.Len+1])
}

func randNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000) + 1
}
