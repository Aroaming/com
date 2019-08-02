//copy form fasthttp
package util

import (
	"reflect"
	"unsafe"
)

type StringHeader struct {
	Data uintptr
	Len  int
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// []byte -> string
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// string -> []byte
func s2b(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
