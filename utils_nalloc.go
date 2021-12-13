package parco

import (
	"reflect"
	"unsafe"
)

func String2Bytes(str string) (bts []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bts))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}

func Bytes2String(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
