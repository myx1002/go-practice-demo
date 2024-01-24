package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/orderrpc/order_pb"
	"go-zero-demo/userapi/internal/svc"
	"go-zero-demo/userapi/internal/types"
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
	user, err := l.svcCtx.OrderRpcClient.GetOrderInfo(l.ctx, &order_pb.GetOrderInfoReq{Id: 1})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{UserId: user.Id, Name: user.Name}, nil
}
