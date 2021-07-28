package html

import (
	"abstract-factory/factory"
	"strings"
)

type HTMLTray struct {
	caption string
	items   []factory.Item
}

func NewHTMLTray(caption string) *HTMLTray {
	return &HTMLTray{
		caption: caption,
		items:   make([]factory.Item, 0),
	}
}

func (tray *HTMLTray) Add(item ...factory.Item) {
	tray.items = append(tray.items, item...)
}

func (tray *HTMLTray) MakeHTML() string {
	stringBuilder := new(strings.Builder)
	stringBuilder.WriteString("<li>\n")
	stringBuilder.WriteString(tray.caption + "\n")
	stringBuilder.WriteString("<ul>\n")

	for _, it := range tray.items {
		stringBuilder.WriteString(it.MakeHTML())
	}
	stringBuilder.WriteString("</ul>\n</li>\n")

	return stringBuilder.String()
}

func (tray *HTMLTray) SetCaption(caption string) {
	tray.caption = caption
}
