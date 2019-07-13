package enum

import (
	"reflect"
)

func (e *enum) Each(fn interface{}) {
	inVal := reflect.ValueOf(e.raw)
	if inVal.Kind() == reflect.Invalid || inVal.IsNil() {
		return
	}
	inLen := inVal.Len()
	if inLen == 0 {
		return
	}

	fnVal := reflect.ValueOf(fn)

	for i := 0; i < inLen; i++ {
		fnVal.Call([]reflect.Value{inVal.Index(i)})
	}
}
