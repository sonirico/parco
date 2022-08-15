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
		pool   Pooler
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
	t.length = length
	if err != nil {
		return nil, err
	}

	values := make([]T, t.length)

	// TODO: Consider using ParseBytes in order to allocate 1 []byte only
	for i := 0; i < t.length; i++ {
		values[i], err = t.inner.Parse(r)
		if err != nil {
			return
		}
	}

	return SliceView[T](values), nil
}

func (t SliceType[T]) Compile(x Iterable[T], w io.Writer) error {
	t.length = x.Len()

	if err := t.header.Compile(t.length, w); err != nil {
		return err
	}

	return x.Range(func(x T) error {
		if err := t.inner.Compile(x, w); err != nil {
			return err
		}
		return nil
	})
}

func Slice[T any](header IntType, inner Type[T]) SliceType[T] {
	return SliceType[T]{
		header: header,
		inner:  inner,
		pool:   SinglePool,
	}
}