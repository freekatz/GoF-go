package prototype

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	manager := NewManager()
	proto1 := NewPrototypeImpl([2]rune{'*', '$'})
	proto2 := NewPrototypeImpl([2]rune{'#', '$'})

	_ = manager.Register("1", proto1)
	_ = manager.Register("2", proto2)

	proto3, _ := manager.Create("1")
	proto4, _ := manager.Create("2")

	fmt.Println(proto1 == proto2)           // 应该输出 false
	fmt.Printf("%p %p\n", &proto1, &proto3) // 地址不同
	fmt.Println(proto1 == proto3)           // 应该输出 true
	fmt.Printf("%p %p\n", &proto2, &proto4) // 地址不同
	fmt.Println(proto2 == proto4)           // 应该输出 true

	proto1.PrintWithWrapper("Hello prototype 1")
	proto2.PrintWithWrapper("Hello prototype 2")
	proto3.PrintWithWrapper("Hello prototype 1")
	proto4.PrintWithWrapper("Hello prototype 2")
}
