package parco

import (
	"io"
)

type (
	Iterable[T any] interface {
		Len() int
		Range(ranger[T]) error
		Unwrap() SliceView[T]
	}

	SliceType[T any] struct {
		length int
		header IntType
		inner  Type[T]
	}
)

func (t SliceType[T]) ByteLength() int {
	return t.header.ByteLength() + t.length*t.inner.ByteLength()
}

func (t SliceType[T]) Parse(r io.Reader) (res Iterable[T], err error) {
	var (
		length int
	)
	length, err = t.header.Parse(r)
	if err != nil {
		return nil, err
	}

	arrType := Array[T](length, t.inner)

	t.length = length

	return arrType.Parse(r)
}

func (t SliceType[T]) Compile(x Iterable[T], w io.Writer) error {
	t.length = x.Len()

	if err := t.header.Compile(t.length, w); err != nil {
		return err
	}

	arrType := Array[T](t.length, t.inner)

	return arrType.Compile(x, w)
}

func Slice[T any](header IntType, inner Type[T]) SliceType[T] {
	return SliceType[T]{
		header: header,
		inner:  inner,
	}
}
