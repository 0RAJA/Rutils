package link

import "testing"

func TestInitRing(T *testing.T) {
	r := InitRing()
	r.InsertByTail(1, 2, 3, 4, 5)
	r.Print()
}
