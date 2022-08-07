package parco

import (
	"io"
)

type (
	Iterable[T any] interface {
		Len() int
		Range(ranger[T]) error
		Unwrap() Slice[T]
	}

	ArrayType[T any] struct {
		length int
		header IntType
		inner  Type[T]
		pool   Pooler
	}
)

func (t ArrayType[T]) ByteLength() int {
	return t.header.ByteLength() + t.length*t.inner.ByteLength()
}

func (t ArrayType[T]) Parse(r io.Reader) (res Iterable[T], err error) {
	var (
		length int
	)
	length, err = t.header.Parse(r)
	t.length = length
	if err != nil {
		return nil, err
	}

	values := make([]T, t.length)

	// TODO: Consider using ParseBytes in order to allocate 1 []byte only
	for i := 0; i < t.length; i++ {
		values[i], err = t.inner.Parse(r)
		if err != nil {
			return
		}
	}

	return Slice[T](values), nil

	// TODO: pseudocode for LazyArray impl

	//le := t.length * t.inner.ByteLength()
	//bts := t.pool.Get(le)
	//defer t.pool.Put(bts)
	//data := *bts
	//data = data[:le]
	//_, err = r.Read(data)
	//if err != nil {
	//	return nil, err
	//}
	//return ArrayValue[T]{data: data, innerType: t.inner}, nil
}

func (t ArrayType[T]) Compile(x Iterable[T], w io.Writer) error {
	t.length = x.Len()

	if err := t.header.Compile(t.length, w); err != nil {
		return err
	}

	return x.Range(func(x T) error {
		if err := t.inner.Compile(x, w); err != nil {
			return err
		}
		return nil
	})
}

func NewArrayType[T any](header IntType, inner Type[T]) ArrayType[T] {
	return ArrayType[T]{
		header: header,
		inner:  inner,
		pool:   SinglePool,
	}
}

//type ArrayValue[T any] struct {
//	data []byte
//
//	innerType Type[T]
//}
//
//var (
//	NoopArrVal = ArrayValue[any]{}
//	noVal      = Value{}
//)
//
//func (v ArrayValue[T]) At(pos int) (val Value, err error) {
//	offset := pos * v.innerType.Length()
//	if offset >= byteLength(v.data) {
//		err = errors.New("out of bounds error")
//		return
//	}
//
//	limit := offset + v.innerType.Length()
//	rawVal, err := v.innerType.Parse(bytes.NewBuffer(v.data[offset:limit])) // TODO: ParseBytes interface, bytes.Buffer not desired
//	if err != nil {
//		return
//	}
//
//	val.value = rawVal
//	return
//}
//
//func (v ArrayValue[T]) Range(f func(Value)) {
//	limit := byteLength(v.data) / v.innerType.Length()
//
//	for i := 0; i < limit; i++ {
//		val, err := v.At(i)
//		if err == nil {
//			f(val)
//		}
//	}
//}
//
//func (v ArrayValue[T]) Len() int {
//	return byteLength(v.data)
//}
//
//type Value struct {
//	value any
//}
//
//func (v Value) Unwrap() any {
//	return v.value
//}
//
//func (v Value) GetInt8() int8 {
//	return v.value.(int8)
//}
//
//func (v Value) GetUInt8() uint8 {
//	return v.value.(uint8)
//}
//
