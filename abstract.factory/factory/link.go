package factory

/*
	抽象零件
	Link 实现 Item
*/

type Link interface {
	Item

	SetURL(url string)
}
