package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"goframe-shop/api/frontend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var User cUser

type cUser struct {
}

func (c *cUser) RegisterUser(ctx context.Context, req *frontend.RegisterUserReq) (res *frontend.RegisterUserRes, err error) {
	detail := model.RegisterUserInput{}
	err = gconv.Struct(req, &detail)
	if err != nil {
		return nil, err
	}
	user, err := service.User().RegisterUser(ctx, detail)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterUserRes{UserId: user.UserId}, nil
}

func (c *cUser) UserInfo(ctx context.Context, req *frontend.UserInfoReq) (res *frontend.UserInfoRes, err error) {
	return &frontend.UserInfoRes{
		UserId: gconv.Uint(ctx.Value(consts.CtxUserId)),
		Name:   gconv.String(ctx.Value(consts.CtxUserName)),
		Sign:   gconv.String(ctx.Value(consts.CtxUserSign)),
		Status: gconv.Int(ctx.Value(consts.CtxUserStatus)),
		Avatar: gconv.String(ctx.Value(consts.CtxUserAvatar)),
	}, nil
}

func (c *cUser) ChangeUserPassword(ctx context.Context, req *frontend.ChangeUserPasswordReq) (res *frontend.ChangeUserPasswordRes, err error) {
	data := model.ChangeUserPasswordInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().ChangeUserPassword(ctx, data)
	if err != nil {
		return nil, err
	}

	return &frontend.ChangeUserPasswordRes{UserId: out.UserId}, nil
}
