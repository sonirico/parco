package internal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type iterable interface {
	Range(ranger) error
	Len() int
}

var (
	NoopArr = ArrayType{}
)

type ArrayType struct {
	length int
	header Type
	inner  Type
}

func (t ArrayType) Length() int {
	return t.length * t.inner.Length()
}

func (t ArrayType) Parse(r io.Reader) (interface{}, error) {
	data := make([]byte, t.length*t.inner.Length())
	_, err := r.Read(data)
	if err != nil {
		return nil, err
	}
	return ArrayValue{data: data, innerType: t.inner}, nil
}

func (t ArrayType) Compile(x interface{}, w io.Writer) error {
	if err := t.header.Compile(t.length, w); err != nil {
		return err
	}

	iter := x.(iterable)

	if iter.Len() != t.length {
		return NewErrCompile(fmt.Sprintf("unexpected length, want %d but have %d",
			t.length, iter.Len()))
	}

	return iter.Range(func(x interface{}) error {
		if err := t.inner.Compile(x, w); err != nil {
			return err
		}
		return nil
	})
}

func Array(length int, header, inner Type) ArrayType {
	return ArrayType{length: length, header: header, inner: inner}
}

type ArrayValue struct {
	data []byte

	innerType Type
}

var (
	NoopArrVal = ArrayValue{}
	noVal      = Value{}
)

func (v ArrayValue) At(pos int) (Value, error) {
	offset := pos * v.innerType.Length()
	if offset >= len(v.data) {
		return noVal, errors.New("out of bounds error")
	}

	limit := offset + v.innerType.Length()
	rawVal, err := v.innerType.Parse(bytes.NewBuffer(v.data[offset:limit])) // TODO: ParseBytes interface, bytes.Buffer not desired
	if err != nil {
		return noVal, err
	}

	return Value{value: rawVal}, nil
}

func (v ArrayValue) Range(f func(Value)) {
	limit := len(v.data) / v.innerType.Length()

	for i := 0; i < limit; i++ {
		v, err := v.At(i)
		if err == nil {
			f(v)
		}
	}
}

func (v ArrayValue) Len() int {
	return len(v.data)
}

type Value struct {
	value interface{}
}

func (v Value) GetInt8() int8 {
	return v.value.(int8)
}

func (v Value) GetUInt8() uint8 {
	return v.value.(uint8)
}
