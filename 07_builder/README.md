## Builder 模式

Builder 模式（建造者模式）用于组装具有复杂结构的实例。

### 示例程序类图

![builder](./builder_class.png)

### 示例程序时序图

![builder](./builder_sequence.png)

### 拓展思路的要点

1. Director 不关注具体子类，仅使用 Builder 的方法来组装实例。子类具有可替换性。
1. Builder 类应尽可能灵活，来适应未来子类的需求或变化。
