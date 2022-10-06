## Adapter 模式

> **适配器** 是一种结构型设计模式， 它能使不兼容的对象能够相互合作。

### 示例程序类图

1. Print 接口：对象（Target），定义所需方法。
2. Banner 类：被适配（Adaptee），持有既定方法。
3. PrintBanner 类：适配器（Adapter），使用 Adaptee 的既定方法，实现 Target 定义的方法。
4. main：请求者（Client），使用 Target。

![adapter](./adapter.png)

### 拓展思路的要点

1. 适配器模式有两种：
   - 类适配器模式：使用继承的适配器，Golang 没有继承无法实现。PrintBanner 继承 Banner 的两个方法来实现 Print 接口的方法。
   - 对象适配器模式：使用委托的适配器。PrintBanner 将方法调用委派给 Banner 实例来实现 Print 接口的方法。
2. 我们调用的是 Print 接口的方法，Banner 类的两个方法被完全隐藏起来了。
3. 对于现有已充分测试，并被使用的类，可以使用 Adapter 模式进行复用或新增、修改功能。方便排查问题，避免重新进行测试。
4. 版本升级时，可以使用 Adapter 模式兼容新旧版本。