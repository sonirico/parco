package parco

import (
	"bytes"
	"errors"
	"io"
)

type Iterable[T any] interface {
	Range(ranger[T]) error
	Len() int
}

var (
	NoopArr = ArrayType[any, any]{}
)

type (
	ArrayType[T, U any] struct {
		length int
		header IntType
		inner  Type[U]
		getter getter[T, Iterable[U]]
		pooler Pooler
	}
)

func (t ArrayType[T, U]) Length() int {
	return t.length * t.inner.Length()
}

func (t ArrayType[T, U]) Parse(r io.Reader) (any, error) {
	var err error
	t.length, err = t.header.ParseInt(r)
	if err != nil {
		return nil, err
	}

	le := t.length * t.inner.Length()
	bts := t.pooler.Get(le)
	defer t.pooler.Put(bts)
	data := *bts
	data = data[:le]
	_, err = r.Read(data)
	if err != nil {
		return nil, err
	}
	return ArrayValue[U]{data: data, innerType: t.inner}, nil
}

func (t ArrayType[T, U]) Compile(item T, w io.Writer) error {
	x := t.getter(item)
	return t.CompileIterable(x, w)
}

func (t ArrayType[T, U]) CompileIterable(x Iterable[U], w io.Writer) error {
	t.length = x.Len()

	if err := t.header.CompileInt(t.length, w); err != nil {
		return err
	}

	return x.Range(func(x U) error {
		if err := t.inner.Compile(x, w); err != nil {
			return err
		}
		return nil
	})
}

func Array[T, U any](header IntType, inner Type[U], getter getter[T, Iterable[U]]) ArrayType[T, U] {
	return ArrayType[T, U]{
		header: header,
		inner:  inner,
		getter: getter,
	}
}

func AnyArray(header IntType, inner Type[any], getter getter[any, Iterable[any]]) ArrayType[any, any] {
	return ArrayType[any, any]{
		header: header,
		inner:  inner,
		getter: getter,
	}
}

type ArrayValue[T any] struct {
	data []byte

	innerType Type[T]
}

var (
	NoopArrVal = ArrayValue[any]{}
	noVal      = Value{}
)

func (v ArrayValue[T]) At(pos int) (val Value, err error) {
	offset := pos * v.innerType.Length()
	if offset >= len(v.data) {
		err = errors.New("out of bounds error")
		return
	}

	limit := offset + v.innerType.Length()
	rawVal, err := v.innerType.Parse(bytes.NewBuffer(v.data[offset:limit])) // TODO: ParseBytes interface, bytes.Buffer not desired
	if err != nil {
		return
	}

	val.value = rawVal
	return
}

func (v ArrayValue[T]) Range(f func(Value)) {
	limit := len(v.data) / v.innerType.Length()

	for i := 0; i < limit; i++ {
		val, err := v.At(i)
		if err == nil {
			f(val)
		}
	}
}

func (v ArrayValue[T]) Len() int {
	return len(v.data)
}

type Value struct {
	value any
}

func (v Value) Unwrap() any {
	return v.value
}

func (v Value) GetInt8() int8 {
	return v.value.(int8)
}

func (v Value) GetUInt8() uint8 {
	return v.value.(uint8)
}
