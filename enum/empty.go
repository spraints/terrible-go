package enum

type emptyType struct{}

var empty emptyType

func (emptyType) Get() interface{} {
	return nil
}

func (emptyType) Each(interface{}) {
}

func (emptyType) Map(interface{}) Enum {
	return empty
}

func (emptyType) Select(interface{}) Enum {
	return empty
}

func (emptyType) Reject(interface{}) Enum {
	return empty
}

func (emptyType) Reduce(init, _ interface{}) interface{} {
	return init
}

func (emptyType) All(interface{}) bool {
	return true
}

func (emptyType) Any(interface{}) bool {
	return false
}
