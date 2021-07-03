package case1

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	txtBuilder := TXTBuilder{}
	htmlBuilder := HTMLBuilder{}
	txtDirector := NewDirector(&txtBuilder)
	htmlDirector := NewDirector(&htmlBuilder)

	txtDirector.Build("txt", "txt body")
	htmlDirector.Build("html", "html body")

	txt := txtBuilder.GetResult()
	html := htmlBuilder.GetResult()

	if len(txt) == 0 || len(html) == 0 {
		t.Error("txt or html len is zero", txt, html)
	}

	log.Println(txt, html)
}
