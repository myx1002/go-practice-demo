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
	ICoupon interface {
		CouponCreate(ctx context.Context, input model.CouponCreateInput) (output model.CouponCreateOutput, err error)
		CouponUpdate(ctx context.Context, input model.CouponUpdateInput) (output model.CouponUpdateOutput, err error)
		CouponDelete(ctx context.Context, couponId uint) (err error)
		CouponList(ctx context.Context, input model.CouponListInput) (output *model.CouponListOutput, err error)
	}
)

var (
	localCoupon ICoupon
)

func Coupon() ICoupon {
	if localCoupon == nil {
		panic("implement not found for interface ICoupon, forgot register?")
	}
	return localCoupon
}

func RegisterCoupon(i ICoupon) {
	localCoupon = i
}
