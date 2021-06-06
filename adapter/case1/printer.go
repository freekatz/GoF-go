package adapter_case1

// ===
// Printer 定义两种打印 string 的输出结构: (string) 和 **string**
// ===

type Printer interface {
	PrintWeak(string)
	PrintStrong(string)
}
