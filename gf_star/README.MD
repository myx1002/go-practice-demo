# GoFrame Template For SingleRepo

# Quick Start: 
- 文档地址：https://goframe.org/course/starbook
- 源码地址：https://github.com/oldme-git/star

# 本书简介
Web开发领域知识庞杂，各类库、标准、框架等知识盘根错节，像一棵百年老树的树根。新手朋友面对它们如同面对一个刺猬，无从下手。
本书将从GoFrame框架出发，带领新手朋友们实际开发一个项目：星辰英语本。旨在为读者提供一份深入浅出的GoFrame框架学习指南，帮助读者快速掌握GoFrame框架。
以实际项目为导向，分享更多专业、实用的编程技巧和经验，让您学有所成是作者的期望！

# dao文件生成
```bash
> gf gen dao
internal/model/do/users.go
internal/model/entity/users.go
internal/dao/internal/users.go
internal/dao/users.go
```
> model层用作GoFrame管理数据结构，与数据表一一对应，不允许用户修改。model/do/users.go用作数据写入对象，采用泛型设计，方便数据入库；model/entity/users.go用作数据读取对象，类型与数据表保持一致。

> dao层管理数据访问对象，GoFrame ORM通过它实现对数据表的增删改查。dao/internal/users.go保存内部对象实现，不允许用户修改，也不对外暴露。dao/users.go实例化数据访问对象并对外暴露，用户可以在这里添加自定义的方法。

# controller文件生成
```bash
> gf gen ctrl
api\users\users.go
internal\controller\users\users.go
internal\controller\users\users_new.go
internal\controller\users\users_v1_register.go
```
> 这里生成的四个文件，我们只需要关注users_v1_register.go即可，它用作接收HTTP请求，并调用Logic完成业务流程。