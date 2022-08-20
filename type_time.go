package parco

import (
	"encoding/binary"
	"io"
	"time"
)

var (
	noopTime       = time.Time{}
	timeByteLength = 8
)

type (
	TimeType struct {
		locationType OptionalType[string]
		pooler       Pooler
	}
)

func (t TimeType) ByteLength() int {
	return timeByteLength
}

func (t TimeType) Parse(r io.Reader) (time.Time, error) {
	box := t.pooler.Get(timeByteLength)
	defer t.pooler.Put(box)

	data := *box
	data = data[:timeByteLength]

	n, err := r.Read(data)
	if n != timeByteLength || err != nil {
		return noopTime, ErrCannotRead
	}

	tim, err := ParseTime(data, binary.LittleEndian)
	if err != nil {
		return noopTime, err
	}

	locationRaw, err := t.locationType.Parse(r)

	if err != nil || locationRaw == nil {
		return noopTime, err
	}

	loc, err := time.LoadLocation(*locationRaw)

	if err != nil {
		return noopTime, err
	}

	return tim.In(loc), nil
}

func (t TimeType) Compile(tt time.Time, w io.Writer) error {
	box := t.pooler.Get(timeByteLength)
	defer t.pooler.Put(box)

	data := *box
	data = data[:timeByteLength]

	err := CompileTime(tt, data, binary.LittleEndian)
	if err != nil {
		return err
	}

	if n, err := w.Write(data); err != nil || n != timeByteLength {
		return ErrCannotWrite
	}

	if loc := tt.Location(); loc != nil {
		return t.locationType.Compile(Ptr(loc.String()), w)
	}

	return nil
}

func CompileTime(t time.Time, box []byte, order binary.ByteOrder) error {
	return CompileInt64(t.UnixNano(), box, order)
}

func ParseTime(box []byte, order binary.ByteOrder) (time.Time, error) {
	i64, err := ParseInt64(box, order)
	if err != nil {
		return noopTime, err
	}

	return time.Unix(0, i64).UTC(), nil
}

func TimeLocation() Type[time.Time] {
	return TimeType{
		locationType: Option(SmallVarchar()),
		pooler:       SinglePool,
	}
}

func TimeUTC() Type[time.Time] {
	return NewFixedType[time.Time](
		timeByteLength,
		func(data []byte) (time.Time, error) {
			return ParseTime(data, binary.LittleEndian)
		},
		func(t time.Time, box []byte) error {
			return CompileTime(t, box, binary.LittleEndian)
		},
	)
}
