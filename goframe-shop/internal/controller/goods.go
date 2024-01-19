package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/net/context"

	"goframe-shop/api/backend"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Goods cGoods

type cGoods struct {
}

func (a *cGoods) Create(ctx context.Context, req *backend.GoodsCreateReq) (res *backend.GoodsCreateRes, err error) {
	// 类型转换，如果字段太多需要一个个赋值的话很麻烦，所以直接用类型转换，把req的数据scan到model.GoodsCreateInput里面即可
	input := model.GoodsCreateInput{}
	err = gconv.Scan(req, &input)
	if err != nil {
		return nil, err
	}

	out, err := service.Goods().GoodsCreate(ctx, input)
	if err != nil {
		return nil, err
	}
	return &backend.GoodsCreateRes{GoodsId: out.GoodsId}, nil
}

func (a *cGoods) Delete(ctx context.Context, req *backend.GoodsDeleteReq) (res *backend.GoodsDeleteRes, err error) {
	err = service.Goods().GoodsDelete(ctx, req.GoodsId)
	return &backend.GoodsDeleteRes{GoodsId: req.GoodsId}, err
}

func (a *cGoods) Update(ctx context.Context, req *backend.GoodsUpdateReq) (res *backend.GoodsUpdateRes, err error) {
	// 类型转换，如果字段太多需要一个个赋值的话很麻烦，所以直接用类型转换，把req的数据scan到model.GoodsCreateInput里面即可
	input := model.GoodsUpdateInput{}
	err = gconv.Scan(req, &input)
	if err != nil {
		return nil, err
	}

	out, err := service.Goods().GoodsUpdate(ctx, input)
	return &backend.GoodsUpdateRes{GoodsId: out.GoodsId}, err
}

// List
func (a *cGoods) List(ctx context.Context, req *backend.GoodsListReq) (res *backend.GoodsListRes, err error) {
	getListRes, err := service.Goods().GoodsList(ctx, model.GoodsListInput{
		Page: req.Page,
		Size: req.Size,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &backend.GoodsListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  getListRes.List,
			Page:  getListRes.Page,
			Size:  getListRes.Size,
			Total: getListRes.Total,
		},
	}, nil
}

func (a *cGoods) Detail(ctx context.Context, req *frontend.GoodsDetailReq) (res *frontend.GoodsDetailRes, err error) {
	detail, err := service.Goods().GoodsDetail(ctx, model.GoodsDetailInput{GoodsId: req.GoodsId})
	if err != nil {
		return nil, err
	}

	data := &frontend.GoodsDetailRes{}
	err = gconv.Struct(detail, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
