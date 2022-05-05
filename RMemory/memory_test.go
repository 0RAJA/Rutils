package RMemory

import (
	"fmt"
	"testing"
)

func TestSizeOfMemoryInt(t *testing.T) {
	type M struct {
		a int32
		b int64
		c byte
	}
	d := M{}
	fmt.Println(SizeOfMemoryInt(d), SizeOfMemoryInt(d.a), SizeOfMemoryInt(d.b), SizeOfMemoryInt(d.c))
	e := "123"
	f := ""
	fmt.Println(SizeOfMemoryInt(e), SizeOfMemoryInt(f))
	g := []int{1, 2, 3, 4}
	fmt.Println(SizeOfMemoryInt(g))
}
