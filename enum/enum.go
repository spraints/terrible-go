package enum

import "reflect"

// Enum is a wrapper around any type.
//
// You're responsible for making function types match. There is no type safety here.
type Enum interface {
	// Get returns the underlying slice from the current enum.
	Get() interface{}

	// Each applies a func to all elements of an enum.
	//
	//   enum.New([]int{1,2,3}).Each(func(x int) { log.Printf("val: %d", x) })
	Each(fn interface{})

	// Map transforms this enum into another.
	//
	//   enum.New([]int{1,2,3}).Map(func(x int) { return x * x }).Get()
	//   => []int{1,4,9}
	Map(transform interface{}) Enum

	// Select chooses elements from this enum.
	//
	//   enum.New([]int{1,2,3,4}).Select(func(x int) { return x % 2 == 0 }).Get()
	//   => []int{2,4}
	Select(predicate interface{}) Enum

	// Reduce creates a single value from this enum.
	//
	//   enum.New([]int{1,2,3}).Reduce(0, func(sum, x int) int { return sum + x })
	//   => 6
	Reduce(initial, reducer interface{}) interface{}

	// All returns true unless the predicate returns false for a value.
	//
	//   enum.New([]{1,2,3}).All(func(x int) { return x < 100 })
	//   => true
	All(predicate interface{}) bool

	// Any returns false unless the predicate returns true for a value.
	//
	//   enum.New([]{1,2,3}).Any(func(x int) { return x == 2 })
	//   => true
	Any(predicate interface{}) bool
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
