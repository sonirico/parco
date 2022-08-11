package parco

import (
	"bytes"
	"fmt"
	"math"
)

func MaxByteSize64(byteLength int) int64 {
	switch byteLength {
	case 1:
		return math.MaxUint8
	case 2:
		return math.MaxUint16
	case 4:
		return math.MaxUint32
	case 8:
		return math.MaxInt64
	default:
		return (1 << (byteLength / 8)) - 1
	}
}

func MaxSize(byteLength int) int {
	return int(MaxByteSize64(byteLength))
}

func FormatBytes(data []byte) string {
	var buf bytes.Buffer
	buf.WriteString("[]byte{")
	if len(data) > 0 {
		buf.WriteString(fmt.Sprintf("%d", data[0]))
	}
	for i := 1; i < len(data); i++ {
		buf.WriteString(fmt.Sprintf(", %d", data[i]))
	}
	buf.WriteByte('}')
	return buf.String()
}
