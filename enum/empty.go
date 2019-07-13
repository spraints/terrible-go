package enum

type emptyType struct{}

var empty emptyType

func (emptyType) Get() interface{} {
	return nil
}

func (emptyType) Map(fn interface{}) Enum {
	return empty
}

func (emptyType) Each(fn interface{}) {
}

func (emptyType) All(interface{}) bool {
	return true
}
