package iterator

// ===
// 定义 Aggregate 接口
// ===

type Aggregate interface {
	Iterator() Iterator // 定义 Iterator(), 返回同样实现了 Iterator 接口的实例
}
