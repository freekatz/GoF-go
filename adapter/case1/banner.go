package adapter_case1

import "strings"

// ===
// Banner 提供两种包装 string 的方式: (string) 和 **string**
// ===

type Banner struct {
}

func NewBanner() *Banner {
	return &Banner{}
}

func (banner *Banner) WithParen(str string) string {
	return strings.Join([]string{"(", str, ")"}, "")
}

func (banner *Banner) WithAster(str string) string {
	return strings.Join([]string{"**", str, "**"}, "")
}
