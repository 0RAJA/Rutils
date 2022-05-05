package set

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet(-1)
	set.Add(1)
	set.Add(2)
	set.Add(2)
	fmt.Println(set.Len())
	fmt.Println(set.Has(3))
	fmt.Println(set.Has(2))
}
