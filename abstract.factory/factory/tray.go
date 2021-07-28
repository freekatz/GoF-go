package factory

/*
	抽象零件
	Link 实现 Item
*/

type Tray interface {
	Item

	Add(item ...Item)
}
