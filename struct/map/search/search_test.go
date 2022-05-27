package search_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/0RAJA/Rutils/struct/map/search"
)

func TestBSTree_Add(t *testing.T) {
	bSTree := search.NewBSTree()
	for i := 0; i < 100; i++ {
		bSTree.Add(strconv.Itoa(i), strconv.Itoa(i+10))
	}
	bSTree.Del("0")
	for i := 0; i < 100; i++ {
		fmt.Println(bSTree.Find(strconv.Itoa(i)))
	}
}
