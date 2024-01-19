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
	IPraise interface {
		PraiseAdd(ctx context.Context, in model.PraiseAddInput) (out model.PraiseAddOutput, err error)
		PraiseCancel(ctx context.Context, in model.PraiseCancelInput) (err error)
		PraiseList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error)
	}
)

var (
	localPraise IPraise
)

func Praise() IPraise {
	if localPraise == nil {
		panic("implement not found for interface IPraise, forgot register?")
	}
	return localPraise
}

func RegisterPraise(i IPraise) {
	localPraise = i
}
