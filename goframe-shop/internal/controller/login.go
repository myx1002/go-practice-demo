package controller

import (
	"context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/service"
)

var Login = cLogin{}

type cLogin struct{}

// for session
//func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//	res = &backend.LoginDoRes{}
//
//	err = service.Login().Login(ctx, model.UserLoginInput{
//		Name:     req.Name,
//		Password: req.Password,
//	})
//	if err != nil {
//		return
//	}
//
//	res.Info = service.Session().GetUser(ctx)
//	return
//}

// 登录-这个是jwt的token登录方法，因为我们用到了gtoken，所以这个注释掉
//func (c *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
//	res = &backend.LoginDoRes{}
//	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
//	return
//}
//
//func (c *cLogin) RefreshToken(ctx context.Context, req *backend.RefreshTokenReq) (res *backend.RefreshTokenRes, err error) {
//	res = &backend.RefreshTokenRes{}
//	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
//	return
//}

func (c *cLogin) Logout(ctx context.Context, req *backend.LogoutReq) (res *backend.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
