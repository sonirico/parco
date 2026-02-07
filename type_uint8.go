package parco

import "math"

func CompileUInt8(u8 uint8, box []byte) (err error) {
	box[0] = u8
	return
}

func ParseUInt8(data []byte) (uint8, error) {
	if len(data) < 1 {
		return 0, NewErrUnSufficientBytesError(1, 0)
	}
	return data[0], nil
}

func CompileUInt8Header(i int, box []byte) (err error) {
	if i > 255 || i < 0 {
		err = ErrOverflow
		return
	}
	box[0] = byte(i)
	return
}

func ParseUInt8Header(data []byte) (int, error) {
	if len(data) < 1 {
		return 0, NewErrUnSufficientBytesError(1, 0)
	}
	return int(data[0]), nil
}

func UInt8() Type[uint8] {
	return NewFixedType[uint8](
		1,
		ParseUInt8,
		CompileUInt8,
	)
}

func Byte() Type[byte] {
	return UInt8()
}

func UInt8Header() Type[int] {
	return NewFixedType[int](
		1,
		ParseUInt8Header,
		CompileUInt8Header,
	)
}

func CompileInt8(i8 int8, box []byte) (err error) {
	return CompileUInt8(uint8(i8), box)
}

func ParseInt8(data []byte) (int8, error) {
	if len(data) < 1 {
		return 0, NewErrUnSufficientBytesError(1, 0)
	}
	// Direct conversion handles two's complement for negative values
	return int8(data[0]), nil
}

func CompileInt8Header(i int, box []byte) (err error) {
	if i > math.MaxInt8 || i < math.MinInt8 {
		err = ErrOverflow
		return
	}
	box[0] = byte(i)
	return
}

func ParseInt8Header(box []byte) (int, error) {
	x, err := ParseInt8(box)
	if err != nil {
		return 0, err
	}
	return int(x), nil
}

func Int8() Type[int8] {
	return NewFixedType[int8](
		1,
		ParseInt8,
		CompileInt8,
	)
}

func Int8Header() Type[int] {
	return NewFixedType[int](
		1,
		ParseInt8Header,
		CompileInt8Header,
	)
}

func Bool() Type[bool] {
	return NewFixedType[bool](
		1,
		func(data []byte) (bool, error) {
			n, err := ParseUInt8(data)
			if err != nil {
				return false, err
			}

			return n == 1, err
		},
		func(value bool, box []byte) (err error) {
			var n uint8
			if value {
				n = 1
			}

			return CompileUInt8(n, box)
		},
	)
}
