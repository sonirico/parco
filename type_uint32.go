package parco

import (
	"encoding/binary"
	"math"
)

func CompileUInt32(u32 uint32, box []byte, order binary.ByteOrder) (err error) {
	order.PutUint32(box, u32)
	return
}

func ParseUInt32(data []byte, order binary.ByteOrder) (uint32, error) {
	if len(data) < 4 {
		return 0, NewErrUnSufficientBytesError(4, 0)
	}
	return order.Uint32(data), nil
}

func UInt32(order binary.ByteOrder) Type[uint32] {
	return NewFixedType[uint32](
		4,
		func(data []byte) (uint32, error) {
			return ParseUInt32(data, order)
		},
		func(value uint32, box []byte) (err error) {
			return CompileUInt32(value, box, order)
		},
	)
}

func UInt32LE() Type[uint32] {
	return UInt32(binary.LittleEndian)
}

func UInt32BE() Type[uint32] {
	return UInt32(binary.BigEndian)
}

func UInt32Header(order binary.ByteOrder) Type[int] {
	return NewFixedType[int](
		4,
		func(data []byte) (int, error) {
			n, err := ParseUInt32(data, order)
			m := int(n)
			return m, err
		},
		func(value int, box []byte) (err error) {
			if value > math.MaxUint32 || value < 0 {
				return ErrOverflow
			}

			return CompileUInt32(uint32(value), box, order)
		},
	)
}

func UInt32LEHeader() Type[int] {
	return UInt32Header(binary.LittleEndian)
}

func UInt32BEHeader() Type[int] {
	return UInt32Header(binary.BigEndian)
}

func CompileInt32(i32 int32, box []byte, order binary.ByteOrder) (err error) {
	order.PutUint32(box, uint32(i32))
	return
}

func ParseInt32(data []byte, order binary.ByteOrder) (int32, error) {
	if len(data) < 4 {
		return 0, NewErrUnSufficientBytesError(4, 0)
	}

	return int32(order.Uint32(data)), nil
}

func Int32(order binary.ByteOrder) Type[int32] {
	return NewFixedType[int32](
		4,
		func(data []byte) (int32, error) {
			return ParseInt32(data, order)
		},
		func(value int32, box []byte) (err error) {
			return CompileInt32(value, box, order)
		},
	)
}

func Int32LE() Type[int32] {
	return Int32(binary.LittleEndian)
}

func Int32BE() Type[int32] {
	return Int32(binary.BigEndian)
}

func Int32Header(order binary.ByteOrder) Type[int] {
	return NewFixedType[int](
		4,
		func(data []byte) (int, error) {
			n, err := ParseInt32(data, order)
			m := int(n)
			return m, err
		},
		func(value int, box []byte) (err error) {
			if value > math.MaxInt32 || value < math.MinInt32 {
				return ErrOverflow
			}
			return CompileInt32(int32(value), box, order)
		},
	)
}

func Int32LEHeader() Type[int] {
	return Int32Header(binary.LittleEndian)
}

func Int32BEHeader() Type[int] {
	return Int32Header(binary.BigEndian)
}
