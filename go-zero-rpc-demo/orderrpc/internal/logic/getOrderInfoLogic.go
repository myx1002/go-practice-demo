package logic

import (
	"context"

	"go-zero-rpc-demo/orderrpc/internal/svc"
	"go-zero-rpc-demo/orderrpc/order_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoLogic {
	return &GetOrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义一个 OrderInfo 一元 rpc 方法，请求体和响应体必填。
func (l *GetOrderInfoLogic) GetOrderInfo(in *order_pb.GetOrderInfoReq) (*order_pb.GetOrderInfoResp, error) {
	// todo: add your logic here and delete this line

	return &order_pb.GetOrderInfoResp{Id: 11, Name: "apple2222"}, nil
}
