package template

import (
	"fmt"
	"strings"
)

type StringDisplayer struct {
	str      string
	strWidth int
}

func NewStringDisplayer(s string) *StringDisplayer {
	return &StringDisplayer{
		str:      s,
		strWidth: len(s),
	}
}

func (displayer *StringDisplayer) Open() error {
	displayer.printLine()
	return nil
}

func (displayer *StringDisplayer) Print() error {
	fmt.Println(strings.Join([]string{"|", displayer.str, "|"}, ""))
	return nil
}

func (displayer *StringDisplayer) Close() error {
	displayer.printLine()
	return nil
}

func (displayer *StringDisplayer) printLine() {
	line := strings.Builder{}
	line.WriteRune('+')
	for i := 0; i < displayer.strWidth; i++ {
		line.WriteRune('-')
	}
	line.WriteRune('+')
	fmt.Println(line.String())
}
