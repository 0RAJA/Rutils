package RMemory

import (
	"reflect"
)

func SizeOfMemoryInt(value interface{}) int {
	if value == nil {
		return 0
	}
	return int(reflect.TypeOf(value).Size())
}
