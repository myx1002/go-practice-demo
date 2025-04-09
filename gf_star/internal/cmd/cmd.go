package cmd

import (
	"context"
	"gf_star/internal/controller/account"
	"gf_star/internal/controller/user"
	"gf_star/internal/controller/word"
	"gf_star/internal/logic/middleware"

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
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1(),
					)
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.Bind(
							account.NewV1(),
							word.NewV1(),
						)
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
