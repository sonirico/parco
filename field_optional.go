package parco

import "io"

type OptionalField[T, U any] struct {
	id     string
	inner  OptionalType[U]
	setter Setter[T, *U]
	getter Getter[T, *U]
}

func (s OptionalField[T, U]) ID() string {
	return s.id
}

func (s OptionalField[T, U]) Parse(item *T, r io.Reader) error {
	value, err := s.inner.Parse(r)
	if err != nil {
		return err
	}

	s.setter(item, value)
	return nil
}

func (s OptionalField[T, U]) Compile(item *T, w io.Writer) (err error) {
	value := s.getter(item)
	return s.inner.Compile(value, w)
}

func OptionField[T, U any](
	tp Type[U],
	setter Setter[T, *U],
	getter Getter[T, *U],
) OptionalField[T, U] {
	return OptionalField[T, U]{
		inner:  Option[U](tp),
		setter: setter,
		getter: getter,
	}
}

func OptionFieldGetter[T, U any](
	tp Type[U],
	getter Getter[T, *U],
) OptionalField[T, U] {
	return OptionField[T, U](tp, nil, getter)
}

func OptionFieldSetter[T, U any](
	tp Type[U],
	setter Setter[T, *U],
) OptionalField[T, U] {
	return OptionField[T, U](tp, setter, nil)
}
