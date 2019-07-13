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

func (emptyType) All(interface{}) bool {
	return true
}
