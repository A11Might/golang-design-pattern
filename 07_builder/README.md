## Builder 模式

> **建造者模式** 是一种创建型设计模式， 使你能够分步骤创建复杂对象。

### 示例程序类图

1. Builder 接口：建造者（Builder），定义用于生成实例的方法。
2. TextBuilder 类、HTMLBuilder类：具体的建造者（ConcreteBuilder），实现 Builder 定义的方法。
3. Director 类：监工（Director），使用 Builder 定义的方法生成实例。
4. main：使用者（Client），使用建造者模式。

![builder](./builder_class.png)

### 示例程序时序图

![builder](./builder_sequence.png)

### 拓展思路的要点

1. Director 不关注具体子类，仅使用 Builder 的方法来组装实例。子类具有可替换性。
1. Builder 类应尽可能灵活，来适应未来子类的需求或变化。
