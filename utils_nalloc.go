package parco

import (
	"unsafe"
)

func String2Bytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func Bytes2String(data []byte) string {
	return unsafe.String(unsafe.SliceData(data), len(data))
}
