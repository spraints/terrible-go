package enum

import "reflect"

// Enum is a wrapper around any type.
type Enum interface {
	// Get returns the underlying slice from the current enum.
	Get() interface{}

	Map(fn interface{}) Enum
	Each(fn interface{})
	All(predicate interface{}) bool
}

// New makes an Enum out of whatever slice you give it.
func New(raw interface{}) Enum {
	val := reflect.ValueOf(raw)
	switch val.Kind() {
	case reflect.Slice:
		if val.Len() == 0 {
			return empty
		}
		return slice{val}
	default: // Or maybe Invalid?
		return empty
	}
}
