package parco

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

type (
	serializable[T comparable] interface {
		ParcoID() T
	}

	parserAny interface {
		ParseAny(io.Reader) (any, error)
	}

	compilerAny interface {
		CompileAny(any, io.Writer) error
	}

	builderAny interface {
		parserAny
		compilerAny
	}

	ModelMultiBuilder[T comparable] struct {
		header Type[T]

		parsers map[T]parserAny

		compilers map[T]compilerAny
	}
)

func MultiBuilder[T comparable](header Type[T]) *ModelMultiBuilder[T] {
	return &ModelMultiBuilder[T]{
		header:    header,
		parsers:   make(map[T]parserAny),
		compilers: make(map[T]compilerAny),
	}
}

func (b *ModelMultiBuilder[T]) Register(id T, builder builderAny) (*ModelMultiBuilder[T], error) {
	if _, ok := b.parsers[id]; !ok {
		b.parsers[id] = builder
		b.compilers[id] = builder
		return b, nil
	}
	return b, errors.Wrapf(ErrAlreadyRegistered, "id: %v", id)
}

func (b *ModelMultiBuilder[T]) MustRegister(id T, builder builderAny) *ModelMultiBuilder[T] {
	if _, ok := b.parsers[id]; !ok {
		b.parsers[id] = builder
		b.compilers[id] = builder
		return b
	}
	panic(fmt.Sprintf("this id is registered already: %v", id))
}

func (b *ModelMultiBuilder[T]) Parse(r io.Reader) (id T, res any, err error) {
	id, err = b.header.Parse(r)
	if err != nil {
		return
	}

	p, ok := b.parsers[id]

	if !ok {
		err = errors.Wrapf(ErrUnknownType, "%v", id)
		return
	}

	res, err = p.ParseAny(r)
	return
}

func (b *ModelMultiBuilder[T]) Compile(item serializable[T], w io.Writer) (err error) {
	return b.compile(item.ParcoID(), item, w)
}

func (b *ModelMultiBuilder[T]) CompileAny(id T, item any, w io.Writer) (err error) {
	return b.compile(id, item, w)
}

func (b *ModelMultiBuilder[T]) compile(id T, item any, w io.Writer) (err error) {
	err = b.header.Compile(id, w)

	if err != nil {
		return
	}

	c, ok := b.compilers[id]

	if !ok {
		err = errors.Wrapf(ErrUnknownType, "%v", id)
		return
	}

	err = c.CompileAny(item, w)
	return
}
