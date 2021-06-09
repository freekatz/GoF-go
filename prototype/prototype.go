package prototype

type Prototype interface {
	PrintWithWrapper(s string)
	CreateClone() Prototype // 根据实际需要选择浅拷贝还是深拷贝
}
