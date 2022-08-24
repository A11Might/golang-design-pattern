Golang 通过结构体（Foo）匿名组合（IFoo）接口来实现抽象类，由于接口中不能包含字段，所以只能通过 Get()、Set() 方法（GetFooField()、SetFooField()）来操作抽象类中的字段。

抽象类（Foo）声明的抽象方法（AbstracMethod()），具体由子类（xxFoo）实现。

具体使用参考[示例代码](./main.go)。

![tutorial](./tutorial.png)
