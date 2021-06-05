package iterator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	og := NewObjectGroup(10)
	_ = og.AddObject(NewObject(0))
	_ = og.AddObject(NewObject(1))
	_ = og.AddObject(NewObject(2))
	_ = og.AddObject(NewObject(3))

	iter := og.Iterator()
	for iter.HasNext() {
		o, _ := iter.Next()
		fmt.Println(o)
	}

}
