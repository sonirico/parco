package internal

import (
	"errors"
	"io"

	"github.com/sonirico/parco/internal/utils"
)

func parseUInt8(r io.Reader) (uint8, error) {
	data := make([]byte, 1)
	n, err := r.Read(data)
	if err != nil || n != 1 {
		return 0, errors.New("TODO")
	}
	if len(data) < 1 {
		return 0, NewErrUnSufficientBytesError(1, 0)
	}
	return data[0], nil
}

type UInt8Type struct{}

func (i UInt8Type) Length() int {
	return 1
}

func (i UInt8Type) Parse(r io.Reader) (interface{}, error) {
	return parseUInt8(r)
}

func (i UInt8Type) Compile(x interface{}, w io.Writer) (err error) {
	data, ok := utils.AnyIntToInt(x)
	if !ok {
		return NewErrTypeAssertionError("int", "whatnot")
	}
	u8 := uint8(data)

	if data != int(u8) {
		return NewErrTypeAssertionError("int8", "int")
	}
	_, err = w.Write([]byte{u8})
	return
}

func UInt8() UInt8Type {
	return UInt8Type{}
}
