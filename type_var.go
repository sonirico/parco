package parco

import (
	"io"
)

const (
	// MaxReasonableVarSize is the maximum allowed size for variable-length types
	// to prevent malicious or corrupted data from causing excessive memory allocation.
	// Set to 100MB - adjust based on your use case.
	MaxReasonableVarSize = 100 * 1024 * 1024 // 100MB
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

	// Validate size to prevent excessive memory allocation
	if size < 0 || size > MaxReasonableVarSize {
		err = ErrOverflow
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
