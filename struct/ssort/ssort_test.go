package ssort

import (
	"fmt"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	BubbleSort(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	fmt.Println(num)
}

func TestSelectSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	SelectSort(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	fmt.Println(num)
}

func TestInsertSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	InsertSort(num, func(i, j int) bool {
		return i < j
	})
	fmt.Println(num)
}

func TestShellSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	ShellSort(num, func(i, j int) bool {
		return i < j
	})
	fmt.Println(num)
}

func TestMergeSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	MergeSort(num, 0, len(num))
	fmt.Println(num)
}

func TestMergeSort2(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	MergeSort2(num, 0, len(num))
	fmt.Println(num)
	//Output:[1 2 3 4 5 5 6 7 8 9]
}

func TestQSort(t *testing.T) {
	num := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5}
	QSort(num, 0, len(num)-1, func(i, j int) bool {
		return i > j
	})
	fmt.Println(num)
}
