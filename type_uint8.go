package parco

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

func UInt8Header() Type[int] {
	return NewFixedType[int](
		1,
		ParseUInt8Header,
		CompileUInt8Header,
	)
}
