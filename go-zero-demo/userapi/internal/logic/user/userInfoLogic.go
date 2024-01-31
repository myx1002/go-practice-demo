package user

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/orderrpc/order_pb"
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

	// 获取jwt解析出来的user_phone
	fmt.Println("get UserInfo....", l.ctx.Value("user_phone"))
	user, err := l.svcCtx.OrderRpcClient.GetOrderInfo(l.ctx, &order_pb.GetOrderInfoReq{Id: req.UserId})

	if err != nil || err == model.ErrNotFound {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.UserInfoResp{UserId: user.Id, Name: user.Name}, nil
}
