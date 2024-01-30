package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"go-zero-rpc-demo/orderrpc/internal/config"
	"go-zero-rpc-demo/orderrpc/model"
)

type ServiceContext struct {
	Config      config.Config
	BOrderModel model.BOrdersModel
	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       1,
	})

	return &ServiceContext{
		Config:      c,
		AsynqClient: client,
		BOrderModel: model.NewBOrdersModel(sqlConn, c.Cache),
	}
}
