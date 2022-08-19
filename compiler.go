package parco

import (
	"encoding/binary"
	"io"
)

type (
	fieldCompiler[T any] interface {
		Compile(item *T, writer io.Writer) error
	}

	Compiler[T any] struct {
		fields []fieldCompiler[T]
	}
)

func (c Compiler[T]) Compile(value T, w io.Writer) error {
	for _, f := range c.fields {
		if err := f.Compile(&value, w); err != nil {
			return err
		}
	}

	return nil
}

func CompilerModel[T any]() *Compiler[T] {
	return &Compiler[T]{}
}

func (c *Compiler[T]) Struct(field fieldCompiler[T]) *Compiler[T] {
	return c.register(field)
}

func (c *Compiler[T]) Slice(field fieldCompiler[T]) *Compiler[T] {
	return c.register(field)
}

func (c *Compiler[T]) Array(field fieldCompiler[T]) *Compiler[T] {
	return c.register(field)
}

func (c *Compiler[T]) Map(field fieldCompiler[T]) *Compiler[T] {
	return c.register(field)
}

func (c *Compiler[T]) Varchar(getter Getter[T, string]) *Compiler[T] {
	return c.register(StringFieldGetter[T](Varchar(), getter))
}

func (c *Compiler[T]) SmallVarchar(getter Getter[T, string]) *Compiler[T] {
	return c.register(StringFieldGetter[T](SmallVarchar(), getter))
}

func (c *Compiler[T]) Bool(getter Getter[T, bool]) *Compiler[T] {
	return c.register(BoolFieldGetter[T](Bool(), getter))
}

func (c *Compiler[T]) Byte(getter Getter[T, byte]) *Compiler[T] {
	return c.register(UInt8FieldGetter[T](Byte(), getter))
}

func (c *Compiler[T]) UInt8(getter Getter[T, uint8]) *Compiler[T] {
	return c.register(UInt8FieldGetter[T](UInt8(), getter))
}

func (c *Compiler[T]) Int8(getter Getter[T, int8]) *Compiler[T] {
	return c.register(Int8FieldGetter[T](Int8(), getter))
}

func (c *Compiler[T]) UInt16(order binary.ByteOrder, getter Getter[T, uint16]) *Compiler[T] {
	return c.register(UInt16FieldGetter[T](UInt16(order), getter))
}

func (c *Compiler[T]) Int16(order binary.ByteOrder, getter Getter[T, int16]) *Compiler[T] {
	return c.register(Int16FieldGetter[T](Int16(order), getter))
}

func (c *Compiler[T]) UInt16LE(getter Getter[T, uint16]) *Compiler[T] {
	return c.register(UInt16FieldGetter[T](UInt16LE(), getter))
}

func (c *Compiler[T]) UInt16BE(getter Getter[T, uint16]) *Compiler[T] {
	return c.register(UInt16FieldGetter[T](UInt16BE(), getter))
}

func (c *Compiler[T]) UInt32(order binary.ByteOrder, getter Getter[T, uint32]) *Compiler[T] {
	return c.register(UInt32FieldGetter[T](UInt32(order), getter))
}

func (c *Compiler[T]) Int32(order binary.ByteOrder, getter Getter[T, int32]) *Compiler[T] {
	return c.register(Int32FieldGetter[T](Int32(order), getter))
}

func (c *Compiler[T]) UInt64(order binary.ByteOrder, getter Getter[T, uint64]) *Compiler[T] {
	return c.register(UInt64FieldGetter[T](UInt64(order), getter))
}

func (c *Compiler[T]) Int64(order binary.ByteOrder, getter Getter[T, int64]) *Compiler[T] {
	return c.register(Int64FieldGetter[T](Int64(order), getter))
}

func (c *Compiler[T]) Int(order binary.ByteOrder, getter Getter[T, int]) *Compiler[T] {
	return c.register(IntFieldGetter[T](Int(order), getter))
}

func (c *Compiler[T]) Option(f fieldCompiler[T]) *Compiler[T] {
	return c.register(f)
}

func (c *Compiler[T]) Field(f fieldCompiler[T]) *Compiler[T] {
	return c.register(f)
}

func (c *Compiler[T]) register(field fieldCompiler[T]) *Compiler[T] {
	c.fields = append(c.fields, field)
	return c
}
