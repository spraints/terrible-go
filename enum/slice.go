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

	return sliceFrom(itemType, outVals)
}

func (s slice) Select(predicate interface{}) Enum {
	fnVal := reflect.ValueOf(predicate)

	inLen := s.Len()
	selectedVals := make([]reflect.Value, 0, inLen)

	eachValue(s.Value, func(item reflect.Value) bool {
		if fnVal.Call([]reflect.Value{item})[0].Bool() {
			selectedVals = append(selectedVals, item)
		}
		return true
	})

	return sliceFrom(s.Type().Elem(), selectedVals)
}

func (s slice) Reject(predicate interface{}) Enum {
	fnVal := reflect.ValueOf(predicate)

	inLen := s.Len()
	rejectedVals := make([]reflect.Value, 0, inLen)

	eachValue(s.Value, func(item reflect.Value) bool {
		if !fnVal.Call([]reflect.Value{item})[0].Bool() {
			rejectedVals = append(rejectedVals, item)
		}
		return true
	})

	return sliceFrom(s.Type().Elem(), rejectedVals)
}

func (s slice) Partition(predicate interface{}) (Enum, Enum) {
	fnVal := reflect.ValueOf(predicate)

	inLen := s.Len()
	selectedVals := make([]reflect.Value, 0, inLen)
	rejectedVals := make([]reflect.Value, 0, inLen)

	eachValue(s.Value, func(item reflect.Value) bool {
		if fnVal.Call([]reflect.Value{item})[0].Bool() {
			selectedVals = append(selectedVals, item)
		} else {
			rejectedVals = append(rejectedVals, item)
		}
		return true
	})

	itemType := s.Type().Elem()
	return sliceFrom(itemType, selectedVals), sliceFrom(itemType, rejectedVals)
}

func (s slice) Reduce(init, reducer interface{}) interface{} {
	fnVal := reflect.ValueOf(reducer)
	memo := reflect.ValueOf(init)

	eachValue(s.Value, func(item reflect.Value) bool {
		memo = fnVal.Call([]reflect.Value{memo, item})[0]
		return true
	})

	return memo.Interface()
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

func sliceFrom(itemType reflect.Type, vals []reflect.Value) slice {
	resVal := reflect.MakeSlice(reflect.SliceOf(itemType), 0, len(vals))
	resVal = reflect.Append(resVal, vals...)
	return slice{resVal}
}
