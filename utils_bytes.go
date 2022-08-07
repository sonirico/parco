package parco

import "math"

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
