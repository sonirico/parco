package parco

import (
	"encoding/binary"
	"io"
	"time"
)

type (
	fieldBuilder[T any] interface {
		fieldCompiler[T]
		fieldParser[T]
	}

	ModelBuilder[T any] struct {
		fields []fieldBuilder[T]

		parser *Parser[T]

		compiler *Compiler[T]
	}
)

func Builder[T any](factory Factory[T]) ModelBuilder[T] {
	return ModelBuilder[T]{
		parser:   ParserModel(factory),
		compiler: CompilerModel[T](),
	}
}

func (b ModelBuilder[T]) Compile(value T, w io.Writer) error {
	return b.compiler.Compile(value, w)
}

func (b ModelBuilder[T]) CompileAny(value any, w io.Writer) error {
	t, ok := value.(T)
	if !ok {
		return ErrUnknownType
	}
	return b.compiler.Compile(t, w)
}

func (b ModelBuilder[T]) Parse(r io.Reader) (T, error) {
	return b.parser.Parse(r)
}

func (b ModelBuilder[T]) ParseAny(r io.Reader) (any, error) {
	return b.parser.Parse(r)
}

func (b ModelBuilder[T]) Parco() (*Parser[T], *Compiler[T]) {
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

func (b ModelBuilder[T]) Slice(field fieldBuilder[T]) ModelBuilder[T] {
	b.parser.Slice(field)
	b.compiler.Slice(field)
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

func (b ModelBuilder[T]) Byte(getter Getter[T, byte], setter Setter[T, byte]) ModelBuilder[T] {
	b.parser.Byte(setter)
	b.compiler.Byte(getter)
	return b
}

func (b ModelBuilder[T]) Int8(getter Getter[T, int8], setter Setter[T, int8]) ModelBuilder[T] {
	b.parser.Int8(setter)
	b.compiler.Int8(getter)
	return b
}

func (b ModelBuilder[T]) UInt16(
	order binary.ByteOrder,
	getter Getter[T, uint16],
	setter Setter[T, uint16],
) ModelBuilder[T] {
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

func (b ModelBuilder[T]) Int32(
	order binary.ByteOrder,
	getter Getter[T, int32],
	setter Setter[T, int32],
) ModelBuilder[T] {
	b.parser.Int32(order, setter)
	b.compiler.Int32(order, getter)
	return b
}

func (b ModelBuilder[T]) UInt32(
	order binary.ByteOrder,
	getter Getter[T, uint32],
	setter Setter[T, uint32],
) ModelBuilder[T] {
	b.parser.UInt32(order, setter)
	b.compiler.UInt32(order, getter)
	return b
}

func (b ModelBuilder[T]) Int64(
	order binary.ByteOrder,
	getter Getter[T, int64],
	setter Setter[T, int64],
) ModelBuilder[T] {
	b.parser.Int64(order, setter)
	b.compiler.Int64(order, getter)
	return b
}

func (b ModelBuilder[T]) UInt64(
	order binary.ByteOrder,
	getter Getter[T, uint64],
	setter Setter[T, uint64],
) ModelBuilder[T] {
	b.parser.UInt64(order, setter)
	b.compiler.UInt64(order, getter)
	return b
}

func (b ModelBuilder[T]) Int(order binary.ByteOrder, getter Getter[T, int], setter Setter[T, int]) ModelBuilder[T] {
	b.parser.Int(order, setter)
	b.compiler.Int(order, getter)
	return b
}

func (b ModelBuilder[T]) Float32(
	order binary.ByteOrder,
	getter Getter[T, float32],
	setter Setter[T, float32],
) ModelBuilder[T] {
	b.parser.Float32(order, setter)
	b.compiler.Float32(order, getter)
	return b
}

func (b ModelBuilder[T]) Float64(
	order binary.ByteOrder,
	getter Getter[T, float64],
	setter Setter[T, float64],
) ModelBuilder[T] {
	b.parser.Float64(order, setter)
	b.compiler.Float64(order, getter)
	return b
}

func (b ModelBuilder[T]) Time(
	withLocation bool,
	getter Getter[T, time.Time],
	setter Setter[T, time.Time],
) ModelBuilder[T] {
	b.parser.Time(withLocation, setter)
	b.compiler.Time(withLocation, getter)
	return b
}

func (b ModelBuilder[T]) TimeUTC(getter Getter[T, time.Time], setter Setter[T, time.Time]) ModelBuilder[T] {
	return b.Time(false, getter, setter)
}

func (b ModelBuilder[T]) TimeLocation(getter Getter[T, time.Time], setter Setter[T, time.Time]) ModelBuilder[T] {
	return b.Time(true, getter, setter)
}

func (b ModelBuilder[T]) Option(field fieldBuilder[T]) ModelBuilder[T] {
	b.parser.Option(field)
	b.compiler.Option(field)
	return b
}

func (b ModelBuilder[T]) Field(f fieldBuilder[T]) ModelBuilder[T] {
	b.parser.Field(f)
	b.compiler.Field(f)
	return b
}
