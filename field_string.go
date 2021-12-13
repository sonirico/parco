package parco

import "io"

type (
	Setter[T, U any] func(*T, U)
	Getter[T, U any] func(*T) U

	StringField[T any] struct {
		Type   VarcharType[T]
		Setter func(*T, string)
		Getter func(*T) string
	}
)

func (s StringField[T]) Parse(item *T, r io.Reader) error {
	value, err := s.Type.ParseString(r)
	if err != nil {
		return err
	}

	s.Setter(item, value)
	return nil
}

func FieldString[T any](
	tp VarcharType[T],
	getter Getter[T, string],
	setter Setter[T, string],
) StringField[T] {
	return StringField[T]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func FieldStringGetter[T any](
	tp VarcharType[T],
	getter Getter[T, string],
) StringField[T] {
	return StringField[T]{
		Type:   tp,
		Getter: getter,
	}
}

func FieldStringSetter[T any](
	tp VarcharType[T],
	setter Setter[T, string],
) StringField[T] {
	return StringField[T]{
		Type:   tp,
		Setter: setter,
	}
}
