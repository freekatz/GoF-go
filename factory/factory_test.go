package factory

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	simpleFactory := NewSimpleObjectFactory()
	factory := NewObjectFactoryTemplate(simpleFactory)
	o1 := factory.Create("o1")
	o2 := factory.Create("o2")
	o3 := factory.Create("o3")
	_ = o1.DoSomething()
	_ = o2.DoSomething()
	_ = o3.DoSomething()

	o4, err := simpleFactory.GetObject("o2")
	if err != nil {
		fmt.Println(err.Error())
	}
	if o2 == o4 {
		_ = o4.DoSomething()
	}
}
