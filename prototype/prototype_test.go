package prototype

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	manager := NewManager()
	proto1 := NewPrototypeImpl('*')
	proto2 := NewPrototypeImpl('#')

	_ = manager.Register("*", proto1)
	_ = manager.Register("#", proto2)

	proto3, _ := manager.Create("*")
	proto4, _ := manager.Create("#")

	fmt.Println(proto1 == proto3) // 应该输出 false
	fmt.Println(proto2 == proto4) // 应该输出 false

	proto1.PrintWithWrapper("Hello prototype 1")
	proto2.PrintWithWrapper("Hello prototype 2")
	proto3.PrintWithWrapper("Hello prototype 1")
	proto4.PrintWithWrapper("Hello prototype 2")
}
