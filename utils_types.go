package parco

type (
	ranger[T any] func(x T) error

	sizer[T any] interface {
		Len(T) int
	}

	SizerFunc[T any] func(T) int
)

type (
	String string

	Slice[T any] []T
)

func (s String) Len() int {
	return len(s)
}

func (s String) Unwrap() string {
	return string(s)
}

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Range(fn ranger[T]) error {
	for _, x := range s {
		if err := fn(x); err != nil {
			return err
		}
	}

	return nil
}

func (s Slice[T]) Unwrap() Slice[T] {
	return s
}

func (s SizerFunc[T]) Len(item T) int {
	return s(item)
}