# Abstract Factory Method 模式

本例中，factory 包定义了**抽象零件**，**抽象产品**以及**抽象工厂**
- 抽象零件： `Item` 接口, `Link` 和 `Tray` 接口继承自 `Item` 接口
- 抽象产品： `Page` 接口
- 抽象工厂： 创建抽象零件和抽象产品

html 包实现了具体的零件，产品和工厂

抽象工厂的优点：新增具体工厂很简单
抽象工厂的缺点：新增抽象零件很繁琐

> Builder 模式： 抽象工厂通过调用抽象产品的 API 来组装抽象产品，生成复杂实例； Builder 则分阶段地制作复杂实例；[builder/case2](https://github.com/1uvu/GoF-go/tree/main/builder/case2) 其实是将抽象工厂和 builder 理念进行融合的例子
> Factory Method 模式： 抽象工厂有时会使用到工厂方法来生成具体实例，如本例中 `HTMLFactory` 实现 `Factory` 接口，使用了 `NewHTML*()` 模板方法
> Composite 模式： 有时抽象工厂制作产品时会用到 Composite 模式
> Singleton 模式： 具体工厂实例一般是单例的
