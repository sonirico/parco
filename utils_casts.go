package parco

func AnyIntToInt(x any) (int, bool) {
	switch val := x.(type) {
	case int:
		return val, true
	case int8:
		return int(val), true
	case int16:
		return int(val), true
	case int32:
		return int(val), true
	case int64:
		return int(val), true
	case uint8:
		return int(val), true
	case uint16:
		return int(val), true
	case uint32:
		return int(val), true
	case uint64:
		return int(val), true
	}
	return 0, false
}

func Identity[T any](x T) T {
	return x
}
