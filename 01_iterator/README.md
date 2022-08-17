## Iterator 模式

Iterator 模式（迭代器模式）用于在数据集合中按照顺序遍历集合。

### 示例程序类图

![iterator](iterator.png)

### 拓展思路的要点

1. Iterator 将遍历和实现分开，BookShelf 实现变化并不会影响遍历。
1. 使用 Aggregate 接口和 Iterator 接口，弱化类之间的耦合。
1. Aggregate 和 Iterator 是对应的，BookShelf 实现发生变化，BookShelfIterator 也需要变化。
1. 一个具体集合可以实现多个不同的 Iterator：从前往后遍历迭代器、从后往前遍历迭代器等等 。