## Factory Method 模式

> **工厂方法** 是一种创建型设计模式， 解决了在不指定具体类的情况下创建产品对象的问题。

### 示例程序类图

1. Product 接口：产品（Product），工厂方法中生成的实例的抽象。
2. Factory 抽象类：创建者（Creator），负责生成 Product。
3. IDCard 类：具体的产品（ConcreteProduct）。
4. IDCardFactory 类：具体的创建者（ConcreteFactory），负责生成 IDCard。

![factory](./factory_method.png)

### 拓展思路的要点

1. 用 Template Method 模式（模板方法模式）来构建生成实例的工厂，就是 Factory Method 模式（工厂方法模式）。
2. 父类只决定实例的生成方式（Create()），具体处理由子类负责（CreateProduct()、RegisterProduct()），解耦父类与具体类。
