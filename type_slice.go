package parco

import (
	"io"
)

const (
	// MaxReasonableSliceLength is the maximum allowed length for slices
	// to prevent malicious or corrupted data from causing excessive memory allocation.
	MaxReasonableSliceLength = 10_000_000 // 10 million elements
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

	// Validate length to prevent excessive memory allocation
	if length < 0 || length > MaxReasonableSliceLength {
		return nil, ErrOverflow
	}

	arrType := Array[T](length, t.inner)

	return arrType.Parse(r)
}

func (t SliceType[T]) Compile(x Iterable[T], w io.Writer) error {
	length := x.Len()

	if err := t.header.Compile(length, w); err != nil {
		return err
	}

	arrType := Array[T](length, t.inner)

	return arrType.Compile(x, w)
}

func Slice[T any](header IntType, inner Type[T]) SliceType[T] {
	return SliceType[T]{
		header: header,
		inner:  inner,
	}
}
