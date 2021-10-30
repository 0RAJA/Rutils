package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewHeap(t *testing.T) {
	heap := NewHeap()
	for i := 0; i < 20; i++ {
		time.Sleep(1000 * time.Microsecond)
		heap.Push(randNum())
	}
	fmt.Println(heap.Array[1 : heap.Len+1])
	//for i := 0; i < 20; i++ {
	//	fmt.Print(heap.Pop(), " ")
	//}
	heap.HeapSort()
	fmt.Println(heap.Array[1 : heap.Len+1])
}

func randNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000) + 1
}
