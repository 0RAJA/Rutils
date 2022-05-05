package BTree

import (
	"fmt"
	"testing"
)

func TestNewProTree(t *testing.T) {
	tree := NewProTree("ABC##DE#G##F###", '#')
	tree.Count()
	fmt.Println(tree.count.height)
}
