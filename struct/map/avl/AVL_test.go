package avl_test

import (
	"fmt"
	"github.com/0RAJA/Rutils/struct/map/avl"
	"strconv"
	"testing"
)

func TestAvlTree_Add(t *testing.T) {
	avlTree := avl.NewAVLTree()
	for i := 0; i < 100; i++ {
		avlTree.Add(strconv.Itoa(i), strconv.Itoa(i+10))
	}
	avlTree.Del("1")
	for i := 0; i < 100; i++ {
		fmt.Println(avlTree.Find(strconv.Itoa(i)))
	}
}
