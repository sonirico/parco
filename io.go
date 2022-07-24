package parco

import (
	"io"
)

type (
	Writer interface {
		io.Writer
		WriteByte(byte) error
	}
)
