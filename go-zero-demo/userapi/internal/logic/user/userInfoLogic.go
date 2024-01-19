package user

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/userapi/internal/svc"
	"go-zero-demo/userapi/internal/types"
	"go-zero-demo/userapi/model"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询数据失败！")
	}
	if user == nil {
		return nil, errors.New("该用户不存在！")
	}

	return &types.UserInfoResp{UserId: req.UserId, Name: user.Name}, nil
}
