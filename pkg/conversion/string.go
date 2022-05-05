package conversion

import (
	"reflect"
	"unsafe"
)

func UnsafeStringToBytes(str string) []byte {
	p := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	b := reflect.SliceHeader{
		Data: p.Data,
		Len:  p.Len,
		Cap:  p.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func UnsafeBytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
