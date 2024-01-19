package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/net/context"

	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Praise cPraise

type cPraise struct {
}

func (a *cPraise) Create(ctx context.Context, req *frontend.PraiseAddReq) (res *frontend.PraiseAddRes, err error) {
	out, err := service.Praise().PraiseAdd(ctx, model.PraiseAddInput{
		Type:     req.Type,
		ObjectId: req.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseAddRes{PraiseId: out.PraiseId}, nil
}

func (a *cPraise) Cancel(ctx context.Context, req *frontend.PraiseCancelReq) (res *frontend.PraiseCancelRes, err error) {
	err = service.Praise().PraiseCancel(ctx, model.PraiseCancelInput{
		Type:     req.Type,
		ObjectId: req.ObjectId,
	})
	return nil, err
}

func (a *cPraise) List(ctx context.Context, req *frontend.PraiseListReq) (*frontend.PraiseListRes, error) {
	data := model.PraiseListInput{}
	err := gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	list, err := service.Praise().PraiseList(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseListRes{
		CommonPaginationRes: frontend.CommonPaginationRes{
			List:  list.List,
			Total: list.Total,
			Page:  list.Page,
			Size:  list.Size,
		}}, nil
}
