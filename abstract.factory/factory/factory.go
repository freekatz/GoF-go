package factory

/*
	抽象工厂
*/

type Factory interface {
	CreateLink(caption, url string) Link
	CreateTray(caption string) Tray
	CreatePage(title, author string) Page
}
