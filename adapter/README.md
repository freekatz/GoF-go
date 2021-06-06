# Adapter 模式

Adapter 位于**实际情况**与**需求**之间, 用来填补两者之间的差异

Adapter 模式也称为 **Wrapper 模式**

包括两种类型:

- 类 Adapter 模式 (使用**继承**的适配器)
- 对象 Adapter 模式 (使用**委托**的适配器)

**继承**就是 `Adapter` 同时继承和实现需要适配的两方 (这两方一方是结构体, 另一方是接口), 达到适配的目的

**委托**就是将工作交给其他方, 也就是将某个方法的处理交给其它实例的方法 (也就是说, 两方都是结构体)

需要知道的是: Golang 是原生支持**委托**的, 因为同一个结构体可以混合多个结构体, 但是依然有必要按照对象 Adapter 模式进行适配 (这里面有一些不同, 看下面的 cases)

示例:

- [case1](./case1) 是一个简单的类 Adapter 模式的例子: 使用 `Adapter` 类型将 `Banner` 提供的实际情况适配到 `Printer` 的需求上面
- [case2](./case2) 是一个解决**如何部分实现一个接口, 而不引发 panic** 的例子, 这里使用类 Adapter 模式来解决, 是使用 Adapter 模式的一种常见特例
- [case3](./case3) 是一个简单的对象 Adapter 模式的例子: `Banner` 与 case1 完全相同, 把 `Printer` 改为结构体, `Adapter` 同时混合两个结构, 将 `Banner` 提供的实际情况适配到 `Printer` 的需求上面, 与 `Adapter` 同时继承 `Banner` 和 `Printer` 不同的是 **Adapter extends Printer, and Adapter only has Banner**

> Adaptor 模式用于连接 **API 不同**的类或接口, 而 Bridge 模式则用于连接类的**功能层次**与**实现层次**