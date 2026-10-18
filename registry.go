//go:build go1.27

package parco

import (
	"fmt"
	"io"
)

type (
	Entry[ID comparable] struct {
		ID    ID
		Value any
	}

	registryModel interface {
		parseAny(r io.Reader) (any, error)
		compileAny(value any, w io.Writer) error
	}

	Registry[ID comparable] struct {
		header Type[ID]
		models map[ID]registryModel
	}
)

func (m *Model[T]) parseAny(r io.Reader) (any, error) {
	return m.Parse(r)
}

func (m *Model[T]) compileAny(value any, w io.Writer) error {
	v, ok := value.(T)
	if !ok {
		return ErrUnknownType
	}
	return m.Compile(v, w)
}

func NewRegistry[ID comparable](header Type[ID]) *Registry[ID] {
	return &Registry[ID]{header: header, models: make(map[ID]registryModel)}
}

func (r *Registry[ID]) Register[U any](id ID, model *Model[U]) *Registry[ID] {
	if _, ok := r.models[id]; ok {
		panic(fmt.Sprintf("parco: model %v already registered", id))
	}
	r.models[id] = model
	return r
}

func (r *Registry[ID]) ByteLength() int {
	return r.header.ByteLength()
}

func (r *Registry[ID]) Parse(reader io.Reader) (Entry[ID], error) {
	id, err := r.header.Parse(reader)
	if err != nil {
		return Entry[ID]{}, err
	}
	model, ok := r.models[id]
	if !ok {
		return Entry[ID]{ID: id}, fmt.Errorf("%w: %v", ErrUnknownType, id)
	}
	value, err := model.parseAny(reader)
	if err != nil {
		return Entry[ID]{ID: id}, err
	}
	return Entry[ID]{ID: id, Value: value}, nil
}

func (r *Registry[ID]) Compile(entry Entry[ID], w io.Writer) error {
	model, ok := r.models[entry.ID]
	if !ok {
		return fmt.Errorf("%w: %v", ErrUnknownType, entry.ID)
	}
	if err := r.header.Compile(entry.ID, w); err != nil {
		return err
	}
	return model.compileAny(entry.Value, w)
}
