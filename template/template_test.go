package template

import "testing"

func Test(t *testing.T) {
	s := "Hello"
	stringDisplayer := NewDisplayerTemplate(NewStringDisplayer(s))
	bytesDisplayer := NewDisplayerTemplate(NewBytesDisplayer([]byte(s)))
	stringDisplayer.Display()
	bytesDisplayer.Display()
}
