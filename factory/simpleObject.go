package factory

import "fmt"

type SimpleObject struct {
	ID string
}

func NewSimpleObject(id string) SimpleObject {
	return SimpleObject{
		ID: id,
	}
}

// 实现 Object 接口

func (object SimpleObject) DoSomething() error {
	fmt.Printf("%s do something\n", object.ID)
	return nil
}
