package parco

func If[T any](condition bool, a, b T) T {
	if condition {
		return a
	}

	return b
}

func Ptr[T any](t T) *T {
	return &t
}
