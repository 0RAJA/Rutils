package skip_list_test

import (
	"Rutils/struct/skip_list"
	"fmt"
	"testing"
)

type M int

func (m M) CompareTo(comparable skip_list.Comparable) int {
	c, _ := comparable.(M)
	return int(m - c)
}

func TestNewSkipList(t *testing.T) {
	sList := skip_list.NewSkipList()
	sList.Put(M(1), 111)
	fmt.Println(sList.Get(M(1)))
	sList.Delete(M(1))
	fmt.Println(sList.Get(M(1)))
}
