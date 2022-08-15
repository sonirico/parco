package parco

type (
	ranger[T any] func(x T) error

	sizer[T any] interface {
		Len(T) int
	}

	SizerFunc[T any] func(T) int
)

type (
	SliceView[T any] []T
)

func (s SliceView[T]) Len() int {
	return len(s)
}

func (s SliceView[T]) Range(fn ranger[T]) error {
	for _, x := range s {
		if err := fn(x); err != nil {
			return err
		}
	}

	return nil
}

func (s SliceView[T]) Unwrap() SliceView[T] {
	return s
}

func (s SizerFunc[T]) Len(item T) int {
	return s(item)
}
