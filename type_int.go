package parco

import (
	"encoding/binary"
	"math"
)

const (
	is64BitsArch = (^uint(0) >> 63) == 1
)

var (
	uintBitSize = If[int](is64BitsArch, 8, 4)
)

func CompileUInt(u uint, box []byte, order binary.ByteOrder) (err error) {
	if is64BitsArch {
		return CompileUInt64(uint64(u), box, order)
	}
	return CompileUInt32(uint32(u), box, order)
}

func ParseUInt(data []byte, order binary.ByteOrder) (uint, error) {
	if is64BitsArch {
		u64, err := ParseUInt64(data, order)
		if err != nil {
			return 0, err
		}

		return uint(u64), err
	}

	u32, err := ParseUInt32(data, order)
	if err != nil {
		return 0, err
	}

	return uint(u32), err
}

func UInt(order binary.ByteOrder) Type[uint] {
	return NewFixedType[uint](
		uintBitSize,
		func(data []byte) (uint, error) {
			return ParseUInt(data, order)
		},
		func(value uint, box []byte) (err error) {
			return CompileUInt(value, box, order)
		},
	)
}

func UIntLE() Type[uint] {
	return UInt(binary.LittleEndian)
}

func UIntBE() Type[uint] {
	return UInt(binary.BigEndian)
}

func UIntHeader(order binary.ByteOrder) Type[int] {
	return NewFixedType[int](
		uintBitSize,
		func(data []byte) (int, error) {
			n, err := ParseUInt(data, order)
			m := int(n)
			return m, err
		},
		func(value int, box []byte) (err error) {
			if value < 0 {
				return ErrOverflow
			}

			return CompileUInt(uint(value), box, order)
		},
	)
}

func UIntLEHeader() Type[int] {
	return UIntHeader(binary.LittleEndian)
}

func UIntBEHeader() Type[int] {
	return UIntHeader(binary.BigEndian)
}

func CompileInt(i int, box []byte, order binary.ByteOrder) (err error) {
	if is64BitsArch {
		return CompileInt64(int64(i), box, order)
	}
	return CompileInt32(int32(i), box, order)
}

func ParseInt(data []byte, order binary.ByteOrder) (int, error) {
	if is64BitsArch {
		i64, err := ParseInt64(data, order)
		if err != nil {
			return 0, err
		}

		return int(i64), err
	}

	i32, err := ParseInt32(data, order)
	if err != nil {
		return 0, err
	}

	return int(i32), err
}

func Int(order binary.ByteOrder) Type[int] {
	return NewFixedType[int](
		uintBitSize,
		func(data []byte) (int, error) {
			return ParseInt(data, order)
		},
		func(value int, box []byte) (err error) {
			return CompileInt(value, box, order)
		},
	)
}

func IntLE() Type[int] {
	return Int(binary.LittleEndian)
}

func IntBE() Type[int] {
	return Int(binary.BigEndian)
}

func IntHeader(order binary.ByteOrder) Type[int] {
	return NewFixedType[int](
		uintBitSize,
		func(data []byte) (int, error) {
			n, err := ParseInt(data, order)
			m := int(n)
			return m, err
		},
		func(value int, box []byte) (err error) {
			if value > math.MaxInt || value < math.MinInt {
				return ErrOverflow
			}
			return CompileInt(value, box, order)
		},
	)
}

func IntLEHeader() Type[int] {
	return IntHeader(binary.LittleEndian)
}

func IntBEHeader() Type[int] {
	return IntHeader(binary.BigEndian)
}
