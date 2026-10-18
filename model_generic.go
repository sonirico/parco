//go:build go1.27

package parco

import "io"

type (
	modelField[T any] interface {
		byteLength() int
		parse(item *T, r io.Reader) error
		compile(item *T, w io.Writer) error
	}

	typedField[T, U any] struct {
		typ    Type[U]
		getter Getter[T, U]
		setter Setter[T, U]
	}

	Model[T any] struct {
		factory Factory[T]
		fields  []modelField[T]
	}
)

func (f typedField[T, U]) byteLength() int {
	return f.typ.ByteLength()
}

func (f typedField[T, U]) parse(item *T, r io.Reader) error {
	value, err := f.typ.Parse(r)
	if err != nil {
		return err
	}
	f.setter(item, value)
	return nil
}

func (f typedField[T, U]) compile(item *T, w io.Writer) error {
	return f.typ.Compile(f.getter(item), w)
}

func NewModel[T any](factory Factory[T]) *Model[T] {
	return &Model[T]{factory: factory}
}

func (m *Model[T]) Field[U any](typ Type[U], getter Getter[T, U], setter Setter[T, U]) *Model[T] {
	m.fields = append(m.fields, typedField[T, U]{typ: typ, getter: getter, setter: setter})
	return m
}

func (m *Model[T]) Skip(pad int) *Model[T] {
	return m.Field(SkipType(pad), func(*T) any { return nil }, func(*T, any) {})
}

func (m *Model[T]) ByteLength() int {
	total := 0
	for _, f := range m.fields {
		total += f.byteLength()
	}
	return total
}

func (m *Model[T]) ParseBytes(data []byte) (T, error) {
	pr := getParseReaderBytes(data)
	defer putParseReader(pr)
	return m.parse(pr)
}

func (m *Model[T]) Parse(r io.Reader) (T, error) {
	if pr, ok := r.(*parseReader); ok {
		// Nested inside another Parse call: the outer one owns the reader.
		return m.parse(pr)
	}
	pr := getParseReader(r)
	defer putParseReader(pr)
	return m.parse(pr)
}

func (m *Model[T]) parse(r io.Reader) (T, error) {
	item := m.factory.Get()
	for _, f := range m.fields {
		if err := f.parse(&item, r); err != nil {
			return item, err
		}
	}
	return item, nil
}

func (m *Model[T]) Compile(value T, w io.Writer) error {
	if _, ok := w.(*compileWriter); ok {
		return m.compile(value, w)
	}

	cw := getCompileWriter(w)
	defer putCompileWriter(cw)

	if err := m.compile(value, cw); err != nil {
		return err
	}
	return cw.flush()
}

func (m *Model[T]) compile(value T, w io.Writer) error {
	for _, f := range m.fields {
		if err := f.compile(&value, w); err != nil {
			return err
		}
	}
	return nil
}
