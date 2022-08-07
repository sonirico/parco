package parco

import (
	"encoding/binary"
	"io"
)

func Blob(header IntType) Type[[]byte] {
	return varType[[]byte]{
		header:   header,
		sizer:    SizerFunc[[]byte](func(x []byte) int { return len(x) }),
		pool:     SinglePool,
		parser:   ParseBlob,
		compiler: CompileBlob,
	}
}

func NewVarcharType(header IntType) Type[string] {
	return varType[string]{
		header:   header,
		sizer:    SizerFunc[string](func(x string) int { return len(x) }),
		pool:     SinglePool,
		parser:   ParseStringFactory(),
		compiler: CompileStringWriter,
	}
}

func SmallVarchar() Type[string] {
	return NewVarcharType(UInt8Header())
}

func Varchar() Type[string] {
	return NewVarcharType(UInt16HeaderLE())
}

func VarcharOrder(order binary.ByteOrder) Type[string] {
	return NewVarcharType(UInt16Header(order))
}

func ParseStringFactory() ParserFunc[string] {
	return func(data []byte) (string, error) {
		return ParseString(data)
	}
}

func ParseString(data []byte) (res string, err error) {
	return string(data), nil
}

func CompileStringWriter(x string, w io.Writer) (err error) {
	var written int
	data := String2Bytes(x)
	written, err = w.Write(data)
	if written != len(x) {
		err = ErrCannotWrite
	}
	return
}

func CompileString(x string, box []byte) (err error) {
	bites := String2Bytes(x)

	if copy(box, bites) != len(bites) {
		return ErrCannotWrite
	}

	return
}

func CompileStringFactory() CompilerFunc[string] {
	return func(s string, box []byte) error {
		return CompileString(s, box)
	}
}

func ParseBlob(data []byte) ([]byte, error) {
	return data, nil
}

func CompileBlob(x []byte, w io.Writer) (err error) {
	var written int
	written, err = w.Write(x)
	if written != len(x) {
		err = ErrCannotWrite
	}
	return
}
