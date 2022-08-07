package parco

import "io"

type FixedField[T, U any] struct {
	Id     string
	Type   Type[U]
	Setter Setter[T, U]
	Getter Getter[T, U]
	Pooler Pooler
}

func (s FixedField[T, U]) ID() string {
	return s.Id
}

func (s FixedField[T, U]) Parse(item *T, r io.Reader) error {
	value, err := s.Type.Parse(r)
	if err != nil {
		return err
	}

	s.Setter(item, value)
	return nil
}

func (s FixedField[T, U]) Compile(item *T, w io.Writer) (err error) {
	value := s.Getter(item)
	return s.Type.Compile(value, w)
}

func UInt8Field[T any](
	tp Type[uint8],
	getter Getter[T, uint8],
	setter Setter[T, uint8],
) Field[T, uint8] {
	return FixedField[T, uint8]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
		Pooler: SinglePool,
	}
}

func UInt8FieldGetter[T any](
	tp Type[uint8],
	getter Getter[T, uint8],
) Field[T, uint8] {
	return UInt8Field[T](tp, getter, nil)
}

func UInt8FieldSetter[T any](
	tp Type[uint8],
	setter Setter[T, uint8],
) Field[T, uint8] {
	return UInt8Field[T](tp, nil, setter)
}

func UInt16Field[T any](
	tp Type[uint16],
	getter Getter[T, uint16],
	setter Setter[T, uint16],
) Field[T, uint16] {
	return FixedField[T, uint16]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
		Pooler: SinglePool,
	}
}

func UInt16FieldGetter[T any](
	tp Type[uint16],
	getter Getter[T, uint16],
) Field[T, uint16] {
	return UInt16Field[T](tp, getter, nil)
}

func UInt16FieldSetter[T any](
	tp Type[uint16],
	setter Setter[T, uint16],
) Field[T, uint16] {
	return UInt16Field[T](tp, nil, setter)
}

func SkipField[T any](
	tp Type[any],
) Field[T, any] {
	return FixedField[T, any]{
		Type:   tp,
		Pooler: SinglePool,
	}
}

func DefaultSkipField[T any](pad int) Field[T, any] {
	return SkipField[T](SkipTypeFactory(pad))
}
