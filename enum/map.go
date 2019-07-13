package enum

import (
	"reflect"
)

func (e *enum) Map(fn interface{}) Enum {
	inVal := reflect.ValueOf(e.raw)
	if inVal.Kind() == reflect.Invalid || inVal.IsNil() {
		return &enum{}
	}
	inLen := inVal.Len()
	if inLen == 0 {
		return &enum{}
	}

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
	return &enum{raw: resVal.Interface()}
}
