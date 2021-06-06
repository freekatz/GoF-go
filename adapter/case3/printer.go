package adapter_case1

// ===
// Printer 实现两种打印 string 的输出结构: (string) 和 **string**
// ===

type Printer struct {
}

func NewPrinter() *Printer {
	return &Printer{}
}

func (printer *Printer) PrintWeak(str string) {
}

func (printer *Printer) PrintStrong(str string) {
}
