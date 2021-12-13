package parco

import (
	"encoding/binary"
	"io"
)

type ParserResult struct {
	fields []field[any]
}

func NewParserResult() ParserResult {
	return ParserResult{}
}

func (p ParserResult) ParseBytes(data []byte) (Result, error) {
	buf := NewBufferCursor(data, 0)

	return p.parse(&buf)
}

func (p ParserResult) Parse(r io.Reader) (Result, error) {
	return p.parse(r)
}

func (p ParserResult) parse(r io.Reader) (res Result, e error) {
	s := newResult()

	for _, f := range p.fields {
		// TODO: Parse lazy
		value, err := f.Type.Parse(r)

		if err != nil {
			return s, nil
		}

		if _, ok := f.Type.(SkipType[any]); !ok {
			s.data[f.Name] = structItem[any]{
				field: f,
				value: value,
			}
		}
	}
	return s, nil
}

func (c ParserResult) SmallVarchar(name string) ParserResult {
	return c.register(name, SmallVarchar[any](nil))
}

func (c ParserResult) UInt8(name string) ParserResult {
	return c.register(name, UInt8C[any](nil))
}

func (c ParserResult) UInt16(name string, order binary.ByteOrder) ParserResult {
	return c.register(name, UInt16[any](order, nil))
}

func (c ParserResult) UInt16BE(name string) ParserResult {
	return c.register(name, UInt16[any](binary.BigEndian, nil))
}

func (c ParserResult) UInt16LE(name string) ParserResult {
	return c.register(name, UInt16[any](binary.LittleEndian, nil))
}

func (c ParserResult) Array(name string, inner Type[any]) ParserResult {
	return c.register(name, inner)
}

func (c ParserResult) Field(name string, tp Type[any]) ParserResult {
	return c.register(name, tp)
}

func (c ParserResult) Skip(pad int) ParserResult {
	return c.register("", SkipType[any]{pad: pad})
}

func (c ParserResult) register(name string, tp Type[any]) ParserResult {
	c.fields = append(c.fields, field[any]{Name: name, Type: tp})
	return c
}
