package prototype

type Prototype interface {
	PrintWithWrapper(s string)
	CreateClone() Prototype
}
