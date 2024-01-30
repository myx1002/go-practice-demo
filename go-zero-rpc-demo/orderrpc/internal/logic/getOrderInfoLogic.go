package logic

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"go-zero-rpc-demo/orderrpc/internal/svc"
	"go-zero-rpc-demo/orderrpc/model"
	"go-zero-rpc-demo/orderrpc/order_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type DeferQueueDemoPayload struct {
	TaskTime string
	TaskId   int64
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
	// 获取客户端拦截器传过来的meta值
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("md_test_name_key")
		fmt.Println("获取客户端传过来的metadata:", md)
		if len(tmp) > 0 {
			fmt.Println("获取客户端传过来的metadata", tmp[0])
		}
	}

	order, err := l.svcCtx.BOrderModel.FindOne(l.ctx, in.Id)
	if err != nil || err == model.ErrNotFound {
		return nil, err
	}
	return &order_pb.GetOrderInfoResp{Id: order.OrderId, Name: order.OrderNo}, nil
}
