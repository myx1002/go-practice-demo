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
	ICategory interface {
		CategoryCreate(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error)
		CategoryUpdate(ctx context.Context, in model.CategoryUpdateInput) (out model.CategoryUpdateOutput, err error)
		CategoryDelete(ctx context.Context, CategoryId uint) error
		CategoryList(ctx context.Context, in model.CategoryListInput) (out *model.CategoryListOutput, err error)
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
