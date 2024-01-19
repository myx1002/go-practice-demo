package controller

import (
	"golang.org/x/net/context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Coupon cCoupon

type cCoupon struct {
}

func (a *cCoupon) Create(ctx context.Context, req *backend.CouponCreateReq) (res *backend.CouponCreateRes, err error) {
	out, err := service.Coupon().CouponCreate(ctx, model.CouponCreateInput{
		Name:       req.Name,
		GoodsIds:   req.GoodsIds,
		CategoryId: req.CategoryId,
		Price:      req.Price,
	})
	if err != nil {
		return nil, err
	}
	return &backend.CouponCreateRes{CouponId: out.CouponId}, nil
}

func (a *cCoupon) Delete(ctx context.Context, req *backend.CouponDeleteReq) (res *backend.CouponDeleteRes, err error) {
	err = service.Coupon().CouponDelete(ctx, req.CouponId)
	return &backend.CouponDeleteRes{CouponId: req.CouponId}, err
}

func (a *cCoupon) Update(ctx context.Context, req *backend.CouponUpdateReq) (res *backend.CouponUpdateRes, err error) {
	out, err := service.Coupon().CouponUpdate(ctx, model.CouponUpdateInput{
		CouponId:   req.CouponId,
		Name:       req.Name,
		GoodsIds:   req.GoodsIds,
		CategoryId: req.CategoryId,
		Price:      req.Price,
	})
	return &backend.CouponUpdateRes{CouponId: out.CouponId}, err
}

// List
func (a *cCoupon) List(ctx context.Context, req *backend.CouponListReq) (res *backend.CouponListRes, err error) {
	getListRes, err := service.Coupon().CouponList(ctx, model.CouponListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.CouponListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  getListRes.List,
			Page:  getListRes.Page,
			Size:  getListRes.Size,
			Total: getListRes.Total,
		},
	}, nil
}
