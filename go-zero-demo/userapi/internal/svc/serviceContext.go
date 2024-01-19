package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

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
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserModel:      model.NewUserModel(sqlConn, c.Cache),
		UserDataModel:  model.NewUserDataModel(sqlConn, c.Cache),
	}
}
