package parco

import (
	"encoding/binary"
	"math"
)

func CompileFloat32(f32 float32, box []byte, order binary.ByteOrder) (err error) {
	return CompileUInt32(math.Float32bits(f32), box, order)
}

func ParseFloat32(data []byte, order binary.ByteOrder) (float32, error) {
	if len(data) < 4 {
		return 0, NewErrUnSufficientBytesError(4, 0)
	}

	u32, err := ParseUInt32(data, order)
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(u32), nil
}

func Float32(order binary.ByteOrder) Type[float32] {
	return NewFixedType[float32](
		4,
		func(data []byte) (float32, error) {
			return ParseFloat32(data, order)
		},
		func(value float32, box []byte) (err error) {
			return CompileFloat32(value, box, order)
		},
	)
}

func Float32LE() Type[float32] {
	return Float32(binary.LittleEndian)
}

func Float32BE() Type[float32] {
	return Float32(binary.BigEndian)
}
