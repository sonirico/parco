package pkg

import (
	"io"

	"github.com/sonirico/parco/internal"
)

type Compiler struct {
	fields []field
}

func (c Compiler) Compile(value interface{}, w io.Writer) error {
	for _, f := range c.fields {
		if _, ok := f.Type.(internal.SkipType); ok {
			continue
		}

		err := f.Type.Compile(f.Getter(value), w)
		if err != nil {
			return err
		}

	}
	return nil
}
