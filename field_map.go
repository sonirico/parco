package parco

import "io"

type (
	mapField[T any, K comparable, V any] struct {
		id     string
		inner  mapType[K, V]
		setter Setter[T, map[K]V]
		getter Getter[T, map[K]V]
		pooler Pooler
	}
)

func (s mapField[T, K, V]) ID() string {
	return s.id
}

func (s mapField[T, K, V]) Parse(item *T, r io.Reader) error {
	values, err := s.inner.Parse(r)
	if err != nil {
		return err
	}
	s.setter(item, values)
	return nil
}

func (s mapField[T, K, V]) Compile(item *T, w io.Writer) error {
	value := s.getter(item)
	return s.inner.Compile(value, w)
}

func MapField[T any, K comparable, V any](
	header IntType,
	keyType Type[K],
	valueType Type[V],
	setter Setter[T, map[K]V],
	getter Getter[T, map[K]V],
) Field[T, map[K]V] {
	return mapField[T, K, V]{
		inner:  MapType[K, V](header, keyType, valueType),
		setter: setter,
		getter: getter,
		pooler: SinglePool,
	}
}

func MapFieldGetter[T any, K comparable, V any](
	header IntType,
	keyType Type[K],
	valueType Type[V],
	getter Getter[T, map[K]V],
) Field[T, map[K]V] {
	return MapField[T, K, V](
		header,
		keyType,
		valueType,
		nil,
		getter,
	)
}

func MapFieldSetter[T any, K comparable, V any](
	header IntType,
	keyType Type[K],
	valueType Type[V],
	setter Setter[T, map[K]V],
) Field[T, map[K]V] {
	return MapField[T, K, V](
		header,
		keyType,
		valueType,
		setter,
		nil,
	)
}
