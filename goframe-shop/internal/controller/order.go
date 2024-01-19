package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Order cOrder

type cOrder struct {
}

func (c *cOrder) PlaceAnOrder(ctx context.Context, req *frontend.PlaceAnOrderReq) (res *frontend.PlaceAnOrderRes, err error) {
	data := model.PlaceAnOrderInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Order().PlaceOrder(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.PlaceAnOrderRes{OrderId: uint(out.OrderId)}, nil
}
