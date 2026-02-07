package parco

import (
	"io"
)

const (
	// MaxReasonableMapLength is the maximum allowed length for maps
	// to prevent malicious or corrupted data from causing excessive memory allocation.
	MaxReasonableMapLength = 10_000_000 // 10 million entries
)

type (
	mapType[K comparable, V any] struct {
		length    int
		header    IntType
		keyType   Type[K]
		valueType Type[V]
		pool      Pooler
	}
)

func (t mapType[K, V]) ByteLength() int {
	return t.header.ByteLength() + t.length*(t.keyType.ByteLength()+t.valueType.ByteLength())
}

func (t mapType[K, V]) Parse(r io.Reader) (res map[K]V, err error) {
	var (
		length int
	)
	length, err = t.header.Parse(r)
	t.length = length
	if err != nil {
		return nil, err
	}

	// Validate length to prevent excessive memory allocation
	if length < 0 || length > MaxReasonableMapLength {
		return nil, ErrOverflow
	}

	values := make(map[K]V, t.length)

	for i := 0; i < t.length; i++ {
		var (
			k K
			v V
		)
		if k, err = t.keyType.Parse(r); err != nil {
			return
		}
		if v, err = t.valueType.Parse(r); err != nil {
			return
		}

		values[k] = v
	}

	return values, nil
}

func (t mapType[K, V]) Compile(x map[K]V, w io.Writer) error {
	t.length = len(x)

	if err := t.header.Compile(t.length, w); err != nil {
		return err
	}

	for k, v := range x {
		if err := t.keyType.Compile(k, w); err != nil {
			return err
		}
		if err := t.valueType.Compile(v, w); err != nil {
			return err
		}
	}
	return nil
}

func MapType[K comparable, V any](header IntType, keyType Type[K], valueType Type[V]) mapType[K, V] {
	return mapType[K, V]{
		header:    header,
		keyType:   keyType,
		valueType: valueType,
		pool:      SinglePool,
	}
}
