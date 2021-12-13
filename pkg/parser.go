package pkg

import (
	"io"

	"github.com/sonirico/parco/internal"
	"github.com/sonirico/parco/internal/utils"
)

type Parser struct {
	fields []field
}

func (p Parser) ParseBytes(data []byte) (Result, error) {
	buf := utils.NewBufferCursor(data, 0)

	return p.parse(&buf)
}

func (p Parser) Parse(r io.Reader) (Result, error) {
	return p.parse(r)
}

func (p Parser) parse(r io.Reader) (Result, error) {
	s := newResult()

	for _, f := range p.fields {
		value, err := f.Type.Parse(r)

		if err != nil {
			return s, nil
		}

		if _, ok := f.Type.(internal.SkipType); !ok {
			s.data[f.Name] = structItem{
				field: f,
				value: value,
			}
		}
	}
	return s, nil
}
