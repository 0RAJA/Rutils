package KMP

import (
	"fmt"
	"testing"
)

func TestGetNext(t *testing.T) {
	s := "abcac"
	fmt.Println(GetNext2(s))
}
