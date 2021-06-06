# Iterator 模式

Iterator 模式适合用在, 当循环遍历时, 只有**少数**字段**有规律更改**的情况

由 Aggregate 和 Iterator 两个接口组成

特别注意: Aggregate 和 Iterator 是一一对应的, Iterator 只能由 Aggregate 生成, 尽量不要贪图简单而违背模式

实现过程：

Aggregate定义 => Iterator定义 => Aggregate实现 => Iterator实现