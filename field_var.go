package parco

func StringField[T any](
	tp Type[string],
	getter Getter[T, string],
	setter Setter[T, string],
) Field[T, string] {
	return FixedField[T, string]{
		Type:   tp,
		Setter: (setter),
		Getter: (getter),
		Pooler: SinglePool,
	}
}

func StringFieldGetter[T any](
	tp Type[string],
	getter Getter[T, string],
) Field[T, string] {
	return StringField[T](tp, getter, nil)
}

func StringFieldSetter[T any](
	tp Type[string],
	setter Setter[T, string],
) Field[T, string] {
	return StringField[T](tp, nil, setter)
}

func stringGetter[T any](getter Getter[T, string]) Getter[T, String] {
	return func(item *T) String {
		return String(getter(item))
	}
}

func stringSetter[T any](setter Setter[T, string]) Setter[T, String] {
	return func(item *T, s String) {
		setter(item, s.Unwrap())
	}
}