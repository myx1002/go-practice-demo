package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Comment cComment

type cComment struct {
}

func (c *cComment) CommentAdd(ctx context.Context, req *frontend.CommentAddReq) (res *frontend.CommentAddRes, err error) {
	data := model.CommentAddInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Comment().CommentAdd(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CommentAddRes{CommentId: out.CommentId}, nil
}

func (c *cComment) CommentDelete(ctx context.Context, req *frontend.CommentDeleteReq) (res *frontend.CommentDeleteRes, err error) {
	data := model.CommentDeleteInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.Comment().CommentDelete(ctx, data)
	return &frontend.CommentDeleteRes{}, err
}

func (c *cComment) CommentList(ctx context.Context, req *frontend.CommentListReq) (*frontend.CommentListRes, error) {
	data := model.CommentListInput{}
	err := gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	list, err := service.Comment().CommentList(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.CommentListRes{
		CommonPaginationRes: frontend.CommonPaginationRes{
			List:  list.List,
			Total: list.Total,
			Page:  list.Page,
			Size:  list.Size,
		}}, nil
}
