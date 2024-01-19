package main

import (
	"flag"
	"fmt"

	"go-zero-demo/userapi/internal/config"
	"go-zero-demo/userapi/internal/handler"
	"go-zero-demo/userapi/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// 解析配置
var configFile = flag.String("f", "etc/userApi.yaml", "the config file")

func main() {
	flag.Parse()

	// 把解析出来的配置文件内容复制给config对象
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 启动server服务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 初始化上下文context，把一些配置加载到context中，例如上面的配置，还有Kafka、Redis等等可以引入
	ctx := svc.NewServiceContext(c)

	// 路由注册
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
