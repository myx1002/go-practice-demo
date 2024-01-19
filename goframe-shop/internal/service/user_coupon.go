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
	IUserCoupon interface {
		UserCouponCreate(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error)
		UserCouponUpdate(ctx context.Context, in model.UserCouponUpdateInput) (out model.UserCouponUpdateOutput, err error)
		UserCouponDelete(ctx context.Context, userCouponId uint) (err error)
		UserCouponList(ctx context.Context, in model.UserCouponListInput) (out *model.UserCouponListOutput, err error)
	}
)

var (
	localUserCoupon IUserCoupon
)

func UserCoupon() IUserCoupon {
	if localUserCoupon == nil {
		panic("implement not found for interface IUserCoupon, forgot register?")
	}
	return localUserCoupon
}

func RegisterUserCoupon(i IUserCoupon) {
	localUserCoupon = i
}
