package enum

// Enum is a wrapper around any type.
type Enum interface {
	// Get returns the underlying slice from the current enum.
	Get() interface{}

	Map(fn interface{}) Enum
	Each(fn interface{})
}

// New makes an Enum out of whatever slice you give it.
func New(raw interface{}) Enum {
	return &enum{raw}
}

type enum struct {
	raw interface{}
}

func (e *enum) Get() interface{} {
	return e.raw
}
