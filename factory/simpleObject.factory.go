package factory

import "fmt"

type SimpleObjectFactory struct {
	objectMap map[string]Object
}

func NewSimpleObjectFactory() SimpleObjectFactory {
	return SimpleObjectFactory{
		objectMap: make(map[string]Object),
	}
}

// 实现 ObjectFactory 接口

func (factory SimpleObjectFactory) createObject(id string) Object {
	return NewSimpleObject(id)
}

func (factory SimpleObjectFactory) registerObject(id string, object Object) {
	factory.objectMap[id] = object
}

// 额外实现的方法
func (factory SimpleObjectFactory) GetObject(id string) (Object, error) {
	object, ok := factory.objectMap[id]
	if !ok {
		return nil, fmt.Errorf("id: %s do not exist\n", id)
	}
	return object, nil
}
