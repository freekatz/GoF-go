package prototype

import (
	"fmt"
	"strings"
)

type PrototypeImpl struct {
	wrapper [2]rune
}

func NewPrototypeImpl(ws [2]rune) PrototypeImpl {
	return PrototypeImpl{
		wrapper: ws,
	}
}

func (impl PrototypeImpl) PrintWithWrapper(s string) {
	width := len(s)
	printLine(impl.wrapper[0], width)
	if width%2 == 0 {
		s += " "
	}
	fmt.Printf("%s %s %s\n", string(impl.wrapper[1]), s, string(impl.wrapper[1]))
	printLine(impl.wrapper[0], width)
}

func (impl PrototypeImpl) CreateClone() Prototype {
	return NewPrototypeImpl(impl.wrapper)
}

func printLine(w rune, width int) {
	sb := strings.Builder{}
	for i := 0; i < width/2+2; i++ {
		sb.WriteRune(w)
		sb.WriteRune(' ')
	}
	sb.WriteRune(w)
	sb.WriteRune('\n')
	fmt.Print(sb.String())
}
