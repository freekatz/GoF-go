# Builder 模式

## case1

[case1](./case1) 是一个简单的标准的 Builder 模式示例，Builder 接口定义构建方法，Director 实现约束构建方法调用顺序

## case2

[case2](./case2) 展示了一个更加灵活的 Builder 模式示例，它不依赖 builder 方法的调用顺序

它用来构建 `Ethernet`, `IPv4` 和 `IPv6` 数据包的例子 (省略了很多详细字段).

`PacketBuilder` 接口定义构建相关的方法

`Packet` 接口定义属性字段及其获取方法

这里为了灵活性, 并未像其他 Builder 模式实现一样, 一次性构建完成, 这里定义了各种构建方法, 而没有严格要求构建方法调用的顺序, 且方法默认返回当前 `builder`, 以便**流式构建**.

当然是否定义顺序要根据具体场景来确定, 如需要控制构建方法顺序, 则参考 case1，否则想要获取更大的灵活性，参考 case2

