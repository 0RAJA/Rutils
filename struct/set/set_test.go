package set

import (
	"fmt"
	"testing"
)

var set *Set

func testNewSet(t *testing.T) {
	set = NewSet(-1)
}

func testSetAdd(t *testing.T) {
	set.Add(1)
	set.Add(2)
	set.Add(2)
}

func testSetList(t *testing.T) {
	fmt.Println(set.List())
}

func TestNewSet(t *testing.T) {
	t.Run("", testNewSet)
	t.Run("", testSetAdd)
	t.Run("", testSetList)
}
