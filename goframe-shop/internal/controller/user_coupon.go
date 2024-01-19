package controller

import (
	"golang.org/x/net/context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var UserCoupon cUserCoupon

type cUserCoupon struct {
}

func (a *cUserCoupon) Create(ctx context.Context, req *backend.UserCouponCreateReq) (res *backend.UserCouponCreateRes, err error) {
	out, err := service.UserCoupon().UserCouponCreate(ctx, model.UserCouponCreateInput{
		UserId:   req.UserId,
		CouponId: req.CouponId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.UserCouponCreateRes{UserCouponId: out.UserCouponId}, nil
}

func (a *cUserCoupon) Delete(ctx context.Context, req *backend.UserCouponDeleteReq) (res *backend.UserCouponDeleteRes, err error) {
	err = service.UserCoupon().UserCouponDelete(ctx, req.UserCouponId)
	return &backend.UserCouponDeleteRes{UserCouponId: req.UserCouponId}, err
}

func (a *cUserCoupon) Update(ctx context.Context, req *backend.UserCouponUpdateReq) (res *backend.UserCouponUpdateRes, err error) {
	out, err := service.UserCoupon().UserCouponUpdate(ctx, model.UserCouponUpdateInput{
		UserCouponId: req.UserCouponId,
		CouponId:     req.CouponId,
		UserId:       req.UserId,
		Status:       req.Status,
	})
	return &backend.UserCouponUpdateRes{UserCouponId: out.UserCouponId}, err
}

// List
func (a *cUserCoupon) List(ctx context.Context, req *backend.UserCouponListReq) (res *backend.UserCouponListRes, err error) {
	getListRes, err := service.UserCoupon().UserCouponList(ctx, model.UserCouponListInput{
		Page:       req.Page,
		Size:       req.Size,
		UserName:   req.UserName,
		CouponName: req.CouponName,
	})
	if err != nil {
		return nil, err
	}

	return &backend.UserCouponListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  getListRes.List,
			Page:  getListRes.Page,
			Size:  getListRes.Size,
			Total: getListRes.Total,
		},
	}, nil
}
