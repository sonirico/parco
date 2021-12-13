package parco

import (
	"encoding/binary"
	"io"
)

type Compiler[T any] struct {
	fields []field[T]
}

func (c Compiler[T]) Compile(value T, w io.Writer) error {
	for _, f := range c.fields {
		if _, ok := f.Type.(SkipType[T]); ok {
			continue
		}

		err := f.Type.Compile(value, w)
		if err != nil {
			return err
		}

	}
	return nil
}

func NewCompiler[T any]() Compiler[T] {
	return Compiler[T]{}
}

func (c Compiler[T]) Array(name string, tp Type[T]) Compiler[T] {
	return c.register(name, tp)
}

func (c Compiler[T]) SmallVarchar(name string, getter getter[T, string]) Compiler[T] {
	return c.register(name, SmallVarchar[T](getter))
}

func (c Compiler[T]) Varchar(name string, getter getter[T, string]) Compiler[T] {
	return c.register(name, Varchar[T](getter))
}

func (c Compiler[T]) UInt8(name string, getter getter[T, uint8]) Compiler[T] {
	return c.register(name, UInt8C[T](getter))
}

func (c Compiler[T]) UInt16(name string, order binary.ByteOrder, getter getter[T, uint16]) Compiler[T] {
	return c.register(name, UInt16[T](order, getter))
}

func (c Compiler[T]) UInt16BE(name string, getter getter[T, uint16]) Compiler[T] {
	return c.register(name, UInt16[T](binary.BigEndian, getter))
}

func (c Compiler[T]) UInt16LE(name string, getter getter[T, uint16]) Compiler[T] {
	return c.register(name, UInt16[T](binary.LittleEndian, getter))
}

func (c Compiler[T]) Field(name string, tp Type[T]) Compiler[T] {
	return c.register(name, tp)
}

func (c Compiler[T]) Skip(pad int) Compiler[T] {
	return c.register("", SkipType[T]{pad: pad})
}

func (c Compiler[T]) register(name string, tp Type[T]) Compiler[T] {
	c.fields = append(c.fields, field[T]{Name: name, Type: tp})
	return c
}
