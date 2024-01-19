package session

import (
	"context"

	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sSession struct{}

const (
	sessionKeyUser = "SessionKeyUser" // 用户信息存放在Session中的Key
)

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return &sSession{}
}

// 设置用户Session.
func (s *sSession) SetUser(ctx context.Context, admin *entity.AdminInfo) error {
	return service.BizCtx().Get(ctx).Session.Set(sessionKeyUser, admin)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sSession) GetUser(ctx context.Context) *entity.AdminInfo {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyUser)
		if !v.IsNil() {
			var user *entity.AdminInfo
			_ = v.Struct(&user)
			return user
		}
	}
	return &entity.AdminInfo{}
}

// 删除用户Session。
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}
