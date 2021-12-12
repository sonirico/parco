package internal

import (
	"io"
)

type Head interface {
	Length() int

	ParseLength([]byte) (int, error)
}

type Type interface {
	Length() int

	Parse(r io.Reader) (interface{}, error)

	Compile(x interface{}, w io.Writer) error
}

type SkipType struct {
	pad int
}

func (t SkipType) Length() int {
	return t.pad
}

func (t SkipType) Parse(_ io.Reader) (interface{}, error) {
	return nil, nil
}

func (t SkipType) Compile(_ interface{}, _ io.Writer) error {
	return nil
}
