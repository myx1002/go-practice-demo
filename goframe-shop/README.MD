# GoFrame Template For SingleRepo

## Quick Start:

- https://goframe.org/pages/viewpage.action?pageId=1114399

## GoFrame V2开发流程:

- 1.设计表结构,初始化项目，修改配置文件
- 2.使用gf gen dao生成对应的dao/do/model
- 3.编写api层:义业务侧数据结构,提供对外接口的输入/输出数据结构
- 4.编写model层:定义数据侧数据结构,提供对内的数据处理的输入输出数据结构
- 5.编写logic层，自动生成service层代码。 (通过 配置goland File Watcher自动生成,也可以通过 gf gen service手动执行脚本生成,建议前者)
- 6.在service层代码生成RegisterXX()方法后，在对应的logic模块注册服务(每个模块只要写一次)
- 7.编写controller层，接收/解析用户输入的参数，调用service层的服务。
- 8.注册路由，对外暴露接口，比如这个项目是编写cmd.go文件。
- 9.在main.go中加入一行_ "project-name/internallogic" (只需写-次)
- 10.在main.go中加入一行_ "github.com/gogf/gficontrib/drivers/mysqlv2" (如果你使用的是
  mysq|;只写- --次)

## 工程目录结构

```text
/
├── api （对外接口，对外提供服务的输入、输入数据结构定义）
├── hack（工具脚本，存放项目的开发工具、脚本等内容，如CLI工具的配置，各种shell、bat脚本文件）
├── internal （内部逻辑，业务逻辑存放的目录）
│   ├── cmd （入口指令，命令行管理目录，可以管理维护多个命令行）
│   ├── consts（常量定义，项目中所有的常量定义）
│   ├── controller（接口处理，接收、解析用户输入参数的入口/接口层）
│   ├── dao（数据访问，数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的CURD方法）
│   ├── logic（业务封装，业务逻辑封装管理，特定的业务逻辑实现和封装，往往是项目中最复杂的部分）
│   ├── model（结构模型，数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义）
│   |   ├── do（领域对象，用于dao数据操作中业务模型与示例模型转化，由工具维护，用户不能修改）
│   │   └── entity（数据模型，数据模型是模型与数据集合的一对一关系，由工具维护，用户不能修改）
│   └── service（业务接口，用于业务模块解耦的接口定义层，具体的接口实现再logic中进行注入）
├── manifest（交付清淡，包含程序编译、部署、运行、配置的文件）
├── resource（静态资源）
├── utility（工具类）
├── go.mod（依赖管理）
└── main.go （程序入口文件）
```

## 请求分层流转

### cmd

cmd层负责引导程序启动，显著的工作是初始化逻辑、注册路由对象、启动server监听、阻塞运行程序直至server退出。

### api

上层server服务接收客户端请求，转换为api中定义的Req接收对象、执行请求参数到Req对象属性的类型转换、执行Req对象中绑定的基础校验并转交Req请求对象给controller层。

### controller

controller层负责接收Req请求对象后做一些业务逻辑校验，随后调用一个或多个service实现业务逻辑，将执行结构封装为约定的Res数据结构对象返回。

### model

model层中管理了所有的业务模型，service资源的Input/Output输入输出数据结构都由model层来维护。

### service

service是接口层，用于解耦业务模块，service没有具体的业务逻辑实现，具体的业务实现是依靠logic层注入的。

### logic

logic层的业务逻辑需要通过调用dao来实现数据的操作，调用dao时需要传递do数据结构对象，用于传递查询条件、输入数据。dao执行完毕后通过Entity数据模型将数据结果返回给service层。

### dao

dao层通过框架的ORM抽象层组件与底层真实的数据库交互。
![](https://moyuxing.oss-cn-guangzhou.aliyuncs.com/gridea/20231117173328.png)
