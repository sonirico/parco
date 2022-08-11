package parco

import (
	"encoding/binary"
	"io"
)

type (
	fieldCompiler[T any] interface {
		Compile(item *T, writer io.Writer) error
	}

	Compiler[T any] interface {
		Compile(T, io.Writer) error
	}

	ModelCompiler[T any] struct {
		fields []fieldCompiler[T]
	}
)

func (c ModelCompiler[T]) Compile(value T, w io.Writer) error {
	for _, f := range c.fields {
		if err := f.Compile(&value, w); err != nil {
			return err
		}
	}

	return nil
}

func CompilerModel[T any]() *ModelCompiler[T] {
	return &ModelCompiler[T]{}
}

func (c *ModelCompiler[T]) Struct(field fieldCompiler[T]) *ModelCompiler[T] {
	return c.register(field)
}

func (c *ModelCompiler[T]) Array(field fieldCompiler[T]) *ModelCompiler[T] {
	return c.register(field)
}

func (c *ModelCompiler[T]) Map(field fieldCompiler[T]) *ModelCompiler[T] {
	return c.register(field)
}

func (c *ModelCompiler[T]) Varchar(getter Getter[T, string]) *ModelCompiler[T] {
	return c.register(StringFieldGetter[T](Varchar(), getter))
}

func (c *ModelCompiler[T]) SmallVarchar(getter Getter[T, string]) *ModelCompiler[T] {
	return c.register(StringFieldGetter[T](SmallVarchar(), getter))
}

func (c *ModelCompiler[T]) Bool(getter Getter[T, bool]) *ModelCompiler[T] {
	return c.register(BoolFieldGetter[T](Bool(), getter))
}

func (c *ModelCompiler[T]) UInt8(getter Getter[T, uint8]) *ModelCompiler[T] {
	return c.register(UInt8FieldGetter[T](UInt8(), getter))
}

func (c *ModelCompiler[T]) UInt16(order binary.ByteOrder, getter Getter[T, uint16]) *ModelCompiler[T] {
	return c.register(UInt16FieldGetter[T](UInt16(order), getter))
}

func (c *ModelCompiler[T]) UInt16LE(getter Getter[T, uint16]) *ModelCompiler[T] {
	return c.register(UInt16FieldGetter[T](UInt16LE(), getter))
}

func (c *ModelCompiler[T]) UInt16BE(getter Getter[T, uint16]) *ModelCompiler[T] {
	return c.register(UInt16FieldGetter[T](UInt16BE(), getter))
}

func (c *ModelCompiler[T]) Field(f fieldCompiler[T]) *ModelCompiler[T] {
	return c.register(f)
}

func (c *ModelCompiler[T]) register(field fieldCompiler[T]) *ModelCompiler[T] {
	c.fields = append(c.fields, field)
	return c
}
