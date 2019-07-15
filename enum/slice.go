package enum

import "reflect"

type slice struct {
	reflect.Value
}

func (s slice) Get() interface{} {
	return s.Interface()
}

func (s slice) Each(fn interface{}) {
	fnVal := reflect.ValueOf(fn)

	eachValue(s.Value, func(item reflect.Value) bool {
		fnVal.Call([]reflect.Value{item})
		return true
	})
}

func (s slice) Map(transform interface{}) Enum {
	fnVal := reflect.ValueOf(transform)
	fnType := fnVal.Type()
	itemType := fnType.Out(0)

	inLen := s.Len()

	outVals := make([]reflect.Value, 0, inLen)
	eachValue(s.Value, func(item reflect.Value) bool {
		outVal := fnVal.Call([]reflect.Value{item})[0]
		outVals = append(outVals, outVal)
		return true
	})

	resVal := reflect.MakeSlice(reflect.SliceOf(itemType), 0, inLen)
	resVal = reflect.Append(resVal, outVals...)
	return slice{resVal}
}

func (s slice) All(predicate interface{}) bool {
	fnVal := reflect.ValueOf(predicate)

	result := true
	eachValue(s.Value, func(item reflect.Value) bool {
		outVal := fnVal.Call([]reflect.Value{item})[0]
		result = outVal.Bool()
		return result // continue iterating as long as the result is still true.
	})

	return result
}

func (s slice) Any(predicate interface{}) bool {
	fnVal := reflect.ValueOf(predicate)

	result := false
	eachValue(s.Value, func(item reflect.Value) bool {
		outVal := fnVal.Call([]reflect.Value{item})[0]
		result = outVal.Bool()
		return !result // continue iterating as long as the result is still false.
	})

	return result
}

func eachValue(slice reflect.Value, fn func(reflect.Value) bool) {
	inLen := slice.Len()

	for i := 0; i < inLen; i++ {
		if !fn(slice.Index(i)) {
			break
		}
	}
}
