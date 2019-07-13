package enum

import "reflect"

type slice struct {
	reflect.Value
}

func (s slice) Get() interface{} {
	return s.Interface()
}

func (s slice) Map(fn interface{}) Enum {
	fnVal := reflect.ValueOf(fn)
	fnType := fnVal.Type()
	itemType := fnType.Out(0)

	inLen := s.Len()

	outVals := make([]reflect.Value, 0, inLen)
	eachValue(s.Value, func(item reflect.Value) {
		outVal := fnVal.Call([]reflect.Value{item})[0]
		outVals = append(outVals, outVal)
	})

	resVal := reflect.MakeSlice(reflect.SliceOf(itemType), 0, inLen)
	resVal = reflect.Append(resVal, outVals...)
	return slice{resVal}
}

func (s slice) Each(fn interface{}) {
	fnVal := reflect.ValueOf(fn)

	eachValue(s.Value, func(item reflect.Value) {
		fnVal.Call([]reflect.Value{item})
	})
}

func (s slice) All(predicate interface{}) bool {
	fnVal := reflect.ValueOf(predicate)

	all := true
	eachValue(s.Value, func(item reflect.Value) {
		outVal := fnVal.Call([]reflect.Value{item})[0]
		all = all && outVal.Bool()
	})

	return all
}

func eachValue(slice reflect.Value, fn func(reflect.Value)) {
	inLen := slice.Len()

	for i := 0; i < inLen; i++ {
		fn(slice.Index(i))
	}
}
