package iterator

// ===
// 定义 Iterator 接口, 由 Aggregate 生成
// ===

type Iterator interface {
	HasNext() bool              // 判断是否存在下一个元素
	Next() (interface{}, error) // 得到下一个元素
}
