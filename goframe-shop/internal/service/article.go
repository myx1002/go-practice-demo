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
	IArticle interface {
		ArticleCreate(ctx context.Context, in model.ArticleCreateInput) (out *model.ArticleCreateOutput, err error)
		ArticleUpdate(ctx context.Context, in model.ArticleUpdateInput) (out model.ArticleUpdateOutput, err error)
		ArticleDelete(ctx context.Context, goodsId uint) (err error)
		ArticleList(ctx context.Context, in model.ArticleListInput) (out *model.ArticleListOutput, err error)
	}
)

var (
	localArticle IArticle
)

func Article() IArticle {
	if localArticle == nil {
		panic("implement not found for interface IArticle, forgot register?")
	}
	return localArticle
}

func RegisterArticle(i IArticle) {
	localArticle = i
}
