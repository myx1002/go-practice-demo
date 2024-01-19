package controller

import (
	"context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Category = cCategory{}

type cCategory struct {
}

func (c *cCategory) CategoryCreate(ctx context.Context, req *backend.CategoryCreateReq) (res *backend.CategoryCreateRes, err error) {
	out, err := service.Category().CategoryCreate(ctx, model.CategoryCreateInput{
		ParentId: req.ParentId,
		Name:     req.Name,
		Level:    req.Level,
		Sort:     req.Sort,
		PicUrl:   req.PicUrl,
	})
	if err != nil {
		return nil, err
	}

	return &backend.CategoryCreateRes{
		CategoryId: out.CategoryId,
	}, nil
}

func (c *cCategory) CategoryUpdate(ctx context.Context, req *backend.CategoryUpdateReq) (res *backend.CategoryUpdateRes, err error) {
	out, err := service.Category().CategoryUpdate(ctx, model.CategoryUpdateInput{
		CategoryId: req.CategoryId,
		ParentId:   req.ParentId,
		Name:       req.Name,
		Level:      req.Level,
		Sort:       req.Sort,
		PicUrl:     req.PicUrl,
	})
	if err != nil {
		return nil, err
	}

	return &backend.CategoryUpdateRes{
		CategoryId: out.CategoryId,
	}, nil
}

func (c *cCategory) CategoryDelete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().CategoryDelete(ctx, req.CategoryId)
	if err != nil {
		return nil, err
	}

	return &backend.CategoryDeleteRes{
		CategoryId: req.CategoryId,
	}, nil
}

// List
func (a *cPosition) CategoryList(ctx context.Context, req *backend.CategoryListReq) (res *backend.CategoryListRes, err error) {
	out, err := service.Category().CategoryList(ctx, model.CategoryListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.CategoryListRes{
		List:  out.List,
		Total: out.Total,
	}, nil
}
