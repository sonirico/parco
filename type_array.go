package parco

import (
	"io"
)

const (
	// maxInitialCapacity caps the capacity pre-allocated from a wire-provided
	// length, so a lying header cannot force a large allocation before any
	// element is actually parsed. Collections grow past it as data arrives.
	maxInitialCapacity = 4096
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
	values := make([]T, 0, min(t.length, maxInitialCapacity))

	for range t.length {
		var value T
		if value, err = t.inner.Parse(r); err != nil {
			return
		}
		values = append(values, value)
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
