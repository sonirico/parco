package parco

import (
	"encoding/binary"
	"math"
)

func CompileFloat64(f32 float64, box []byte, order binary.ByteOrder) (err error) {
	return CompileUInt64(math.Float64bits(f32), box, order)
}

func ParseFloat64(data []byte, order binary.ByteOrder) (float64, error) {
	if len(data) < 8 {
		return 0, NewErrUnSufficientBytesError(8, 0)
	}

	u32, err := ParseUInt64(data, order)
	if err != nil {
		return 0, err
	}

	return math.Float64frombits(u32), nil
}

func Float64(order binary.ByteOrder) Type[float64] {
	return NewFixedType[float64](
		8,
		func(data []byte) (float64, error) {
			return ParseFloat64(data, order)
		},
		func(value float64, box []byte) (err error) {
			return CompileFloat64(value, box, order)
		},
	)
}

func Float64LE() Type[float64] {
	return Float64(binary.LittleEndian)
}

func Float64BE() Type[float64] {
	return Float64(binary.BigEndian)
}
