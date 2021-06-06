package adapter_case2

import "fmt"

// ===
// ImplementedObjectAdapter 混合了 UnimplementedObjectAdapter
// 只部分实现了 Object 接口的 Secret() 方法
// ===

type ImplementedObject struct {
	*UnimplementedObjectAdapter // 混合一个已实现全部接口方法的 UnimplementedObjectAdapter

	ID int
}

func NewImplementedObject(id int) *ImplementedObject {
	return &ImplementedObject{
		ID: id,
	}
}

func (o *ImplementedObject) Secret() SecretType {
	switch {
	case o.ID%3 == 0:
		return Love
	case o.ID%3 == 1:
		return Life
	default:
		return NoneSecret
	}
}

// // Info()  并没有实现, 也就是 Object 接口并没有完全实现
// func (o *ImplementedObject) Info() string {
// 	return ""
// }

func (o *ImplementedObject) String() string {
	return fmt.Sprintf("Object with ID: %d, has secert called: %s.\n", o.ID, o.Secret().String())
}
