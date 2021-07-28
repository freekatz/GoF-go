package main

import (
	"abstract-factory/factory"
	"abstract-factory/html"
	"fmt"
)

var factoryMap = map[string]factory.Factory{
	"html": html.HTMLFactory{},
}

func GetFactory(factoryName string) factory.Factory {
	if v, ok := factoryMap[factoryName]; ok {
		return v
	}

	return nil
}

func main() {
	htmlFactory := GetFactory("html")

	link1 := htmlFactory.CreateLink("link1", "http://link1.com")
	link2 := htmlFactory.CreateLink("link2", "http://link2.com")
	link3 := htmlFactory.CreateLink("link3", "http://link3.com")

	tray1 := htmlFactory.CreateTray("tray1")
	tray1.Add(link1)
	tray1.Add(link2)

	tray2 := htmlFactory.CreateTray("tray2")
	tray2.Add(link2)
	tray2.Add(link3)

	tray3 := htmlFactory.CreateTray("tray3")
	tray3.Add(tray2)
	tray3.Add(link3)

	page := htmlFactory.CreatePage("page", "zjh")
	page.Add(tray1)
	page.Add(tray3)

	fmt.Println(page.MakeHTML())

	_ = page.Output("./")
}
