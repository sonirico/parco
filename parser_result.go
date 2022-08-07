package parco

//import (
//	"reflect"
//)
//
//type Result struct {
//	data map[string]structItem[any]
//}
//
//func (r *Result) getItem(k string) (item structItem[any], err error) {
//	item, ok := r.data[k]
//	if !ok {
//		return item, NewErrFieldNotFoundError(k)
//	}
//
//	return item, nil
//}
//
//func (r *Result) get(k string) (data any, err error) {
//	item, ok := r.data[k]
//	if !ok {
//		return nil, NewErrFieldNotFoundError(k)
//	}
//
//	return item.value, nil
//}
//
//func (r *Result) Get(k string) (data []byte, err error) {
//	item, ok := r.data[k]
//	if !ok {
//		return nil, NewErrFieldNotFoundError(k)
//	}
//
//	return item.data, nil
//
//}
//
//func (r *Result) GetString(k string) (string, error) {
//	value, err := r.get(k)
//	if err != nil {
//		return "", err
//	}
//
//	str, ok := value.(string)
//	if !ok {
//		return "", NewErrTypeAssertionError("string", reflect.TypeOf(value).String())
//	}
//	return str, nil
//}
//
//func (r *Result) GetUInt8(k string) (uint8, error) {
//	bts, err := r.get(k)
//	if err != nil {
//		return 0, err
//	}
//
//	str, ok := bts.(uint8)
//	if !ok {
//		return 0, NewErrTypeAssertionError("uint8", reflect.TypeOf(bts).String())
//	}
//	return str, nil
//}
//
//func (r *Result) GetArray(k string) (arr ArrayValue[any], err error) {
//	item, err := r.getItem(k)
//	if err != nil {
//		return
//	}
//
//	var ok bool
//	arr, ok = item.value.(ArrayValue[any])
//	if !ok {
//		err = NewErrTypeAssertionError("array", reflect.TypeOf(item.value).String())
//		return
//	}
//	return arr, nil
//
//}
//
//func newResult() Result {
//	return Result{
//		data: make(map[string]structItem[any]),
//	}
//}
//
