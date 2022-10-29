# 设计模式(Go语言版本)
## 定义
> 设计模式是在特定环境下人们解决某类重复出现问题的一套成功或有效的解决方案。

## 设计原则(7个):SOLID+合成复用原则+迪米特法则
代码质量好坏评判标准:
* 可维护性(maintainability)
* 可读性(readability)
* 可扩展性(extensibility)
* 灵活性(flexibility)
* 简洁性(simplicity)
* 可复用性(reusability)
* 可测试性(testability)

如何同时提高一个软件系统的可维护性和可复用性是面向对象设计需要解决的核心问题之一。
面向对象设计原则为支持可维护性复用而诞生，这些原则蕴含在很多设计模式中，它们是从许多设计方案中总结出的指导性原则。

| 设计原则 | 定义 |
|--| --- |
| 单一职责原则(Single Responsibility Principle, SRP) | 一个类只负责一个功能领域中的相应职责 |
| 开闭原则(Open-Closed Principle, OCP) | 软件实体应对扩展开放，而对修改关闭 |
| 里氏代换原则(Liskov Substitution Principle, LSP) | 所有引用基类对象的地方能够透明地使用其子类的对象 |
| 依赖倒转原则(Dependence  Inversion Principle, DIP) |抽象不应该依赖于细节，细节应该依赖于抽象 |
|接口隔离原则(Interface Segregation Principle, ISP) | 使用多个专门的接口，而不使用单一的总接口|
|合成复用原则(Composite Reuse Principle, CRP) |尽量使用对象组合，而不是继承来达到复用的目的 |
|迪米特法则(Law of Demeter, LoD) | 一个软件实体应当尽可能少地与其他实体发生相互作用 |

## 种类与个数
GoF的23种 + “简单工厂模式” = 24种。
### 创建型模式(6个): 如何创建对象
* [单例模式](https://github.com/SnDragon/go-design-pattern/blob/master/creator/singleton)

### 结构型模式(7个): 如何实现类或对象的组合


### 行为型模式(11个): 类或对象怎样交互以及怎样分配职责



## 学习资料
### 博客&专栏
* [史上最全设计模式导学目录](https://blog.csdn.net/LoveLion/article/details/17517213)
* [极客时间设计模式之美](https://time.geekbang.org/column/intro/100039001)

### 视频
* [Go设计模式入门到实践](https://www.imooc.com/learn/1226)
* [模式宗师养成宝典之Java版](https://www.imooc.com/course/programdetail/pid/18)
* [Easy搞定Golang设计模式(刘丹冰)](https://space.bilibili.com/373073810/channel/collectiondetail?sid=734579)

### 书籍
* [HeadFirst设计模式](https://awesome-programming-books.github.io/design-pattern/HeadFirst%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F.pdf)
* [Gof设计模式](https://book.douban.com/subject/1052241/)
### 网站
* https://refactoringguru.cn/design-patterns/