// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop/internal/model"
)

type (
	IGoods interface {
		GoodsCreate(ctx context.Context, in model.GoodsCreateInput) (out *model.GoodsCreateOutput, err error)
		GoodsUpdate(ctx context.Context, in model.GoodsUpdateInput) (out model.GoodsUpdateOutput, err error)
		GoodsDelete(ctx context.Context, goodsId uint) (err error)
		GoodsList(ctx context.Context, in model.GoodsListInput) (out *model.GoodsListOutput, err error)
		GoodsDetail(ctx context.Context, input model.GoodsDetailInput) (out *model.GoodsDetailOutput, err error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
