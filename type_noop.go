package parco

import "io"

type (
	ParserFunc[T any] func([]byte) (T, error)

	CompilerFunc[T any] func(data T, box []byte) error

	CompilerType[T any] interface {
		Compile(T, io.Writer) error
	}

	ParserType[T any] interface {
		Parse(reader io.Reader) (T, error)
	}

	Type[T any] interface {
		// ByteLength represents type byte length for this type. E.g: uint8=1, uint16=2, uint32=4
		// For non-fixed types, returns the byte length of the header
		ByteLength() int

		ParserType[T]
		CompilerType[T]
	}

	IntType = Type[int]
)
