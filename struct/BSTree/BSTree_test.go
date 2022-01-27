package BSTree_test

import (
	"fmt"
	"github.com/0RAJA/Rutils/struct/BSTree"
	"testing"
)

type T1 int

func (t T1) CompareTo(x BSTree.Comparable) int {
	k, _ := x.(T1)
	return int(t - k)
}

func TestNewBSTree(t *testing.T) {
	tree := BSTree.NewBSTree()
	for i := 0; i < 100; i++ {
		tree.Put(T1(i), i)
	}
	tree.Delete(T1(1))
	for i := 0; i < 100; i++ {
		fmt.Println(tree.Get(T1(i)))
	}
}
