// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"go-zero-rpc-demo/orderrpc/internal/logic"
	"go-zero-rpc-demo/orderrpc/internal/svc"
	"go-zero-rpc-demo/orderrpc/order_pb"
)

type OrderInfoServer struct {
	svcCtx *svc.ServiceContext
	order_pb.UnimplementedOrderInfoServer
}

func NewOrderInfoServer(svcCtx *svc.ServiceContext) *OrderInfoServer {
	return &OrderInfoServer{
		svcCtx: svcCtx,
	}
}

// 定义一个 OrderInfo 一元 rpc 方法，请求体和响应体必填。
func (s *OrderInfoServer) GetOrderInfo(ctx context.Context, in *order_pb.GetOrderInfoReq) (*order_pb.GetOrderInfoResp, error) {
	l := logic.NewGetOrderInfoLogic(ctx, s.svcCtx)
	return l.GetOrderInfo(in)
}