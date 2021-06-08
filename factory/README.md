# Factory Method 模式

Factory Method 模式是 Template Method 模式的一个特例, 即在模板方法规定实例生成的逻辑流程

父类决定要生成实例的方式, 而不决定生成的具体类型, 这些处理都交给子类进行

本实例中, `Object` 接口定义 `DoSomething()`, `ObjectFactory` 接口定义生成 `Object` 的方法, `ObjectFactoryTemplate` 实现生成 `Object` 的模板方法, 定义逻辑流程

`SimpleObject` 实现 `Object`, `SimpleObjectFactory` 实现 `ObjectFactoryTemplate`


