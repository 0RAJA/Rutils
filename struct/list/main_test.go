package list

import (
	"fmt"
	"testing"
)

func testInit() *DoubleList {
	list := DoubleListInit()
	list.AddNodeFromTail(0, "1")
	list.AddNodeFromTail(1, "12")
	list.AddNodeFromTail(1, "123")
	list.AddNodeFromTail(1, "1234")
	return list
}
func TestDoubleListInit(t *testing.T) {
	list := DoubleListInit()
	fmt.Println(list.len)
}

func TestDoubleList_AddNodeFromHead(t *testing.T) {
	list := DoubleListInit()
	list.AddNodeFromHead(0, "123")
	fmt.Println(list.head, list.tail, list.len)
}

func TestDoubleList_AddNodeFromTail(t *testing.T) {
	list := testInit()
	fmt.Println(list.len)
	fmt.Println(list.PrintHead())
	fmt.Println(list.PrintTail())
}

func TestDoubleList_PopFromHead(t *testing.T) {
	list := testInit()
	fmt.Println(list.PopFromHead(1).GetValue())
	fmt.Println(list.PopFromTail(1).GetValue())
	fmt.Println(list.PopFromTail(1).GetValue())
	fmt.Println(list.PopFromTail(1).GetValue())
	fmt.Println(list.PrintTail())
}

func TestDoubleList_IndexFromHead(t *testing.T) {
	list := testInit()
	fmt.Println(list.IndexFromHead(2).GetValue())
}
