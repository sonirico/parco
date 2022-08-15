package parco

import "io"

type (
	BasicSliceField[T, U any] struct {
		id     string
		header IntType
		inner  SliceType[U]
		setter Setter[T, SliceView[U]]
		getter Getter[T, SliceView[U]]
	}
)

func (s BasicSliceField[T, U]) ID() string {
	return s.id
}

func (s BasicSliceField[T, U]) Parse(item *T, r io.Reader) error {
	values, err := s.inner.Parse(r)
	if err != nil {
		return err
	}
	s.setter(item, values.Unwrap())
	return nil
}

func (s BasicSliceField[T, U]) Compile(item *T, w io.Writer) error {
	value := s.getter(item)
	return s.inner.Compile(value, w)
}

func SliceField[T, U any](
	header IntType,
	inner Type[U],
	setter Setter[T, SliceView[U]],
	getter Getter[T, SliceView[U]],
) Field[T, U] {
	return BasicSliceField[T, U]{
		header: header,
		inner:  Slice[U](header, inner),
		setter: setter,
		getter: getter,
	}
}

func SliceFieldGetter[T any, U any](
	header IntType,
	inner Type[U],
	getter Getter[T, SliceView[U]],
) Field[T, U] {
	return SliceField[T, U](
		header,
		inner,
		nil,
		getter,
	)
}

func SliceFieldSetter[T, U any](
	header IntType,
	inner Type[U],
	setter Setter[T, SliceView[U]],
) Field[T, U] {
	return SliceField[T, U](
		header,
		inner,
		setter,
		nil,
	)
}
