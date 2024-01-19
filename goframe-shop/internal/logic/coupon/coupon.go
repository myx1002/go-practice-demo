package coupon

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sCoupon struct {
}

func init() {
	service.RegisterCoupon(New())
}

func New() *sCoupon {
	return &sCoupon{}
}

func (s *sCoupon) CouponCreate(ctx context.Context, input model.CouponCreateInput) (output model.CouponCreateOutput, err error) {
	couponId, err := dao.CouponInfo.Ctx(ctx).Data(input).InsertAndGetId()
	if err != nil {
		return model.CouponCreateOutput{}, err
	}

	return model.CouponCreateOutput{CouponId: uint(couponId)}, nil
}

func (s *sCoupon) CouponUpdate(ctx context.Context, input model.CouponUpdateInput) (output model.CouponUpdateOutput, err error) {
	err = dao.CouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.CouponInfo.Ctx(ctx).WherePri(input.CouponId).Data(input).Update()
		return err
	})
	return model.CouponUpdateOutput{CouponId: input.CouponId}, err
}

func (s *sCoupon) CouponDelete(ctx context.Context, couponId uint) (err error) {
	err = dao.CategoryInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.CategoryInfo.Ctx(ctx).WherePri(couponId).Delete()
		return err
	})

	return
}

func (s *sCoupon) CouponList(ctx context.Context, input model.CouponListInput) (output *model.CouponListOutput, err error) {
	m := dao.CouponInfo.Ctx(ctx)
	output = &model.CouponListOutput{
		Size: input.Size,
		Page: input.Page,
	}

	output.Total, err = m.Count()
	if err != nil {
		return output, err
	}

	if err := m.Page(input.Page, input.Size).OrderDesc(dao.CouponInfo.Columns().Id).Scan(&output.List); err != nil {
		return output, err
	}
	return
}
