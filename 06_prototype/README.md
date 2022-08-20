## Prototype 模式

根据实例来生成新实例的模式就是 Prototype 模式（原型模式）。

### 示例程序类图

![prototype](./prototype.png)

### 拓展思路的要点

1. Manager 调用 Product 的 CreateClone() 复制实例，不关注具体复制的是哪一个类。解耦 Manager 与具体类。
2. 当对象种类繁多无法整合、难以根据类生成实例或想解耦框架与生成实例时，可以使用原型模式生成实例。
