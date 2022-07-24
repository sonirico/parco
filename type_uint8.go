package parco

import (
	"errors"
	"io"
)

func parseUInt8(r io.Reader, pooler Pooler) (uint8, error) {
	//data := make([]byte, 1)
	b := pooler.Get(1)
	defer pooler.Put(b)
	data := *b
	data = data[:1]

	n, err := r.Read(data)
	if err != nil || n != 1 {
		return 0, errors.New("TODO")
	}
	if len(data) < 1 {
		return 0, NewErrUnSufficientBytesError(1, 0)
	}
	return data[0], nil
}

type (
	UInt8Type[T any] struct {
		getter getter[T, uint8]
		pooler Pooler
	}
)

func (i UInt8Type[T]) Length() int {
	return 1
}

func (i UInt8Type[T]) Parse(r io.Reader) (val any, err error) {
	val, err = parseUInt8(r, i.pooler)
	return
}

func (i UInt8Type[T]) ParseInt(r io.Reader) (int, error) {
	res, err := parseUInt8(r, i.pooler)
	val := int(res)
	return val, err
}

func (i UInt8Type[T]) ParseUInt8(r io.Reader) (val uint8, err error) {
	val, err = parseUInt8(r, i.pooler)
	return
}

func (i UInt8Type[T]) Compile(item T, w io.Writer) error {
	data := i.getter(item)
	return i.CompileUInt8(data, w)
}

func (i UInt8Type[T]) CompileInt(in int, w io.Writer) error {
	if in > 255 {
		return ErrUnSufficientBytes{want: in, have: 255}
	}
	return i.CompileUInt8(uint8(in), w)
}

func (i UInt8Type[T]) CompileUInt8(u8 uint8, w io.Writer) (err error) {
	_, err = w.Write([]byte{u8})
	return
}

func UInt8C[T any](getter getter[T, uint8]) UInt8Type[T] {
	return UInt8Type[T]{getter: getter, pooler: SinglePool}
}

func UInt8Body() UInt8Type[uint8] {
	return UInt8Type[uint8]{getter: Identity[uint8], pooler: SinglePool}
}

func AnyUInt8Body() UInt8Type[any] {
	return UInt8Type[any]{
		getter: func(x any) uint8 {
			return x.(uint8)
		},
		pooler: SinglePool,
	}
}

func UInt8Header() UInt8Type[int] {
	return UInt8Type[int]{pooler: SinglePool}
}
