package parco

import (
	"encoding/binary"
	"io"
)

type (
	fieldParser[T any] interface {
		Parse(item *T, reader io.Reader) error
	}

	Parser[T any] interface {
		Parse(io.Reader) (T, error)
		ParseBytes([]byte) (T, error)
	}

	ModelParser[T any] struct {
		fields  []fieldParser[T]
		factory Factory[T]
	}
)

func ParserModel[T any](factory Factory[T]) *ModelParser[T] {
	return &ModelParser[T]{factory: factory}
}

func (p *ModelParser[T]) ParseBytes(data []byte) (T, error) {
	buf := NewBufferCursor(data, 0)

	return p.parse(&buf)
}

func (p *ModelParser[T]) Parse(r io.Reader) (T, error) {
	return p.parse(r)
}

func (p *ModelParser[T]) parse(r io.Reader) (T, error) {
	model := p.factory.Get()

	for _, f := range p.fields {
		if err := f.Parse(&model, r); err != nil {
			return model, nil
		}
	}

	return model, nil
}

func (p *ModelParser[T]) Struct(field fieldParser[T]) *ModelParser[T] {
	return p.register(field)
}

func (p *ModelParser[T]) Array(field fieldParser[T]) *ModelParser[T] {
	return p.register(field)
}

func (p *ModelParser[T]) Map(field fieldParser[T]) *ModelParser[T] {
	return p.register(field)
}

func (p *ModelParser[T]) SmallVarchar(setter Setter[T, string]) *ModelParser[T] {
	return p.register(StringFieldSetter[T](SmallVarchar(), setter))
}

func (p *ModelParser[T]) Varchar(setter Setter[T, string]) *ModelParser[T] {
	return p.register(StringFieldSetter[T](Varchar(), setter))
}

func (c *ModelParser[T]) Bool(setter Setter[T, bool]) *ModelParser[T] {
	return c.register(BoolFieldSetter[T](Bool(), setter))
}

func (p *ModelParser[T]) UInt8(setter Setter[T, uint8]) *ModelParser[T] {
	return p.register(UInt8FieldSetter[T](UInt8(), setter))
}

func (p *ModelParser[T]) Int8(setter Setter[T, int8]) *ModelParser[T] {
	return p.register(Int8FieldSetter[T](Int8(), setter))
}

func (p *ModelParser[T]) Byte(setter Setter[T, byte]) *ModelParser[T] {
	return p.register(UInt8FieldSetter[T](Byte(), setter))
}

func (p *ModelParser[T]) UInt16(order binary.ByteOrder, setter Setter[T, uint16]) *ModelParser[T] {
	return p.register(UInt16FieldSetter[T](UInt16(order), setter))
}

func (p *ModelParser[T]) UInt16LE(setter Setter[T, uint16]) *ModelParser[T] {
	return p.register(UInt16FieldSetter[T](UInt16LE(), setter))
}

func (p *ModelParser[T]) UInt16BE(setter Setter[T, uint16]) *ModelParser[T] {
	return p.register(UInt16FieldSetter[T](UInt16BE(), setter))
}

func (p *ModelParser[T]) UInt32(order binary.ByteOrder, setter Setter[T, uint32]) *ModelParser[T] {
	return p.register(UInt32FieldSetter[T](UInt32(order), setter))
}

func (p *ModelParser[T]) Int32(order binary.ByteOrder, setter Setter[T, int32]) *ModelParser[T] {
	return p.register(Int32FieldSetter[T](Int32(order), setter))
}

func (p *ModelParser[T]) UInt64(order binary.ByteOrder, setter Setter[T, uint64]) *ModelParser[T] {
	return p.register(UInt64FieldSetter[T](UInt64(order), setter))
}

func (p *ModelParser[T]) Int64(order binary.ByteOrder, setter Setter[T, int64]) *ModelParser[T] {
	return p.register(Int64FieldSetter[T](Int64(order), setter))
}

func (p *ModelParser[T]) Int(order binary.ByteOrder, setter Setter[T, int]) *ModelParser[T] {
	return p.register(IntFieldSetter[T](Int(order), setter))
}

func (p *ModelParser[T]) Field(f fieldParser[T]) *ModelParser[T] {
	return p.register(f)
}

func (p *ModelParser[T]) Skip(pad int) *ModelParser[T] {
	return p.register(DefaultSkipField[T](pad))
}

func (p *ModelParser[T]) register(f fieldParser[T]) *ModelParser[T] {
	p.fields = append(p.fields, f)
	return p
}
