package parco

import "io"

type FixedField[T, U any] struct {
	Id     string
	Type   Type[U]
	Setter Setter[T, U]
	Getter Getter[T, U]
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

func BoolField[T any](
	tp Type[bool],
	getter Getter[T, bool],
	setter Setter[T, bool],
) Field[T, bool] {
	return FixedField[T, bool]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func BoolFieldGetter[T any](
	tp Type[bool],
	getter Getter[T, bool],
) Field[T, bool] {
	return BoolField[T](tp, getter, nil)
}

func BoolFieldSetter[T any](
	tp Type[bool],
	setter Setter[T, bool],
) Field[T, bool] {
	return BoolField[T](tp, nil, setter)
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

func Int8Field[T any](
	tp Type[int8],
	getter Getter[T, int8],
	setter Setter[T, int8],
) Field[T, int8] {
	return FixedField[T, int8]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func Int8FieldGetter[T any](
	tp Type[int8],
	getter Getter[T, int8],
) Field[T, int8] {
	return Int8Field[T](tp, getter, nil)
}

func Int8FieldSetter[T any](
	tp Type[int8],
	setter Setter[T, int8],
) Field[T, int8] {
	return Int8Field[T](tp, nil, setter)
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

func Int16Field[T any](
	tp Type[int16],
	getter Getter[T, int16],
	setter Setter[T, int16],
) Field[T, int16] {
	return FixedField[T, int16]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func Int16FieldGetter[T any](
	tp Type[int16],
	getter Getter[T, int16],
) Field[T, int16] {
	return Int16Field[T](tp, getter, nil)
}

func Int16FieldSetter[T any](
	tp Type[int16],
	setter Setter[T, int16],
) Field[T, int16] {
	return Int16Field[T](tp, nil, setter)
}

func UInt32Field[T any](
	tp Type[uint32],
	getter Getter[T, uint32],
	setter Setter[T, uint32],
) Field[T, uint32] {
	return FixedField[T, uint32]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func UInt32FieldGetter[T any](
	tp Type[uint32],
	getter Getter[T, uint32],
) Field[T, uint32] {
	return UInt32Field[T](tp, getter, nil)
}

func UInt32FieldSetter[T any](
	tp Type[uint32],
	setter Setter[T, uint32],
) Field[T, uint32] {
	return UInt32Field[T](tp, nil, setter)
}

func Int32Field[T any](
	tp Type[int32],
	getter Getter[T, int32],
	setter Setter[T, int32],
) Field[T, int32] {
	return FixedField[T, int32]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func Int32FieldGetter[T any](
	tp Type[int32],
	getter Getter[T, int32],
) Field[T, int32] {
	return Int32Field[T](tp, getter, nil)
}

func Int32FieldSetter[T any](
	tp Type[int32],
	setter Setter[T, int32],
) Field[T, int32] {
	return Int32Field[T](tp, nil, setter)
}

func UInt64Field[T any](
	tp Type[uint64],
	getter Getter[T, uint64],
	setter Setter[T, uint64],
) Field[T, uint64] {
	return FixedField[T, uint64]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func UInt64FieldGetter[T any](
	tp Type[uint64],
	getter Getter[T, uint64],
) Field[T, uint64] {
	return UInt64Field[T](tp, getter, nil)
}

func UInt64FieldSetter[T any](
	tp Type[uint64],
	setter Setter[T, uint64],
) Field[T, uint64] {
	return UInt64Field[T](tp, nil, setter)
}

func Int64Field[T any](
	tp Type[int64],
	getter Getter[T, int64],
	setter Setter[T, int64],
) Field[T, int64] {
	return FixedField[T, int64]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func Int64FieldGetter[T any](
	tp Type[int64],
	getter Getter[T, int64],
) Field[T, int64] {
	return Int64Field[T](tp, getter, nil)
}

func Int64FieldSetter[T any](
	tp Type[int64],
	setter Setter[T, int64],
) Field[T, int64] {
	return Int64Field[T](tp, nil, setter)
}

func IntField[T any](
	tp Type[int],
	getter Getter[T, int],
	setter Setter[T, int],
) Field[T, int] {
	return FixedField[T, int]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func IntFieldGetter[T any](
	tp Type[int],
	getter Getter[T, int],
) Field[T, int] {
	return IntField[T](tp, getter, nil)
}

func IntFieldSetter[T any](
	tp Type[int],
	setter Setter[T, int],
) Field[T, int] {
	return IntField[T](tp, nil, setter)
}

func Float32Field[T any](
	tp Type[float32],
	getter Getter[T, float32],
	setter Setter[T, float32],
) Field[T, float32] {
	return FixedField[T, float32]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func Float32FieldGetter[T any](
	tp Type[float32],
	getter Getter[T, float32],
) Field[T, float32] {
	return Float32Field[T](tp, getter, nil)
}

func Float32FieldSetter[T any](
	tp Type[float32],
	setter Setter[T, float32],
) Field[T, float32] {
	return Float32Field[T](tp, nil, setter)
}

func Float64Field[T any](
	tp Type[float64],
	getter Getter[T, float64],
	setter Setter[T, float64],
) Field[T, float64] {
	return FixedField[T, float64]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
	}
}

func Float64FieldGetter[T any](
	tp Type[float64],
	getter Getter[T, float64],
) Field[T, float64] {
	return Float64Field[T](tp, getter, nil)
}

func Float64FieldSetter[T any](
	tp Type[float64],
	setter Setter[T, float64],
) Field[T, float64] {
	return Float64Field[T](tp, nil, setter)
}

func SkipField[T any](
	tp Type[any],
) Field[T, any] {
	return FixedField[T, any]{
		Type: tp,
	}
}

func DefaultSkipField[T any](pad int) Field[T, any] {
	return SkipField[T](SkipTypeFactory(pad))
}
