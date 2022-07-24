package parco

import (
	"io"
)

type Head interface {
	Length() int

	ParseLength([]byte) (int, error)
}

type (
	Type[T any] interface {
		Length() int

		Parse(r io.Reader) (any, error)

		Compile(item T, w io.Writer) error
	}

	IntType interface {
		Length() int

		CompileInt(i int, w io.Writer) error
		ParseInt(r io.Reader) (int, error)
	}

	SkipType[T any] struct {
		pad int
	}
)

func (t SkipType[T]) Length() int {
	return t.pad
}

func (t SkipType[T]) Parse(_ io.Reader) (res any, err error) {
	return
}

func (t SkipType[T]) Compile(_ T, _ io.Writer) error {
	return nil
}
