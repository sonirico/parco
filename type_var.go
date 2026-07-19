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
	var size int

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

	if size <= cap(*box) {
		data := (*box)[:size]
		if err = readFull(r, data); err != nil {
			return
		}
		return v.ParseBytes(data)
	}

	// The declared size exceeds the pooled buffer: read in chunks so a
	// corrupted or malicious header cannot force a large upfront allocation.
	var data []byte
	if data, err = readChunked(r, size, cap(*box)); err != nil {
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
