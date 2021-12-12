package internal

type ranger func(x interface{}) error

type (
	u8Iter  []uint8
	u16Iter []uint16
)

func (i u8Iter) Range(fn ranger) error {
	for _, x := range i {
		if err := fn(x); err != nil {
			return err
		}
	}

	return nil
}
func (i u8Iter) Len() int {
	return len(i)
}

func UInt8Iter(x []uint8) u8Iter {
	return u8Iter(x)
}

func (i u16Iter) Range(fn ranger) error {
	for _, x := range i {
		if err := fn(x); err != nil {
			return err
		}
	}

	return nil
}
func (i u16Iter) Len() int {
	return len(i)
}

func UInt16Iter(x []uint16) u16Iter {
	return u16Iter(x)
}
