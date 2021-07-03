package case1

import (
	"bytes"
)

type HTMLBuilder struct {
	htmlBuffer bytes.Buffer
}

func (builder *HTMLBuilder) MakeHead() {
	builder.htmlBuffer.WriteByte('\n')
	builder.htmlBuffer.WriteString("<head>Txt Head</head>")
}

func (builder *HTMLBuilder) MakeTitle(title string) {
	builder.htmlBuffer.WriteByte('\n')
	builder.htmlBuffer.WriteString("<title>")
	builder.htmlBuffer.WriteString(title)
	builder.htmlBuffer.WriteString("</title>")
}

func (builder *HTMLBuilder) MakeBody(body string) {
	builder.htmlBuffer.WriteByte('\n')
	builder.htmlBuffer.WriteString("<body>\n\t")
	builder.htmlBuffer.WriteString(body)
	builder.htmlBuffer.WriteString("\n</body>")
}

func (builder *HTMLBuilder) MakeFooter() {
	builder.htmlBuffer.WriteByte('\n')
	builder.htmlBuffer.WriteString("<footer>Txt Footer</footer>")
}

func (builder *HTMLBuilder) GetResult() string {
	return builder.htmlBuffer.String()
}
