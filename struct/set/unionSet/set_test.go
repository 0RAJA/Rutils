package unionSet_test

import (
	"fmt"
	"testing"
)
import "github.com/0RAJA/Rutils/struct/set/unionSet"

func TestNewSet(t *testing.T) {
	s := unionSet.NewSet(5)
	fmt.Println(s.InSameSet(1, 2))
}
