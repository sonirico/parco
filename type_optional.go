package parco

import "io"

type (
	OptionalType[T any] struct {
		header Type[bool]
		inner  Type[T]
		pool   Pooler
	}
)

func (i OptionalType[T]) Parse(r io.Reader) (*T, error) {
	some, err := i.header.Parse(r)
	if err != nil {
		return nil, err
	}

	if !some {
		return nil, nil
	}

	item, err := i.inner.Parse(r)
	if err != nil {
		return nil, err
	}

	return Ptr(item), nil
}

func (i OptionalType[T]) Compile(item *T, w io.Writer) (err error) {
	some := item != nil
	if err = i.header.Compile(some, w); err != nil {
		return
	}
	if !some {
		return
	}
	return i.inner.Compile(*item, w)
}

func Option[T any](inner Type[T]) OptionalType[T] {
	return OptionalType[T]{
		header: Bool(),
		inner:  inner,
	}
}
