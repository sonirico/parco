package parco

func SkipType(pad int) Type[any] {
	return fixedType[any]{
		byteLength: pad,
		parser: func(_ []byte) (any, error) {
			return nil, nil
		},
		compiler: func(_ any, box []byte) error {
			// Zero out the padding bytes
			for i := range box {
				box[i] = 0
			}
			return nil
		},
		pool: SinglePool,
	}
}

func SkipTypeFactory(pad int) Type[any] {
	return SkipType(pad)
}
