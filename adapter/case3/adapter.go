package adapter_case1

import "fmt"

// ===
// 使用 Adapter 类型将 Banner 提供的实际情况适配到 Printer 的需求上面
// ===
type Adapter struct {
	*Printer         // 继承 Printer
	banner   *Banner // 保存一个 Banner
	// 这里的保存, 可以直接选择混合继承 Banner 来做替代, 本实例里面是严格按照模式来做的
}

func NewAdapter() *Adapter {
	return &Adapter{
		Printer: NewPrinter(),
		banner:  NewBanner(),
	}
}

func (adapter *Adapter) PrintWeak(str string) {
	fmt.Println(adapter.banner.WithParen(str))
}

func (adapter *Adapter) PrintStrong(str string) {
	fmt.Println(adapter.banner.WithAster(str))
}
