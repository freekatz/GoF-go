package template

import (
	"fmt"
	"strings"
)

type BytesDisplayer struct {
	bs []byte
}

func NewBytesDisplayer(bs []byte) *BytesDisplayer {
	return &BytesDisplayer{
		bs: bs,
	}
}

func (displayer *BytesDisplayer) Open() error {
	fmt.Print("[")
	return nil
}

func (displayer *BytesDisplayer) Print() error {
	s := strings.Builder{}
	for i := 0; i < len(displayer.bs)-1; i++ {
		b := displayer.bs[i]
		s.WriteByte(b)
		s.WriteRune(',')
	}
	s.WriteByte(displayer.bs[len(displayer.bs)-1])
	fmt.Print(s.String())
	return nil
}

func (displayer *BytesDisplayer) Close() error {
	fmt.Print("]", "\n")
	return nil
}
