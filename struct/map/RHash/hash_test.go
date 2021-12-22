package RHash_test

import (
	"Rutils/struct/map/RHash"
	"fmt"
	"testing"
)

func TestNewM(t *testing.T) {
	m := RHash.NewM(RHash.MyMap1{}, 100)
	m.Set("1", "10")
	fmt.Println(m.Get("1"))
}
