package case1

type Builder interface {
	MakeHead()
	MakeTitle(string)
	MakeBody(string)
	MakeFooter()
	GetResult() string
}
