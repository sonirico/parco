package internal

import (
	"encoding/binary"
	"github.com/sonirico/parco/internal/utils"
	"io"
)

func parseUInt16Reader(r io.Reader, order binary.ByteOrder) (uint16, error) {
	data := make([]byte, 2)
	_, err := r.Read(data)
	if err != nil {
		return 0, err
	}
	return parseUInt16(data, order)
}

func parseUInt16(data []byte, order binary.ByteOrder) (uint16, error) {
	if len(data) < 2 {
		return 0, NewErrUnSufficientBytesError(2, 0)
	}
	return order.Uint16(data), nil
}

type UInt16Type struct {
	order binary.ByteOrder
}

func (i UInt16Type) Head() Head {
	return nil
}

func (i UInt16Type) Fixed() bool {
	return true
}

func (i UInt16Type) Length() int {
	return 2
}

func (i UInt16Type) Parse(r io.Reader) (interface{}, error) {
	return parseUInt16Reader(r, i.order)
}

func (i UInt16Type) Compile(x interface{}, w io.Writer) (err error) {
	data, ok := utils.AnyIntToInt(x)
	if !ok {
		return NewErrTypeAssertionError("int", "whatnot")
	}
	bts := make([]byte, 2)
	u16 := uint16(data)

	if int(u16) != data {
		return NewErrTypeAssertionError("uint16", "int")
	}

	i.order.PutUint16(bts, u16)
	_, err = w.Write(bts)
	return
}

func (i UInt16Type) ParseLength(data []byte) (int, error) {
	u16, err := parseUInt16(data, i.order)
	if err != nil {
		return 0, err
	}
	return int(u16), nil
}

func UInt16(order binary.ByteOrder) UInt16Type {
	return UInt16Type{order: order}
}

func UInt16LE() UInt16Type {
	return UInt16(binary.LittleEndian)
}

func UInt16BE() UInt16Type {
	return UInt16(binary.BigEndian)
}
