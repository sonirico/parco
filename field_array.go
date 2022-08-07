package parco

import "io"

type (
	arrayField[T, U any] struct {
		id     string
		header IntType
		inner  ArrayType[U]
		setter Setter[T, Slice[U]]
		getter Getter[T, Slice[U]]
		pooler Pooler
	}
)

func (s arrayField[T, U]) ID() string {
	return s.id
}

func (s arrayField[T, U]) Parse(item *T, r io.Reader) error {
	values, err := s.inner.Parse(r)
	if err != nil {
		return err
	}
	s.setter(item, values.Unwrap())
	return nil
}

func (s arrayField[T, U]) Compile(item *T, w io.Writer) error {
	value := s.getter(item)
	return s.inner.Compile(value, w)
}

func ArrayField[T, U any](
	header IntType,
	inner Type[U],
	setter Setter[T, Slice[U]],
	getter Getter[T, Slice[U]],
) Field[T, U] {
	return arrayField[T, U]{
		header: header,
		inner:  NewArrayType[U](header, inner),
		setter: setter,
		getter: getter,
		pooler: SinglePool,
	}
}

func ArrayFieldGetter[T any, U any](
	header IntType,
	inner Type[U],
	getter Getter[T, Slice[U]],
) Field[T, U] {
	return ArrayField[T, U](
		header,
		inner,
		nil,
		getter,
	)
}

func ArrayFieldSetter[T, U any](
	header IntType,
	inner Type[U],
	setter Setter[T, Slice[U]],
) Field[T, U] {
	return ArrayField[T, U](
		header,
		inner,
		setter,
		nil,
	)
}
