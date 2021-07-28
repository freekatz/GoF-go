package html

import "abstract-factory/factory"

type HTMLFactory struct {
}

func (htmlFactory HTMLFactory) CreateLink(caption, url string) factory.Link {
	return NewHTMLLink(caption, url)
}

func (htmlFactory HTMLFactory) CreateTray(caption string) factory.Tray {
	return NewHTMLTray(caption)
}

func (htmlFactory HTMLFactory) CreatePage(title, author string) factory.Page {
	return NewHTMLPage(title, author)
}
