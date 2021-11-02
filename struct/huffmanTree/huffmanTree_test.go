package huffmanTree

import (
	"fmt"
	"testing"
)

func TestNewHFMTree(t *testing.T) {
	str := "1223334444"
	tree := NewHFMTree(str)
	code := tree.ToCode(str)
	fmt.Println(code)
	fmt.Println(tree.ReCode(code))
}
