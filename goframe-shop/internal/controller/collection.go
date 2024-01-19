package controller

import (
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"

	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Collection cCollection

type cCollection struct {
}

func (a *cCollection) Create(ctx context.Context, req *frontend.CollectionAddReq) (res *frontend.CollectionAddRes, err error) {
	out, err := service.Collection().CollectionAdd(ctx, model.CollectionAddInput{
		Type:     req.Type,
		ObjectId: req.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionAddRes{CollectionId: out.CollectionId}, nil
}

func (a *cCollection) Cancel(ctx context.Context, req *frontend.CollectionCancelReq) (res *frontend.CollectionCancelRes, err error) {
	_, err = service.Collection().CollectionCancel(ctx, model.CollectionCancelInput{
		Type:     req.Type,
		ObjectId: req.ObjectId,
	})
	return nil, err
}

func (a *cCollection) List(ctx context.Context, req *frontend.CollectionListReq) (res *frontend.CollectionListRes, err error) {
	g.Dump(req)
	out, err := service.Collection().CollectionList(ctx, model.CollectionListInput{
		Size: req.Size,
		Page: req.Page,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CollectionListRes{
		CommonPaginationRes: frontend.CommonPaginationRes{
			List:  out.List,
			Size:  out.Size,
			Page:  out.Page,
			Total: out.Total,
		},
	}, nil
}
