package parco

import (
	"io"
)

type (
	varType[T any] struct {
		header IntType

		pool Pooler

		sizer sizer[T]

		parser ParserFunc[T]

		compiler func(T, io.Writer) error
	}
)

func (v varType[T]) ByteLength() int {
	return v.header.ByteLength()
}

func (v varType[T]) Header() IntType {
	return v.header
}

func (v varType[T]) ParseBytes(box []byte) (res T, err error) {
	return v.parser(box)
}

func (v varType[T]) Parse(r io.Reader) (res T, err error) {
	var (
		read, size int
	)

	if size, err = v.header.Parse(r); err != nil {
		return
	}

	box := v.pool.Get(size)
	defer v.pool.Put(box)
	data := *box
	data = data[:size]

	if read, err = r.Read(data); err != nil {
		return
	}

	if read != size {
		// TODO: wrap
		err = ErrCannotRead
		return
	}

	return v.ParseBytes(data)
}

func (v varType[T]) Compile(value T, w io.Writer) (err error) {
	size := v.sizer.Len(value)
	if MaxSize(v.header.ByteLength()) < size {
		err = ErrOverflow
		return
	}

	if err = v.header.Compile(size, w); err != nil {
		return
	}

	err = v.compiler(value, w)
	return
}
