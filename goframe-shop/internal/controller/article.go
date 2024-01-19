package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/net/context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Article cArticle

type cArticle struct {
}

func (a *cArticle) Create(ctx context.Context, req *backend.ArticleCreateReq) (res *backend.ArticleCreateRes, err error) {
	// 类型转换，如果字段太多需要一个个赋值的话很麻烦，所以直接用类型转换，把req的数据scan到model.ArticleCreateInput里面即可
	input := model.ArticleCreateInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	input.UserId = gconv.Uint(ctx.Value(consts.CtxAdminId))
	out, err := service.Article().ArticleCreate(ctx, input)
	if err != nil {
		return nil, err
	}
	return &backend.ArticleCreateRes{ArticleId: out.ArticleId}, nil
}

func (a *cArticle) Delete(ctx context.Context, req *backend.ArticleDeleteReq) (res *backend.ArticleDeleteRes, err error) {
	err = service.Article().ArticleDelete(ctx, req.ArticleId)
	return &backend.ArticleDeleteRes{ArticleId: req.ArticleId}, err
}

func (a *cArticle) Update(ctx context.Context, req *backend.ArticleUpdateReq) (res *backend.ArticleUpdateRes, err error) {
	// 类型转换，如果字段太多需要一个个赋值的话很麻烦，所以直接用类型转换，把req的数据scan到model.ArticleUpdateInput里面即可
	input := model.ArticleUpdateInput{}
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	input.UserId = gconv.Uint(ctx.Value(consts.CtxAdminId))
	out, err := service.Article().ArticleUpdate(ctx, input)
	return &backend.ArticleUpdateRes{ArticleId: out.ArticleId}, err
}

// List
func (a *cArticle) List(ctx context.Context, req *backend.ArticleListReq) (res *backend.ArticleListRes, err error) {
	getListRes, err := service.Article().ArticleList(ctx, model.ArticleListInput{
		Page:  req.Page,
		Size:  req.Size,
		Title: req.Title,
	})
	if err != nil {
		return nil, err
	}

	return &backend.ArticleListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  getListRes.List,
			Page:  getListRes.Page,
			Size:  getListRes.Size,
			Total: getListRes.Total,
		},
	}, nil
}
