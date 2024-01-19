package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserDataModel = (*customUserDataModel)(nil)

type (
	// UserDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDataModel.
	UserDataModel interface {
		userDataModel
		TransInsert(ctx context.Context, session sqlx.Session, data *UserData) (sql.Result, error)
	}

	customUserDataModel struct {
		*defaultUserDataModel
	}
)

func (m *defaultUserDataModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserData) (sql.Result, error) {
	bubbleUserDataIdKey := fmt.Sprintf("%s%v", cacheBubbleUserDataIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userDataRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Data, data.UserId)
	}, bubbleUserDataIdKey)
	return ret, err
}

// NewUserDataModel returns a model for the database table.
func NewUserDataModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserDataModel {
	return &customUserDataModel{
		defaultUserDataModel: newUserDataModel(conn, c, opts...),
	}
}
