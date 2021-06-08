# Singleton 模式

Singleton 模式限制程序执行过程中只会存在一个实例, 理念很简单, 但是实现不简单

本实例中 [case1](./case1) 和 [case2](./case2) 展示了一个非常简单的 Singleton 示例, 它的目的是表达 Singleton 模式的主要思想:
- 基于定义全局变量来实现程序加载时自动创建实例 (保证只创建一次)
- 基于

其中 case1 的实现是线程不安全的, case2 的实现是线程安全的

假设当前 `singleton == nil`, 此时有两个 goroutine 同时调用 `GetInstance()`, 那么就会创建出两个实例来

所以在 [case3](./case3) 中展示了使用` sync.Once` 改进 case1 来实现线程安全的 Singleton 模式

思考一下, Once 应该加在何处呢? 

很显然, 我们的目的是**同时只能有一个线程来获取实例**, 所以只需加在 `GetInstance()` 中即可

需要注意几点:

- 必须是 `newSingleton()` 而不是 `NewSingleton()`, 防止意外调用生成其它实例
- 除了使用 Once 来限制意外, 还有其它基于锁的方法, 要求是要保证 `GetInstance()` 的原子性
- 在实际应用场景下 case1 和 case3 都是可用的

如下模式经常会用到 Singleton 模式实现仅生成一个实例

> Abstract Factory 模式 
> 
> Builder 模式
> 
> Facade 模式
> 
> Prototype 模式
> 
