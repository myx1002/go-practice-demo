package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BOrdersModel = (*customBOrdersModel)(nil)

type (
	// BOrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBOrdersModel.
	BOrdersModel interface {
		bOrdersModel
	}

	customBOrdersModel struct {
		*defaultBOrdersModel
	}
)

// NewBOrdersModel returns a model for the database table.
func NewBOrdersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) BOrdersModel {
	return &customBOrdersModel{
		defaultBOrdersModel: newBOrdersModel(conn, c, opts...),
	}
}
