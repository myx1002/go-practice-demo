package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/net/context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var GoodsOptions cGoodsOptions

type cGoodsOptions struct {
}

func (a *cGoodsOptions) Create(ctx context.Context, req *backend.GoodsOptionsCreateReq) (res *backend.GoodsOptionsCreateRes, err error) {
	// 类型转换，如果字段太多需要一个个赋值的话很麻烦，所以直接用类型转换，把req的数据scan到model.GoodsOptionsCreateInput里面即可
	input := model.GoodsOptionsCreateInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}

	out, err := service.GoodsOptions().GoodsOptionsCreate(ctx, input)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsOptionsCreateRes{GoodsOptionsId: out.GoodsOptionsId}, nil
}

func (a *cGoodsOptions) Delete(ctx context.Context, req *backend.GoodsOptionsDeleteReq) (res *backend.GoodsOptionsDeleteRes, err error) {
	err = service.GoodsOptions().GoodsOptionsDelete(ctx, req.GoodsOptionsId)
	return &backend.GoodsOptionsDeleteRes{GoodsOptionsId: req.GoodsOptionsId}, err
}

func (a *cGoodsOptions) Update(ctx context.Context, req *backend.GoodsOptionsUpdateReq) (res *backend.GoodsOptionsUpdateRes, err error) {
	// 类型转换，如果字段太多需要一个个赋值的话很麻烦，所以直接用类型转换，把req的数据scan到model.GoodsOptionsUpdateInput里面即可
	input := model.GoodsOptionsUpdateInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}

	out, err := service.GoodsOptions().GoodsOptionsUpdate(ctx, input)
	return &backend.GoodsOptionsUpdateRes{GoodsOptionsId: out.GoodsOptionsId}, err
}

// List
func (a *cGoodsOptions) List(ctx context.Context, req *backend.GoodsOptionsListReq) (res *backend.GoodsOptionsListRes, err error) {
	getListRes, err := service.GoodsOptions().GoodsOptionsList(ctx, model.GoodsOptionsListInput{
		Page: req.Page,
		Size: req.Size,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &backend.GoodsOptionsListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  getListRes.List,
			Page:  getListRes.Page,
			Size:  getListRes.Size,
			Total: getListRes.Total,
		},
	}, nil
}
