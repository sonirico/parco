package parco

import (
	"io"
)

type VarcharType[T any] struct {
	head   IntType
	getter getter[T, string]
	pooler Pooler
}

func (v VarcharType[T]) Length() int {
	return v.head.Length()
}

func (v VarcharType[T]) Compile(item T, w io.Writer) (err error) {
	x := v.getter(item)
	return v.CompileString(x, w)
}

func (v VarcharType[T]) CompileString(x string, w io.Writer) (err error) {
	bites := String2Bytes(x)
	// TODO: Check whether header type and actual value do not overflow

	if err = v.head.CompileInt(len(bites), w); err != nil {
		return err
	}
	_, err = w.Write(bites)
	return err
}

func (v VarcharType[T]) Parse(r io.Reader) (res any, err error) {
	return v.ParseString(r)
}

func (v VarcharType[T]) ParseString(r io.Reader) (res string, err error) {
	length, err := v.head.ParseInt(r)
	if err != nil {
		return
	}

	b := v.pooler.Get(length)
	defer v.pooler.Put(b)
	data := *b
	data = data[:length]

	if _, err = r.Read(data); err != nil {
		return
	}

	res = Bytes2String(data)
	return
}

// Varchar handles strings up to 65535 bytes
func Varchar[T any](getter getter[T, string]) VarcharType[T] {
	return VarcharType[T]{
		getter: getter,
		pooler: SinglePool,
		head:   UInt16LEHeader[int](),
	}
}

// SmallVarchar handles strings up to 255 bytes
func SmallVarchar[T any](getter getter[T, string]) VarcharType[T] {
	return VarcharType[T]{
		getter: getter,
		pooler: SinglePool,
		head:   UInt8Header(),
	}
}
