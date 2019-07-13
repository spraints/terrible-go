package enum

import "reflect"

type slice reflect.Value

func (s slice) Get() interface{} {
	return reflect.Value(s).Interface()
}

func (s slice) Map(fn interface{}) Enum {
	inVal := reflect.Value(s)
	inLen := inVal.Len()

	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()
	itemType := fnType.Out(0)

	outVals := make([]reflect.Value, 0, inLen)
	for i := 0; i < inLen; i++ {
		outVal := fnVal.Call([]reflect.Value{inVal.Index(i)})[0]
		outVals = append(outVals, outVal)
	}

	resVal := reflect.MakeSlice(reflect.SliceOf(itemType), 0, inLen)
	resVal = reflect.Append(resVal, outVals...)
	return slice(resVal)
}

func (s slice) Each(fn interface{}) {
	inVal := reflect.Value(s)
	inLen := inVal.Len()

	fnVal := reflect.ValueOf(fn)

	for i := 0; i < inLen; i++ {
		fnVal.Call([]reflect.Value{inVal.Index(i)})
	}
}
