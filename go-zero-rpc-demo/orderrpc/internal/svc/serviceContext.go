package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"go-zero-rpc-demo/orderrpc/internal/config"
	"go-zero-rpc-demo/orderrpc/model"
)

type ServiceContext struct {
	Config      config.Config
	BOrderModel model.BOrdersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:      c,
		BOrderModel: model.NewBOrdersModel(sqlConn, c.Cache),
	}
}
