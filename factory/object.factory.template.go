package factory

type ObjectFactoryTemplate struct {
	objectFactory ObjectFactory
}

func NewObjectFactoryTemplate(factory ObjectFactory) ObjectFactoryTemplate {
	return ObjectFactoryTemplate{
		objectFactory: factory,
	}
}

func (template ObjectFactoryTemplate) Create(id string) Object {
	object := template.objectFactory.createObject(id)
	template.objectFactory.registerObject(id, object)
	return object
}
