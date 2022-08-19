package parco

import (
	"encoding/binary"
	"io"
)

type (
	fieldParser[T any] interface {
		Parse(item *T, reader io.Reader) error
	}

	Parser[T any] struct {
		fields  []fieldParser[T]
		factory Factory[T]
	}
)

func ParserModel[T any](factory Factory[T]) *Parser[T] {
	return &Parser[T]{factory: factory}
}

func (p *Parser[T]) ParseBytes(data []byte) (T, error) {
	buf := NewBufferCursor(data, 0)

	return p.parse(&buf)
}

func (p *Parser[T]) Parse(r io.Reader) (T, error) {
	return p.parse(r)
}

func (p *Parser[T]) parse(r io.Reader) (T, error) {
	model := p.factory.Get()

	for _, f := range p.fields {
		if err := f.Parse(&model, r); err != nil {
			return model, nil
		}
	}

	return model, nil
}

func (p *Parser[T]) Struct(field fieldParser[T]) *Parser[T] {
	return p.register(field)
}

func (p *Parser[T]) Slice(field fieldParser[T]) *Parser[T] {
	return p.register(field)
}

func (p *Parser[T]) Array(field fieldParser[T]) *Parser[T] {
	return p.register(field)
}

func (p *Parser[T]) Map(field fieldParser[T]) *Parser[T] {
	return p.register(field)
}

func (p *Parser[T]) SmallVarchar(setter Setter[T, string]) *Parser[T] {
	return p.register(StringFieldSetter[T](SmallVarchar(), setter))
}

func (p *Parser[T]) Varchar(setter Setter[T, string]) *Parser[T] {
	return p.register(StringFieldSetter[T](Varchar(), setter))
}

func (c *Parser[T]) Bool(setter Setter[T, bool]) *Parser[T] {
	return c.register(BoolFieldSetter[T](Bool(), setter))
}

func (p *Parser[T]) UInt8(setter Setter[T, uint8]) *Parser[T] {
	return p.register(UInt8FieldSetter[T](UInt8(), setter))
}

func (p *Parser[T]) Int8(setter Setter[T, int8]) *Parser[T] {
	return p.register(Int8FieldSetter[T](Int8(), setter))
}

func (p *Parser[T]) Byte(setter Setter[T, byte]) *Parser[T] {
	return p.register(UInt8FieldSetter[T](Byte(), setter))
}

func (p *Parser[T]) UInt16(order binary.ByteOrder, setter Setter[T, uint16]) *Parser[T] {
	return p.register(UInt16FieldSetter[T](UInt16(order), setter))
}

func (p *Parser[T]) UInt16LE(setter Setter[T, uint16]) *Parser[T] {
	return p.register(UInt16FieldSetter[T](UInt16LE(), setter))
}

func (p *Parser[T]) UInt16BE(setter Setter[T, uint16]) *Parser[T] {
	return p.register(UInt16FieldSetter[T](UInt16BE(), setter))
}

func (p *Parser[T]) UInt32(order binary.ByteOrder, setter Setter[T, uint32]) *Parser[T] {
	return p.register(UInt32FieldSetter[T](UInt32(order), setter))
}

func (p *Parser[T]) Int32(order binary.ByteOrder, setter Setter[T, int32]) *Parser[T] {
	return p.register(Int32FieldSetter[T](Int32(order), setter))
}

func (p *Parser[T]) UInt64(order binary.ByteOrder, setter Setter[T, uint64]) *Parser[T] {
	return p.register(UInt64FieldSetter[T](UInt64(order), setter))
}

func (p *Parser[T]) Int64(order binary.ByteOrder, setter Setter[T, int64]) *Parser[T] {
	return p.register(Int64FieldSetter[T](Int64(order), setter))
}

func (p *Parser[T]) Int(order binary.ByteOrder, setter Setter[T, int]) *Parser[T] {
	return p.register(IntFieldSetter[T](Int(order), setter))
}

func (p *Parser[T]) Float32(order binary.ByteOrder, setter Setter[T, float32]) *Parser[T] {
	return p.register(Float32FieldSetter[T](Float32(order), setter))
}

func (p *Parser[T]) Float64(order binary.ByteOrder, setter Setter[T, float64]) *Parser[T] {
	return p.register(Float64FieldSetter[T](Float64(order), setter))
}

func (p *Parser[T]) Option(f fieldParser[T]) *Parser[T] {
	return p.register(f)
}

func (p *Parser[T]) Field(f fieldParser[T]) *Parser[T] {
	return p.register(f)
}

func (p *Parser[T]) Skip(pad int) *Parser[T] {
	return p.register(DefaultSkipField[T](pad))
}

func (p *Parser[T]) register(f fieldParser[T]) *Parser[T] {
	p.fields = append(p.fields, f)
	return p
}
