package factory

type ObjectFactory interface {
	createObject(id string) Object
	registerObject(id string, object Object)
}
