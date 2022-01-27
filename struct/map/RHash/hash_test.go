package RHash_test

import (
	"fmt"
	"github.com/0RAJA/Rutils/struct/map/RHash"
	"testing"
)

func TestNewM(t *testing.T) {
	m := RHash.NewM(RHash.MyMap1{}, 100)
	m.Set("1", "10")
	fmt.Println(m.Get("1"))
}
