package parco

import (
	"encoding/binary"
	"math"
)

func CompileUInt64(u64 uint64, box []byte, order binary.ByteOrder) (err error) {
	order.PutUint64(box, u64)
	return
}

func ParseUInt64(data []byte, order binary.ByteOrder) (uint64, error) {
	if len(data) < 8 {
		return 0, NewErrUnSufficientBytesError(4, 0)
	}
	return order.Uint64(data), nil
}

func UInt64(order binary.ByteOrder) Type[uint64] {
	return NewFixedType[uint64](
		8,
		func(data []byte) (uint64, error) {
			return ParseUInt64(data, order)
		},
		func(value uint64, box []byte) (err error) {
			return CompileUInt64(value, box, order)
		},
	)
}

func UInt64LE() Type[uint64] {
	return UInt64(binary.LittleEndian)
}

func UInt64BE() Type[uint64] {
	return UInt64(binary.BigEndian)
}

func UInt64Header(order binary.ByteOrder) Type[int] {
	return NewFixedType[int](
		8,
		func(data []byte) (int, error) {
			n, err := ParseUInt64(data, order)
			m := int(n)
			return m, err
		},
		func(value int, box []byte) (err error) {
			if value < 0 {
				return ErrOverflow
			}

			return CompileUInt64(uint64(value), box, order)
		},
	)
}

func UInt64LEHeader() Type[int] {
	return UInt64Header(binary.LittleEndian)
}

func UInt64BEHeader() Type[int] {
	return UInt64Header(binary.BigEndian)
}

func CompileInt64(i64 int64, box []byte, order binary.ByteOrder) (err error) {
	order.PutUint64(box, uint64(i64))
	return
}

func ParseInt64(data []byte, order binary.ByteOrder) (int64, error) {
	if len(data) < 8 {
		return 0, NewErrUnSufficientBytesError(4, 0)
	}

	return int64(order.Uint64(data)), nil
}

func Int64(order binary.ByteOrder) Type[int64] {
	return NewFixedType[int64](
		8,
		func(data []byte) (int64, error) {
			return ParseInt64(data, order)
		},
		func(value int64, box []byte) (err error) {
			return CompileInt64(value, box, order)
		},
	)
}

func Int64LE() Type[int64] {
	return Int64(binary.LittleEndian)
}

func Int64BE() Type[int64] {
	return Int64(binary.BigEndian)
}

func Int64Header(order binary.ByteOrder) Type[int] {
	return NewFixedType[int](
		8,
		func(data []byte) (int, error) {
			n, err := ParseInt64(data, order)
			m := int(n)
			return m, err
		},
		func(value int, box []byte) (err error) {
			// On 64-bit systems int is int64, so no overflow check needed
			//nolint:staticcheck // SA4003: Condition is always false on 64-bit, but kept for clarity
			if value > math.MaxInt64 || value < math.MinInt64 {
				return ErrOverflow
			}
			return CompileInt64(int64(value), box, order)
		},
	)
}

func Int64LEHeader() Type[int] {
	return Int64Header(binary.LittleEndian)
}

func Int64BEHeader() Type[int] {
	return Int64Header(binary.BigEndian)
}
