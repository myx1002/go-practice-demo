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
	ICollection interface {
		CollectionAdd(ctx context.Context, in model.CollectionAddInput) (out model.CollectionAddOutput, err error)
		CollectionCancel(ctx context.Context, in model.CollectionCancelInput) (out model.CollectionCancelInput, err error)
		CollectionList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error)
	}
)

var (
	localCollection ICollection
)

func Collection() ICollection {
	if localCollection == nil {
		panic("implement not found for interface ICollection, forgot register?")
	}
	return localCollection
}

func RegisterCollection(i ICollection) {
	localCollection = i
}
