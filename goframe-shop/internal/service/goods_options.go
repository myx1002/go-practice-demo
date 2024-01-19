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
	IGoodsOptions interface {
		GoodsOptionsCreate(ctx context.Context, in model.GoodsOptionsCreateInput) (out *model.GoodsOptionsCreateOutput, err error)
		GoodsOptionsUpdate(ctx context.Context, in model.GoodsOptionsUpdateInput) (out model.GoodsOptionsUpdateOutput, err error)
		GoodsOptionsDelete(ctx context.Context, goodsId uint) (err error)
		GoodsOptionsList(ctx context.Context, in model.GoodsOptionsListInput) (out *model.GoodsOptionsListOutput, err error)
	}
)

var (
	localGoodsOptions IGoodsOptions
)

func GoodsOptions() IGoodsOptions {
	if localGoodsOptions == nil {
		panic("implement not found for interface IGoodsOptions, forgot register?")
	}
	return localGoodsOptions
}

func RegisterGoodsOptions(i IGoodsOptions) {
	localGoodsOptions = i
}
