package parco

import (
	"encoding/binary"
	"io"
)

type (
	fieldBuilder[T any] interface {
		fieldCompiler[T]
		fieldParser[T]
	}

	ModelBuilder[T any] struct {
		fields []fieldBuilder[T]

		parser *ModelParser[T]

		compiler *ModelCompiler[T]
	}
)

func (b ModelBuilder[T]) Compile(value T, w io.Writer) error {
	return b.compiler.Compile(value, w)
}

func (b ModelBuilder[T]) Parse(r io.Reader) (T, error) {
	return b.parser.Parse(r)
}

func Builder[T any](factory Factory[T]) ModelBuilder[T] {
	return ModelBuilder[T]{
		parser:   ParserModel(factory),
		compiler: CompilerModel[T](),
	}
}

func (b ModelBuilder[T]) ParCo() (Parser[T], Compiler[T]) {
	return b.parser, b.compiler
}

func (b ModelBuilder[T]) Struct(field fieldBuilder[T]) ModelBuilder[T] {
	b.parser.Struct(field)
	b.compiler.Struct(field)
	return b
}

func (b ModelBuilder[T]) Map(field fieldBuilder[T]) ModelBuilder[T] {
	b.parser.Map(field)
	b.compiler.Map(field)
	return b
}

func (b ModelBuilder[T]) Array(field fieldBuilder[T]) ModelBuilder[T] {
	b.parser.Array(field)
	b.compiler.Array(field)
	return b
}

func (b ModelBuilder[T]) Varchar(getter Getter[T, string], setter Setter[T, string]) ModelBuilder[T] {
	b.parser.Varchar(setter)
	b.compiler.Varchar(getter)
	return b
}

func (b ModelBuilder[T]) SmallVarchar(getter Getter[T, string], setter Setter[T, string]) ModelBuilder[T] {
	b.parser.SmallVarchar(setter)
	b.compiler.SmallVarchar(getter)
	return b
}

func (b ModelBuilder[T]) Bool(getter Getter[T, bool], setter Setter[T, bool]) ModelBuilder[T] {
	b.parser.Bool(setter)
	b.compiler.Bool(getter)
	return b
}

func (b ModelBuilder[T]) UInt8(getter Getter[T, uint8], setter Setter[T, uint8]) ModelBuilder[T] {
	b.parser.UInt8(setter)
	b.compiler.UInt8(getter)
	return b
}

func (b ModelBuilder[T]) UInt16(order binary.ByteOrder, getter Getter[T, uint16], setter Setter[T, uint16]) ModelBuilder[T] {
	b.parser.UInt16(order, setter)
	b.compiler.UInt16(order, getter)
	return b
}

func (b ModelBuilder[T]) UInt16LE(getter Getter[T, uint16], setter Setter[T, uint16]) ModelBuilder[T] {
	b.parser.UInt16LE(setter)
	b.compiler.UInt16LE(getter)
	return b
}

func (b ModelBuilder[T]) UInt16BE(getter Getter[T, uint16], setter Setter[T, uint16]) ModelBuilder[T] {
	b.parser.UInt16BE(setter)
	b.compiler.UInt16BE(getter)
	return b
}

func (b ModelBuilder[T]) Field(f fieldBuilder[T]) ModelBuilder[T] {
	b.parser.Field(f)
	b.compiler.Field(f)
	return b
}
