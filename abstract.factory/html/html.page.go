package html

import (
	"abstract-factory/factory"
	"os"
	"path/filepath"
	"strings"
)

type HTMLPage struct {
	title   string
	author  string
	content []factory.Item
}

func NewHTMLPage(title, author string) *HTMLPage {
	return &HTMLPage{
		title:   title,
		author:  author,
		content: make([]factory.Item, 0),
	}
}

func (page *HTMLPage) Add(item ...factory.Item) {
	page.content = append(page.content, item...)
}

func (page *HTMLPage) Output(path string) error {

	filename := page.title + ".html"

	err := os.WriteFile(filepath.Join(path, filename), []byte(page.MakeHTML()), os.FileMode(0770))

	if err != nil {
		return err
	}

	return nil
}

func (page *HTMLPage) MakeHTML() string {
	stringBuilder := new(strings.Builder)
	stringBuilder.WriteString("<html>\n<head>\n<title>" + page.title + "</title>\n</head>\n")
	stringBuilder.WriteString("<body>\n")
	stringBuilder.WriteString("<h1>" + page.title + "</h1>\n")
	stringBuilder.WriteString("<ul>\n")

	for _, it := range page.content {
		stringBuilder.WriteString(it.MakeHTML())
	}
	stringBuilder.WriteString("</ul>\n")
	stringBuilder.WriteString("<hr><address>" + page.author + "</address>\n")
	stringBuilder.WriteString("</body>\n</html>\n")

	return stringBuilder.String()
}
