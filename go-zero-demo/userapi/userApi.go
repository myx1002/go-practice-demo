package main

import (
	"flag"
	"fmt"

	"go-zero-demo/userapi/common/middleware"
	"go-zero-demo/userapi/internal/config"
	"go-zero-demo/userapi/internal/handler"
	"go-zero-demo/userapi/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// 获取启动命令行参数，如果启动时没有-f参数，则默认读取"etc/userApi.yaml"配置
var configFile = flag.String("f", "etc/userApi.yaml", "the config file")

func main() {
	// 解析命令行参数
	flag.Parse()

	// 把解析出来的配置文件内容赋值给config对象
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 启动server服务
	// 包括日志上报、链路追踪、系统监控等等等
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 引入全局中间件
	server.Use(middleware.NewGlobalTestMiddleware().Handle)

	// 初始化上下文context，把一些配置加载到context中，例如上面的配置，还有Kafka、Redis等等可以引入
	ctx := svc.NewServiceContext(c)

	// 路由注册
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
