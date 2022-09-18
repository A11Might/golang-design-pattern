## Tutorial

Golang 通过结构体（Foo）匿名组合（IFoo）接口来实现抽象类，由于接口中不能包含字段，所以只能通过 Get()、Set() 方法（GetFooField()、SetFooField()）来操作抽象类中的字段。

抽象类（Foo）声明的抽象方法（AbstractMethod()），具体由子类（xxFoo）实现。

具体使用参考[示例代码](./main.go)，下面是代码类图。

#### 类图

类图用于表示类、接口、实例等之间的静态关系。

![class](./class.png)

#### 注意

1. 初始化：当我们调用 XxFoo.Method() 这种 XxFoo 没有直接定义，而是嵌入的方法时，实际调用的是 Foo.Method()，并且 Foo.Method() 中调用的 AbstractMethod() 也是 Foo 的。所以我们需要通过初始化赋值，将 Foo.AbstractMethod() 调用路由回 XxFoo.AbstractMethod()。

   ```go
   func NewXxFoo() *XxFoo {
   	xxFoo := &XxFoo{Foo: &Foo{}}
   	xxFoo.Foo.IFoo = xxFoo // there
   	return xxFoo
   }
   ```

2. 本 repo 中的 UML 都是使用 [PlantUML](https://plantuml.com/) 编写的，示例 [类图](./class.puml)。

3. 聚合和组合的区别：聚合体现 has a 关系，组合体现 contains a 关系。

#### 感谢

- [Abstract Class reinvented](https://adrianwit.medium.com/abstract-class-reinvented-with-go-4a7326525034)