package parco

import (
	"io"
)

type (
	ArrayType[T any] struct {
		length int
		inner  Type[T]
	}
)

func (t ArrayType[T]) ByteLength() int {
	return t.length * t.inner.ByteLength()
}

func (t ArrayType[T]) Parse(r io.Reader) (res Iterable[T], err error) {
	values := make([]T, t.length)

	for i := 0; i < t.length; i++ {
		values[i], err = t.inner.Parse(r)
		if err != nil {
			return
		}
	}

	return SliceView[T](values), nil
}

func (t ArrayType[T]) Compile(x Iterable[T], w io.Writer) error {
	if x.Len() != t.length {
		return ErrInvalidLength
	}
	return x.Range(func(x T) error {
		if err := t.inner.Compile(x, w); err != nil {
			return err
		}
		return nil
	})
}

func Array[T any](length int, inner Type[T]) ArrayType[T] {
	return ArrayType[T]{
		length: length,
		inner:  inner,
	}
}
