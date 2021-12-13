package parco

import (
	"encoding/binary"
	"io"
)

func parseUInt16Reader(r io.Reader, order binary.ByteOrder, pooler Pooler) (uint16, error) {
	b := pooler.Get(2)
	defer pooler.Put(b)
	data := *b
	data = data[:2]

	_, err := r.Read(data)
	if err != nil {
		return 0, err
	}
	return parseUInt16(data, order)
}

func parseUInt16(data []byte, order binary.ByteOrder) (uint16, error) {
	if len(data) < 2 {
		return 0, NewErrUnSufficientBytesError(2, 0)
	}
	return order.Uint16(data), nil
}

type UInt16Type[T any] struct {
	order  binary.ByteOrder
	getter getter[T, uint16]
	pooler Pooler
}

func (i UInt16Type[T]) Head() Head {
	return nil
}

func (i UInt16Type[T]) Fixed() bool {
	return true
}

func (i UInt16Type[T]) Length() int {
	return 2
}

func (i UInt16Type[T]) Parse(r io.Reader) (any, error) {
	return i.ParseUint16(r)
}

func (i UInt16Type[T]) ParseUint16(r io.Reader) (u16 uint16, err error) {
	u16, err = parseUInt16Reader(r, i.order, i.pooler)
	return
}

func (i UInt16Type[T]) ParseInt(r io.Reader) (int, error) {
	u16, err := parseUInt16Reader(r, i.order, i.pooler)
	if err != nil {
		return 0, err
	}
	return int(u16), nil
}

func (i UInt16Type[T]) Compile(item T, w io.Writer) error {
	u16 := i.getter(item)
	return i.CompileUInt16(u16, w)
}

func (i UInt16Type[T]) CompileUInt16(u16 uint16, w io.Writer) (err error) {
	bts := i.pooler.Get(2)
	defer i.pooler.Put(bts)
	data := *bts
	data = data[:2]
	i.order.PutUint16(data, u16)
	_, err = w.Write(data)

	// TODO: widen Writer interface
	//b.WriteByte(byte(u16))
	//b.WriteByte(byte(u16 >> 8))

	return
}

func (i UInt16Type[T]) CompileInt(in int, w io.Writer) error {
	if in > 65535 {
		return ErrUnSufficientBytes{want: in, have: 65535}
	}
	return i.CompileUInt16(uint16(in), w)
}

func (i UInt16Type[T]) ParseLength(data []byte) (int, error) {
	u16, err := parseUInt16(data, i.order)
	if err != nil {
		return 0, err
	}
	return int(u16), nil
}

func UInt16[T any](order binary.ByteOrder, getter getter[T, uint16]) UInt16Type[T] {
	return UInt16Type[T]{order: order, getter: getter, pooler: SinglePool}
}

func UInt16LE[T any](getter getter[T, uint16]) UInt16Type[T] {
	return UInt16[T](binary.LittleEndian, getter)
}

func UInt16BE[T any](getter getter[T, uint16]) UInt16Type[T] {
	return UInt16[T](binary.BigEndian, getter)
}

func UInt16Header[T any](order binary.ByteOrder) UInt16Type[T] {
	return UInt16Type[T]{order: order, pooler: SinglePool}
}

func UInt16LEHeader[T any]() UInt16Type[T] {
	return UInt16Header[T](binary.LittleEndian)
}

func UInt16BEHeader[T any]() UInt16Type[T] {
	return UInt16Header[T](binary.BigEndian)
}

func UInt16Body(order binary.ByteOrder) UInt16Type[uint16] {
	return UInt16Type[uint16]{getter: Identity[uint16], order: order, pooler: SinglePool}
}

func UInt16BodyLE() UInt16Type[uint16] {
	return UInt16Body(binary.LittleEndian)
}

func UInt16BodyBE() UInt16Type[uint16] {
	return UInt16Body(binary.BigEndian)
}

func AnyUInt16Body(order binary.ByteOrder) UInt16Type[any] {
	return UInt16Type[any]{
		pooler: SinglePool,
		order:  order,
		getter: func(x any) uint16 {
			return x.(uint16)
		},
	}
}
