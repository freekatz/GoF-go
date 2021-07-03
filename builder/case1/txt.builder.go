package case1

import (
	"bytes"
)

type TXTBuilder struct {
	txtBuffer bytes.Buffer
}

func (builder *TXTBuilder) MakeHead() {
	builder.txtBuffer.WriteByte('\n')
	builder.txtBuffer.WriteString("===TXT Head===")
}

func (builder *TXTBuilder) MakeTitle(title string) {
	builder.txtBuffer.WriteByte('\n')
	builder.txtBuffer.WriteString("TXT title:")
	builder.txtBuffer.WriteString(title)
}

func (builder *TXTBuilder) MakeBody(body string) {
	builder.txtBuffer.WriteByte('\n')
	builder.txtBuffer.WriteString("TXT body:\n\t")
	builder.txtBuffer.WriteString(body)
}

func (builder *TXTBuilder) MakeFooter() {
	builder.txtBuffer.WriteByte('\n')
	builder.txtBuffer.WriteString("==TXT Footer==")
}

func (builder *TXTBuilder) GetResult() string {
	return builder.txtBuffer.String()
}
