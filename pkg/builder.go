package pkg

import (
	"github.com/sonirico/parco/internal"
)

type field struct {
	Type   internal.Type
	Name   string
	Getter Getter
}

type structItem struct {
	field field
	data  []byte
	value interface{}
}

type Getter func(x interface{}) interface{}

type Builder struct {
	fields []field
}

func (b Builder) Field(name string, t internal.Type) Builder {
	if t.Length() < 1 {
		return b
	}
	b.fields = append(b.fields, field{Name: name, Type: t})
	return b
}

func (b Builder) FieldGet(name string, t internal.Type, g Getter) Builder {
	if t.Length() < 1 {
		return b
	}
	b.fields = append(b.fields, field{Name: name, Type: t, Getter: g})
	return b
}

func (b Builder) Skip(_ int) {
	b.fields = append(b.fields, field{})
}

func (b Builder) parser() Parser {
	return Parser{fields: b.fields}
}

func (b Builder) Parser() Parser {
	return Parser{fields: b.fields}
}

func (b Builder) Compiler() Compiler {
	return Compiler{fields: b.fields}
}

func (b Builder) ParCo() (Parser, Compiler) {
	return b.Parser(), b.Compiler()
}

func NewBuilder() Builder {
	// parser/compiler factories
	return Builder{}
}
