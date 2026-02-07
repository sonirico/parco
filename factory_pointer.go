package parco

import "sync"

type (
	Factory[T any] interface {
		Get() T
	}

	PoolFactory[T any] interface {
		Factory[T]

		Put(T)
	}

	FuncFactory[T any] func() T

	nativePooledFactoryOption[T any] interface {
		Configure(f *NativePooledFactory[T])
	}

	nativePooledFactoryOptionFunc[T any] func(factory *NativePooledFactory[T])
)

func (f nativePooledFactoryOptionFunc[T]) Configure(factory *NativePooledFactory[T]) {
	f(factory)
}

func (f FuncFactory[T]) Get() T {
	return f()
}

func ObjectFactory[T any]() Factory[T] {
	return FuncFactory[T](func() (t T) {
		return
	})
}

type NativePooledFactory[T any] struct {
	inner sync.Pool

	resetFunc func(*T)
}

func (f *NativePooledFactory[T]) Get() T {
	//nolint:errcheck // Type assertion is safe - we control pool contents
	return f.inner.Get().(T)
}

func (f *NativePooledFactory[T]) Put(t T) {
	f.inner.Put(t)
}

func PooledFactory[T any](inner Factory[T], options ...nativePooledFactoryOption[T]) PoolFactory[T] {
	f := &NativePooledFactory[T]{
		inner: sync.Pool{New: func() any {
			return inner.Get()
		}},
	}

	for _, opt := range options {
		opt.Configure(f)
	}

	return f
}

func WithResetFunc[T any](fn func(*T)) nativePooledFactoryOption[T] {
	return nativePooledFactoryOptionFunc[T](func(factory *NativePooledFactory[T]) {
		factory.resetFunc = fn
	})
}
