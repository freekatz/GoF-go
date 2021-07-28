package factory

// 抽象零件
// Item 由 link 和 tray 实现，且必须继承自 ItemImpl（这里这样定义其实是间接实现了 Java 中的 abstract class）

type Item interface {
	MakeHTML() string
	SetCaption(caption string)
}
