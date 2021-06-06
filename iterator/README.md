# Iterator 模式

Iterator 模式适合用在, 当循环遍历时, 只有**少数**字段**有规律更改**的情况

由 `Aggregate` 和 `Iterator` 两个接口组成

`Aggregate` 提供 `Iterator()` 来生成 `Iterator`

`Iterator` 一般包括 `HasNext()` 和 `Next()` 来进行迭代

> 定义生成 Iterator 的方法, 可能会用到另一种设计模式: Factory Method 模式
> 当为 Next() 设置了回调函数, 那么就会涉及到另一种设计模式: Visitor 模式

特别注意: `Aggregate` 和 `Iterator` 是一一对应的, `Iterator` 只能由 `Aggregate` 生成, 尽量不要贪图简单而违背模式

实现过程：

Aggregate定义 => Iterator定义 => Aggregate实现 => Iterator实现