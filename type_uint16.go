package parco

import (
	"encoding/binary"
)

func ParseUInt16(data []byte, order binary.ByteOrder) (uint16, error) {
	if len(data) < 2 {
		return 0, NewErrUnSufficientBytesError(2, 0)
	}
	return order.Uint16(data), nil
}

func ParseUInt16Factory(order binary.ByteOrder) ParserFunc[uint16] {
	return func(data []byte) (uint16, error) {
		return ParseUInt16(data, order)
	}
}

func ParseUInt16HeaderFactory(order binary.ByteOrder) ParserFunc[int] {
	return func(raw []byte) (int, error) {
		data, err := ParseUInt16(raw, order)
		if err != nil {
			return 0, err
		}
		return int(data), nil
	}
}

func CompileUInt16Factory(order binary.ByteOrder) CompilerFunc[uint16] {
	return func(u16 uint16, box []byte) (err error) {
		return CompileUInt16(u16, box, order)
	}
}

func CompileUInt16HeaderFactory(order binary.ByteOrder) CompilerFunc[int] {
	return func(u16 int, box []byte) (err error) {
		return CompileUInt16(uint16(u16), box, order)
	}
}

func CompileUInt16(u16 uint16, box []byte, order binary.ByteOrder) (err error) {
	order.PutUint16(box, u16)

	// TODO: widen Writer interface
	//b.WriteByte(byte(u16))
	//b.WriteByte(byte(u16 >> 8))
	return
}

func UInt16(order binary.ByteOrder) Type[uint16] {
	return NewFixedType[uint16](
		(2),
		ParseUInt16Factory(order),
		CompileUInt16Factory(order),
	)
}

func UInt16LE() Type[uint16] {
	return UInt16(binary.LittleEndian)
}

func UInt16BE() Type[uint16] {
	return UInt16(binary.BigEndian)
}

func UInt16Header(order binary.ByteOrder) Type[int] {
	return NewFixedType[int](
		(2),
		ParseUInt16HeaderFactory(order),
		CompileUInt16HeaderFactory(order),
	)
}

func UInt16HeaderLE() Type[int] {
	return UInt16Header(binary.LittleEndian)
}

func UInt16HeaderBE() Type[int] {
	return UInt16Header(binary.BigEndian)
}

func UInt16Factory(order binary.ByteOrder) Type[uint16] {
	return UInt16(order)
}

func UInt16LEFactory() Type[uint16] {
	return UInt16(binary.LittleEndian)
}

func UInt16BEFactory() Type[uint16] {
	return UInt16(binary.BigEndian)
}
