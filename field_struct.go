package parco

import "io"

type (
	structTypeI[T any] interface {
		ParserType[T]
		CompilerType[T]
	}

	structType[T any] struct {
		ParserType[T]
		CompilerType[T]
	}

	structField[T, U any] struct {
		id     string
		inner  structType[U]
		setter Setter[T, U]
		getter Getter[T, U]
	}
)

func (s structField[T, U]) ID() string {
	return s.id
}

func (s structField[T, U]) Parse(item *T, r io.Reader) error {
	model, err := s.inner.Parse(r)
	if err != nil {
		return err
	}
	s.setter(item, model)
	return nil
}

func (s structField[T, U]) Compile(item *T, w io.Writer) error {
	value := s.getter(item)
	return s.inner.Compile(value, w)
}

func newStructField[T, U any](
	setter Setter[T, U],
	getter Getter[T, U],
	compiler CompilerType[U],
	parser ParserType[U],
) Field[T, U] {
	return structField[T, U]{
		inner: structType[U]{
			ParserType:   parser,
			CompilerType: compiler,
		},
		setter: setter,
		getter: getter,
	}
}

func StructField[T, U any](
	getter Getter[T, U],
	setter Setter[T, U],
	inner structTypeI[U],
) Field[T, U] {
	return newStructField[T, U](
		setter,
		getter,
		inner,
		inner,
	)
}

func StructFieldGetter[T, U any](
	getter Getter[T, U],
	compiler CompilerType[U],
) Field[T, U] {
	return newStructField[T, U](
		nil,
		getter,
		compiler,
		nil,
	)
}

func StructFieldSetter[T, U any](
	setter Setter[T, U],
	parser ParserType[U],
) Field[T, U] {
	return newStructField[T, U](
		setter,
		nil,
		nil,
		parser,
	)
}
