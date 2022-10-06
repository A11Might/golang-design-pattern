## Template Method 模式

> **模版方法** 是一种行为设计模式， 它在基类中定义了一个算法的框架， 允许子类在不修改结构的情况下重写算法的特定步骤。

### 示例程序类图

1. Display 抽象类：抽象类（AbstractClass），定义模板方法及模板方法使用的抽象方法。
2. CharDisplay 类、StringDisplay 类：具体类（ConcreteClass），实现 AbstractClass 定义的抽象方法。

![template_method](./template_method.png)

### 拓展思路的要点

1. 在父类模板方法中编写的算法（Display()），无需在子类中重复编写。
2. 需要理解父类中的抽象方法（Display()），才能够编写出子类（CharDisplay, StringDisplay）。
3. 父类变量保存子类实例，然后调用 Display()，无论哪个子类实例都能正常运行，符合 LSP（里氏替换原则）
