package enum

import "reflect"

type slice reflect.Value

func (s slice) Get() interface{} {
	return reflect.Value(s).Interface()
}

func (s slice) Map(fn interface{}) Enum {
	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()
	itemType := fnType.Out(0)

	inLen := reflect.Value(s).Len()

	outVals := make([]reflect.Value, 0, inLen)
	eachValue(reflect.Value(s), func(item reflect.Value) {
		outVal := fnVal.Call([]reflect.Value{item})[0]
		outVals = append(outVals, outVal)
	})

	resVal := reflect.MakeSlice(reflect.SliceOf(itemType), 0, inLen)
	resVal = reflect.Append(resVal, outVals...)
	return slice(resVal)
}

func (s slice) Each(fn interface{}) {
	fnVal := reflect.ValueOf(fn)

	eachValue(reflect.Value(s), func(item reflect.Value) {
		fnVal.Call([]reflect.Value{item})
	})
}

func eachValue(slice reflect.Value, fn func(reflect.Value)) {
	inLen := slice.Len()

	for i := 0; i < inLen; i++ {
		fn(slice.Index(i))
	}
}
