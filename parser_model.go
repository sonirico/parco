package parco

import (
	"io"
)

type (
	ModelField[T any] interface {
		Parse(item *T, reader io.Reader) error
	}

	ParserModel[T any] struct {
		fields  []ModelField[T]
		factory Factory[T]
	}
)

func NewModelParser[T any](factory Factory[T]) ParserModel[T] {
	return ParserModel[T]{factory: factory}
}

func (p ParserModel[T]) ParseBytes(data []byte) (T, error) {
	buf := NewBufferCursor(data, 0)

	return p.parse(&buf)
}

func (p ParserModel[T]) Parse(r io.Reader) (T, error) {
	return p.parse(r)
}

func (p ParserModel[T]) parse(r io.Reader) (T, error) {
	model := p.factory.Get()

	for _, f := range p.fields {
		if err := f.Parse(&model, r); err != nil {
			return model, nil
		}
	}

	return model, nil
}

func (p ParserModel[T]) SmallVarchar(setter Setter[T, string]) ParserModel[T] {
	return p.register(FieldStringSetter[T](SmallVarchar[T](nil), setter))
}

//func (p ParserModel) UInt8(name string) ParserModel {
//	return p.register(name, UInt8C[any](nil))
//}
//
//func (p ParserModel) UInt16(name string, order binary.ByteOrder) ParserModel {
//	return p.register(name, UInt16[any](order, nil))
//}
//
//func (p ParserModel) UInt16BE(name string) ParserModel {
//	return p.register(name, UInt16[any](binary.BigEndian, nil))
//}
//
//func (p ParserModel) UInt16LE(name string) ParserModel {
//	return p.register(name, UInt16[any](binary.LittleEndian, nil))
//}
//
//func (p ParserModel) Array(name string, inner Type[any]) ParserModel {
//	return p.register(name, inner)
//}
//
//func (p ParserModel) Field(name string, tp Type[any]) ParserModel {
//	return p.register(name, tp)
//}
//
//func (p ParserModel) Skip(pad int) ParserModel {
//	return p.register("", SkipType[any]{pad: pad})
//}

func (p ParserModel[T]) register(f ModelField[T]) ParserModel[T] {
	p.fields = append(p.fields, f)
	return p
}
