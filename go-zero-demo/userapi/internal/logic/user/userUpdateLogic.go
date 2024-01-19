package user

import (
	"context"

	"go-zero-demo/userapi/internal/svc"
	"go-zero-demo/userapi/internal/types"
	"go-zero-demo/userapi/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	user := &model.User{}
	user.Name = req.Name
	user.UserId = req.UserId
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}

	userData := &model.UserData{
		Id:     4,
		UserId: user.UserId,
		Data:   "13213323",
	}

	err = l.svcCtx.UserDataModel.Update(l.ctx, userData)

	if err != nil {
		return nil, err
	}

	return &types.UserUpdateResp{Flag: true}, nil
}
