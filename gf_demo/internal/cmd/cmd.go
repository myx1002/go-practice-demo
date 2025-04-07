package cmd

import (
	"context"
	"gf_demo/internal/controller/hello"
	"gf_demo/internal/controller/user"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetPort(8881)
			// 开启静态资源访问
			s.SetServerRoot("resource/public")

			// 分组路由
			//s.Group("/hello", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Bind(
			//		hello.NewHello(),
			//	)
			//})
			//s.Group("/user", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Bind(
			//		user.New(),
			//	)
			//})

			// 嵌套分组
			//s.Group("/api", func(group *ghttp.RouterGroup) {
			//	group.Group("/hello", func(group *ghttp.RouterGroup) {
			//		group.Middleware(ghttp.MiddlewareHandlerResponse)
			//		group.Bind(
			//			hello.NewHello(),
			//		)
			//	})
			//	group.Group("/user", func(group *ghttp.RouterGroup) {
			//		group.Middleware(ghttp.MiddlewareHandlerResponse)
			//		group.GET("/get", user.New().Get)
			//		group.DELETE("/delete", user.New().Delete)
			//		group.POST("/add", user.New().Add)
			//		group.PUT("/update", user.New().Update)
			//	})
			//})

			// 规范路由
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewHello(),
					user.New(),
				)
			})

			s.Run()
			return nil
		},
	}
)
