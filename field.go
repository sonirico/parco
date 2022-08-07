package parco

import "io"

type (
	Setter[T, U any] func(*T, U)
	Getter[T, U any] func(*T) U

	Field[T, U any] interface {
		ID() string
		Parse(*T, io.Reader) error
		Compile(*T, io.Writer) error
	}
)
