package pkg

import (
	"reflect"

	"github.com/sonirico/parco/internal"
)

type Result struct {
	data map[string]structItem
}

func (r *Result) getItem(k string) (item structItem, err error) {
	item, ok := r.data[k]
	if !ok {
		return structItem{}, internal.NewErrFieldNotFoundError(k)
	}

	return item, nil
}

func (r *Result) get(k string) (data interface{}, err error) {
	item, ok := r.data[k]
	if !ok {
		return nil, internal.NewErrFieldNotFoundError(k)
	}

	return item.value, nil
}

func (r *Result) Get(k string) (data []byte, err error) {
	item, ok := r.data[k]
	if !ok {
		return nil, internal.NewErrFieldNotFoundError(k)
	}

	return item.data, nil

}

func (r *Result) GetString(k string) (string, error) {
	value, err := r.get(k)
	if err != nil {
		return "", err
	}

	str, ok := value.(string)
	if !ok {
		return "", internal.NewErrTypeAssertionError("string", reflect.TypeOf(value).String())
	}
	return str, nil
}

func (r *Result) GetUInt8(k string) (uint8, error) {
	bts, err := r.get(k)
	if err != nil {
		return 0, err
	}

	str, ok := bts.(uint8)
	if !ok {
		return 0, internal.NewErrTypeAssertionError("uint8", reflect.TypeOf(bts).String())
	}
	return str, nil
}

func (r *Result) GetArray(k string) (internal.ArrayValue, error) {
	item, err := r.getItem(k)
	if err != nil {
		return internal.NoopArrVal, err
	}

	val, ok := item.value.(internal.ArrayValue)
	if !ok {
		return internal.NoopArrVal, internal.NewErrTypeAssertionError("array", reflect.TypeOf(item.value).String())
	}
	return val, nil

}

func newResult() Result {
	return Result{
		data: make(map[string]structItem),
	}
}
