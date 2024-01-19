package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goframe-shop/internal/controller"
	"goframe-shop/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "GoFrame电商实战",
		Usage: "moyuxing",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 管理后台-启动gtoken
			// 官方文档：https://goframe.org/pages/viewpage.action?pageId=1115974
			gfAdminToken, _ := getBackendAdminGToken()

			// 管理后台-路由组
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse) // 系统返回
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().ResponseHandler, // 自定义返回
				)

				// 绑定路由-不需要登录
				group.Bind(
					//hello.NewV1(),
					controller.Admin.Create, // 管理员
				)

				// 路由组绑定-需要登录
				group.Group("/", func(group *ghttp.RouterGroup) {
					// gtoken中间件绑定
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.Rotation,     // 轮播图
						controller.Position,     // 手工位
						controller.Admin.List,   // 管理员
						controller.Admin.Update, // 管理员
						controller.Admin.Delete, // 管理员
						controller.Admin.Info,   // 管理员
						controller.Data,         // 数据统计卡片
						controller.Role,         // 角色
						controller.Permission,   //权限
						controller.File,         // 文件
						controller.Category,     // 分类
						controller.Coupon,       // 优惠卷
						controller.UserCoupon,   // 用户优惠卷
						controller.Goods,        // 商品
						controller.GoodsOptions, // 商品规格
						controller.Article,      // 文章
					)
				})
			})

			// 前台-路由组
			gfToken, _ := getFrontendGToken()
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().ResponseHandler, // 自定义返回
				)

				// 绑定路由-不需要登录
				group.Bind(
					controller.User.RegisterUser, // 注册
					controller.Goods,             // 商品
				)

				// 路由组绑定-需要登录
				group.Group("/", func(group *ghttp.RouterGroup) {
					// gtoken中间件绑定
					err := gfToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.User.UserInfo,           // 用户信息
						controller.User.ChangeUserPassword, // 修改密码
						controller.Collection,              // 收藏
						controller.Praise,                  // 点赞
						controller.Comment,                 // 评论
						controller.Order,                   // 订单
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
