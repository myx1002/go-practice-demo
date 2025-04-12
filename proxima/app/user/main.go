package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"proxima/app/user/internal/cmd"
	_ "proxima/app/user/internal/packed"
)

func main() {
	// 读取etcd配置
	var ctx = gctx.New()
	conf, err := g.Cfg("etcd").Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}

	// 注册etcd
	var address = conf.String()
	grpcx.Resolver.Register(etcd.New(address))

	cmd.Main.Run(gctx.GetInitCtx())
}
