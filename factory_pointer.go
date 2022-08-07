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
)

type FuncFactory[T any] func() T

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
}

func (f NativePooledFactory[T]) Get() T {
	return f.inner.Get().(T)
}

func (f NativePooledFactory[T]) Put(t T) {
	f.inner.Put(t)
}

func PooledFactory[T any](inner Factory[T]) PoolFactory[T] {
	return NativePooledFactory[T]{
		inner: sync.Pool{New: func() any {
			return inner.Get()
		}},
	}
}
