//copy form fasthttp
package util

import (
	"bytes"
	"encoding/gob"
	"errors"
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
func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// string -> []byte
func String2Byte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

//copy memcache
func StructToByte(value interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err = encoder.Encode(value); err != nil {
		return nil, errors.New("encode fail")
	}
	return buf.Bytes(), nil
}

func ByteToStruct(b []byte, value interface{}) (err error) {
	buf := bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	if err = decoder.Decode(value); err != nil {
		return errors.New("decode fail")
	}
	return nil
}
