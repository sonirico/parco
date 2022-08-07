package parco

import "io"

type (
	fixedType[T any] struct {
		byteLength int
		parser     ParserFunc[T]
		compiler   CompilerFunc[T]
		pool       Pooler
	}
)

func (i fixedType[T]) ByteLength() int {
	return i.byteLength
}

func (i fixedType[T]) ParseBytes(data []byte) (res T, err error) {
	return i.parser(data)
}

func (i fixedType[T]) Parse(r io.Reader) (res T, err error) {
	box := i.pool.Get(i.byteLength)
	defer i.pool.Put(box)
	data := *box
	data = data[:i.byteLength]

	read, err := r.Read(data)

	if err != nil {
		return
	}

	if read != i.byteLength {
		// TODO: wrap
		err = ErrCannotRead
		return
	}

	return i.parser(data)
}

func (i fixedType[T]) Compile(item T, w io.Writer) (err error) {
	box := i.pool.Get(i.byteLength)
	defer i.pool.Put(box)
	data := *box
	data = data[:i.byteLength]

	if err = i.CompileBytes(item, data); err != nil {
		return err
	}

	written, err := w.Write(data)
	if err != nil {
		return
	}

	if written != i.byteLength {
		err = ErrCannotWrite
		return
	}

	return
}

func (i fixedType[T]) CompileBytes(data T, box []byte) error {
	return i.compiler(data, box)
}

func NewFixedType[T any](
	byteLength int,
	parserFunc ParserFunc[T],
	compilerFunc CompilerFunc[T],
) Type[T] {
	return fixedType[T]{
		byteLength: byteLength,
		parser:     parserFunc,
		compiler:   compilerFunc,
		pool:       SinglePool,
	}
}
