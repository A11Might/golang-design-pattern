## Singleton 模式

> **单例** 是一种创建型设计模式， 让你能够保证一个类只有一个实例， 并提供一个访问该实例的全局节点。

### 示例程序类图

1. singleton 类：单例（Singleton），声明返回唯一实例的方法。

![singleton](./singleton.png)

### 拓展思路的要点

1. 第一次调用 GetInstance() 时，生成唯一的一个实例。
