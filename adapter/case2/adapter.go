package adapter_case2

// ===
// UnimplementedObjectAdapter 是一个默认的完全实现接口方法的结构体, 是一个 Adapter
// 这种情况是 case1 的一种特例
// case1 中 Adapter 要适配不同类型的结构体和接口
// 而 case2 则适配属于同一种类型 Object 的结构体和接口
// ===

type UnimplementedObjectAdapter struct {
}

func (o *UnimplementedObjectAdapter) mustEmbedUnimplementedObject() {
}

func (o *UnimplementedObjectAdapter) Secret() SecretType {
	return NoneSecret
}

func (o *UnimplementedObjectAdapter) Info() string {
	return "Info has not been implemented."
}
