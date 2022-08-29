# Go 语言设计模式

<div align="center">
<img src="./gopher.png" width=20%>
</div>

Go 语言设计模式的实现代码，[Quick Start](./tutorial/tutorial.md)。

🚧施工中，进度[9/23]

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [图解设计模式分类](#%E5%9B%BE%E8%A7%A3%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F%E5%88%86%E7%B1%BB)
  - [适应设计模式](#%E9%80%82%E5%BA%94%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F)
  - [交给子类](#%E4%BA%A4%E7%BB%99%E5%AD%90%E7%B1%BB)
  - [生成实例](#%E7%94%9F%E6%88%90%E5%AE%9E%E4%BE%8B)
  - [分开考虑](#%E5%88%86%E5%BC%80%E8%80%83%E8%99%91)
- [GoF 设计模式分类](#gof-%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F%E5%88%86%E7%B1%BB)
  - [创建型设计模式](#%E5%88%9B%E5%BB%BA%E5%9E%8B%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F)
  - [结构性设计模式](#%E7%BB%93%E6%9E%84%E6%80%A7%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F)
  - [行为型设计模式](#%E8%A1%8C%E4%B8%BA%E5%9E%8B%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F)
- [感谢](#%E6%84%9F%E8%B0%A2)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 图解设计模式分类

### 适应设计模式

- [Iterator 模式（迭代器模式）](./01_iterator/)：一个一个遍历
- [Adapter 模式（适配器模式）](./02_adapter/)：加个 “适配器” 以便于复用

### 交给子类

- [Template Method 模式（模板方法模式）](./03_template_method/)：将具体处理交给子类
- [Factory Method 模式（工厂方法模式）](./04_factory_method/)：将实例的生成交给子类

### 生成实例

- [SingleTon 模式（单例模式）](./05_singleton/)：只有一个实例
- [Prototype 模式（原型模式）](./06_prototype/)：通过复制生成实例
- [Builder 模式（建造者模式）](./07_builder/)：组装复杂的实例
- [Abstract Factory 模式（抽象工厂模式）](./08_abstract_factory/)：将关联零件组装成产品

### 分开考虑

- [Bridge 模式（桥接模式）](./09_bridge/)：将类的功能层次结构与实现层次结构分离

## GoF 设计模式分类

### 创建型设计模式

- [Abstract Factory 模式（抽象工厂模式）](./08_abstract_factory/)
- [Builder 模式（建造者模式）](./07_builder/)
- [Factory Method 模式（工厂方法模式）](./04_factory_method/)
- [Prototype 模式（原型模式）](./06_prototype/)
- [SingleTon 模式（单例模式）](./05_singleton/)

### 结构性设计模式

- [Adapter 模式（适配器模式）](./02_adapter/)
- [Bridge 模式（桥接模式）](./09_bridge/)

### 行为型设计模式

- [Iterator 模式（迭代器模式）](./01_iterator/)
- [Template Method 模式（模板方法模式）](./03_template_method/)

## 感谢

- [图解设计模式](https://book.douban.com/subject/26933281/)
- [Refactoring.Guru 设计模式](https://refactoringguru.cn/design-patterns)
- [senghoo/golang-design-pattern](https://github.com/senghoo/golang-design-pattern) 