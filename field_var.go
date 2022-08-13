package parco

func StringField[T any](
	tp Type[string],
	getter Getter[T, string],
	setter Setter[T, string],
) Field[T, string] {
	return FixedField[T, string]{
		Type:   tp,
		Setter: setter,
		Getter: getter,
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
