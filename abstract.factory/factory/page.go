package factory

/*
	抽象产品
*/

type Page interface {
	Add(item ...Item)
	Output(path string) error
	MakeHTML() string
}
