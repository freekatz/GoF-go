# Factory Method 模式

Factory Method 模式是 Template Method 模式的一个特例, 即在模板方法规定实例生成的逻辑流程

父类决定要生成实例的方式, 而不决定生成的具体类型, 这些处理都交给子类进行

本实例中, `Object` 接口定义 `DoSomething()`, `ObjectFactory` 接口定义生成 `Object` 的方法, `ObjectFactoryTemplate` 实现生成 `Object` 的模板方法, 定义逻辑流程

`SimpleObject` 实现 `Object`, `SimpleObjectFactory` 实现 `ObjectFactoryTemplate`

`Object` 定义了在 Factory Method 模式中生成的实例所拥有的行为, 这些行为最终由 `SimpleObject` 这些子类来实现

`ObjectFactory` 与 `SimpleObjectFactory` 的关系和上面两者类似, `ObjectFactoryTemplate` 则定义了 Factory 模板方法

> Factory Method 模式是 Template Method 模式的一个特例
> 
> 一般 `ObjectFactoryTemplate` 只存在一个实例即可, 这是要使用 Singleton 模式
> 
> 有时可以将 Composite 模式应用到 Object 角色上面
> 
> Iterator 模式中生成 Iterator 时会用到 Factory Method 模式
