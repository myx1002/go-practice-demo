package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"go-zero-demo/orderrpc/orderinfo"
	"go-zero-demo/userapi/internal/config"
	"go-zero-demo/userapi/internal/middleware"
	"go-zero-demo/userapi/model"
)

// ServiceContext 初始化配置
type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
	UserModel      model.UserModel
	UserDataModel  model.UserDataModel
	OrderRpcClient orderinfo.OrderInfo // rpc客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserModel:      model.NewUserModel(sqlConn, c.Cache),
		UserDataModel:  model.NewUserDataModel(sqlConn, c.Cache),
		OrderRpcClient: orderinfo.NewOrderInfo(zrpc.MustNewClient(c.OrderRpcConf)), // 引入rpc客户端到service上下文
	}
}
