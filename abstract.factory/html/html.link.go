package html

import (
	"fmt"
)

type HTMLLink struct {
	caption string
	url     string
}

func NewHTMLLink(caption, url string) *HTMLLink {
	return &HTMLLink{
		caption: caption,
		url:     url,
	}
}

func (link *HTMLLink) MakeHTML() string {
	return fmt.Sprintf(`  <li><a href="%s">%s</a></li>`, link.url, link.caption)
}

func (link *HTMLLink) SetCaption(caption string) {
	link.caption = caption
}

func (link *HTMLLink) SetURL(url string) {
	link.url = url
}
