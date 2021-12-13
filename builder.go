package parco

import "encoding/binary"

type (
	field[T any] struct {
		Type Type[T]
		Name string
	}

	structItem[T any] struct {
		data  []byte
		value any
		field field[T]
	}

	getter[T, U any] func(T) U
	setter[T, U any] func(*T, U)

	Builder[T any] struct {
		fields []field[T]
	}
)

func (b Builder[T]) SmallVarchar(name string, getter getter[T, string]) Builder[T] {
	return b.register(name, SmallVarchar[T](getter))
}

func (b Builder[T]) UInt8(name string, getter getter[T, uint8]) Builder[T] {
	return b.register(name, UInt8C[T](getter))
}

func (b Builder[T]) UInt16(name string, order binary.ByteOrder, getter getter[T, uint16]) Builder[T] {
	return b.register(name, UInt16[T](order, getter))
}

func (b Builder[T]) UInt16BE(name string, getter getter[T, uint16]) Builder[T] {
	return b.register(name, UInt16[T](binary.BigEndian, getter))
}

func (b Builder[T]) UInt16LE(name string, getter getter[T, uint16]) Builder[T] {
	return b.register(name, UInt16[T](binary.LittleEndian, getter))
}

func (b Builder[T]) Field(name string, tp Type[T]) Builder[T] {
	return b.register(name, tp)
}

func (b Builder[T]) Skip(pad int) Builder[T] {
	return b.register("", SkipType[T]{pad: pad})
}

func (b Builder[T]) register(name string, tp Type[T]) Builder[T] {
	b.fields = append(b.fields, field[T]{Name: name, Type: tp})
	return b
}

func (b Builder[T]) Compiler() Compiler[T] {
	return Compiler[T]{fields: b.fields}
}

func NewBuilder[T any]() Builder[T] {
	// parser/compiler factories
	return Builder[T]{}
}
