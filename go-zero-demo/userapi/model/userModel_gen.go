// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheBubbleUserUserIdPrefix = "cache:bubble:user:userId:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)

		FindOne(ctx context.Context, userId int64) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		UserId int64  `db:"user_id"` // 用户id
		Name   string `db:"name"`    // 用户名
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, userId int64) error {
	bubbleUserUserIdKey := fmt.Sprintf("%s%v", cacheBubbleUserUserIdPrefix, userId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId)
	}, bubbleUserUserIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, userId int64) (*User, error) {
	bubbleUserUserIdKey := fmt.Sprintf("%s%v", cacheBubbleUserUserIdPrefix, userId)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, bubbleUserUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	bubbleUserUserIdKey := fmt.Sprintf("%s%v", cacheBubbleUserUserIdPrefix, data.UserId)
	// from tpl
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name)
	}, bubbleUserUserIdKey)
	return ret, err
}

func (m *defaultUserModel) TransInsert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	bubbleUserUserIdKey := fmt.Sprintf("%s%v", cacheBubbleUserUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, userRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Name)
	}, bubbleUserUserIdKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	bubbleUserUserIdKey := fmt.Sprintf("%s%v", cacheBubbleUserUserIdPrefix, data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.UserId)
	}, bubbleUserUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheBubbleUserUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
