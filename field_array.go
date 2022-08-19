package parco

import "io"

type (
	BasicArrayField[T, U any] struct {
		id     string
		size   int
		inner  ArrayType[U]
		setter Setter[T, SliceView[U]]
		getter Getter[T, SliceView[U]]
	}
)

func (s BasicArrayField[T, U]) ID() string {
	return s.id
}

func (s BasicArrayField[T, U]) Parse(item *T, r io.Reader) error {
	values, err := s.inner.Parse(r)
	if err != nil {
		return err
	}
	s.setter(item, values.Unwrap())
	return nil
}

func (s BasicArrayField[T, U]) Compile(item *T, w io.Writer) error {
	value := s.getter(item)
	return s.inner.Compile(value, w)
}

func ArrayField[T, U any](
	length int,
	inner Type[U],
	setter Setter[T, SliceView[U]],
	getter Getter[T, SliceView[U]],
) Field[T, U] {
	return BasicArrayField[T, U]{
		inner:  Array[U](length, inner),
		setter: setter,
		getter: getter,
	}
}

func ArrayFieldGetter[T any, U any](
	length int,
	inner Type[U],
	getter Getter[T, SliceView[U]],
) Field[T, U] {
	return ArrayField[T, U](
		length,
		inner,
		nil,
		getter,
	)
}

func ArrayFieldSetter[T, U any](
	length int,
	inner Type[U],
	setter Setter[T, SliceView[U]],
) Field[T, U] {
	return ArrayField[T, U](
		length,
		inner,
		setter,
		nil,
	)
}
