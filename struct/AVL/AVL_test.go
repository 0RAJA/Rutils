package AVL

import (
	"testing"
)

func TestNewAVL(t *testing.T) {
	avlTree := NewAVLTree()
	avlTree.Add(1)
	avlTree.Add(2)
	avlTree.Add(3)
	avlTree.Add(6)
	avlTree.Add(2)
	avlTree.MidOrder()
}
