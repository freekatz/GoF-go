package prototype

import (
	"fmt"
	"strings"
)

type PrototypeImpl struct {
	wrapper rune
}

func NewPrototypeImpl(w rune) PrototypeImpl {
	return PrototypeImpl{
		wrapper: w,
	}
}

func (impl PrototypeImpl) PrintWithWrapper(s string) {
	width := len(s)
	printLine(impl.wrapper, width)
	if width%2 == 0 {
		s += " "
	}
	fmt.Printf("%s %s %s\n", string(impl.wrapper), s, string(impl.wrapper))
	printLine(impl.wrapper, width)
}

func (impl PrototypeImpl) CreateClone() Prototype {
	proto := new(PrototypeImpl)
	proto.wrapper = impl.wrapper
	return proto
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
