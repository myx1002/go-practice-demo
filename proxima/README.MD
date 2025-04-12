# GoFrame Template For MonoRepo

Quick Start: 
- https://goframe.org/course/proxima-book


> grpc安装：go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
> 
> protobuf安装：https://github.com/protocolbuffers/protobuf/releases
> 
> etcd安装
> 
> grpcx组件，让GoFrame支持微服务开发：go get -u github.com/gogf/gf/contrib/rpc/grpcx/v2
> 
> mysql插件：go get -u github.com/gogf/gf/contrib/drivers/mysql/v2
> 
> etcd插件：go get -u github.com/gogf/gf/contrib/registry/etcd/v2

## 1.项目初始化
> 执行命令，初始化一个名为proxima的Monorepo仓库：
> 
> gf init proxima -m
> 
> 执行以下命令，建立名为word的服务，同样保存在app目录下。
> 
> gf init app/word -a

## 2.代码初始化
```yaml
gfcli:
  gen:  
    dao:  
      - link: "mysql:root:12345678@tcp(srv.com:3306)/word"  
        descriptionTag: true  
  
    pbentity:  
      - link: "mysql:root:12345678@tcp(srv.com:3306)/word"
```
> 执行以下命令，生成word服务代码：
> 
> 生成dao模型: gf gen dao  
> 生成pbentity模型:gf gen pbentity
> 编写协议文件，生成控制器：gf gen pb