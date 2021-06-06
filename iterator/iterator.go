package iterator

// ===
// 定义 Iterator 接口, 由 Aggregate 生成
// ===

type Iterator interface {
	HasNext() bool              // 判断当前是否还存在一个元素
	Next() (interface{}, error) // 得到当前元素, 并指向下一个元素
	// 当为 Next() 设置了回调函数, 那么就会涉及到另一种设计模式: Visitor 模式
	// 思考: 如果要求 Iterator 支持并发, 应该如何做呢?
}
