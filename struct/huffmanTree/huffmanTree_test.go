package huffmanTree

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestNewHFMTree(t *testing.T) {
	buf := bufio.NewReader(os.Stdin)
	str, _ := buf.ReadString('#')
	str = str[:len(str)-1]
	_, _ = buf.ReadString('\n')
	tree := NewHFMTree(str)
	fmt.Println(tree.ToCode(str))
	str, _ = buf.ReadString('\n')
	fmt.Println(tree.ReCode(str))
	var nums = 0
	for _, v := range tree.weight {
		nums += v
	}
	var sum float64 = 0
	for k, v := range tree.code {
		sum += (float64(tree.weight[k])) / float64(nums) * float64(len(v))
	}
	fmt.Printf("%.2f", sum)
}
