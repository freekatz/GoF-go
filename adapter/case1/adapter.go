package adapter_case1

import "fmt"

// ===
// 使用 Adapter 类型将 Banner 提供的实际情况适配到 Printer 的需求上面
// ===
type Adapter struct {
	*Banner // 继承 Banner
}

func NewAdapter() *Adapter {
	return &Adapter{
		Banner: NewBanner(),
	}
}

// 实现 Printer 接口
func (adapter *Adapter) PrintWeak(str string) {
	fmt.Println(adapter.WithParen(str))
}

// 实现 Printer 接口
func (adapter *Adapter) PrintStrong(str string) {
	fmt.Println(adapter.WithAster(str))
}
