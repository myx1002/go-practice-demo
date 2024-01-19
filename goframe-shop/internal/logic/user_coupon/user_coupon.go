package coupon

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/service"
)

type sUserCoupon struct {
}

func init() {
	service.RegisterUserCoupon(New())
}

func New() *sUserCoupon {
	return &sUserCoupon{}
}

func (s *sUserCoupon) UserCouponCreate(ctx context.Context, in model.UserCouponCreateInput) (out model.UserCouponCreateOutput, err error) {
	// 查询用户是否存在该优惠卷
	detail, err := dao.UserCouponInfo.Ctx(ctx).Where(g.Map{
		dao.UserCouponInfo.Columns().UserId:   in.UserId,
		dao.UserCouponInfo.Columns().CouponId: in.CouponId,
	}).One()

	if err != nil {
		return model.UserCouponCreateOutput{}, err
	}
	if detail != nil {
		return model.UserCouponCreateOutput{}, gerror.New("该用户已存在所选优惠卷，请勿重复添加！")
	}

	userCouponId, err := dao.UserCouponInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.UserCouponCreateOutput{}, err
	}

	return model.UserCouponCreateOutput{UserCouponId: uint(userCouponId)}, nil
}

func (s *sUserCoupon) UserCouponUpdate(ctx context.Context, in model.UserCouponUpdateInput) (out model.UserCouponUpdateOutput, err error) {
	// 查询用户是否存在该优惠卷
	detail, err := dao.UserCouponInfo.Ctx(ctx).
		WhereNot(dao.UserCouponInfo.Columns().Id, in.UserCouponId).
		Where(g.Map{
			dao.UserCouponInfo.Columns().UserId:   in.UserId,
			dao.UserCouponInfo.Columns().CouponId: in.CouponId,
		}).One()

	if err != nil {
		return model.UserCouponUpdateOutput{}, err
	}

	if detail != nil {
		return model.UserCouponUpdateOutput{}, gerror.New("该用户已存在所选优惠卷，请重新编辑！")
	}

	err = dao.UserCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.UserCouponInfo.Ctx(ctx).WherePri(in.UserCouponId).Data(in).Update()
		return err
	})
	return model.UserCouponUpdateOutput{UserCouponId: in.UserCouponId}, err
}

func (s *sUserCoupon) UserCouponDelete(ctx context.Context, userCouponId uint) (err error) {
	err = dao.UserCouponInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.UserCouponInfo.Ctx(ctx).WherePri(userCouponId).Delete()
		return err
	})

	return
}

func (s *sUserCoupon) UserCouponList(ctx context.Context, in model.UserCouponListInput) (out *model.UserCouponListOutput, err error) {
	m := dao.UserCouponInfo.Ctx(ctx).OmitEmpty()
	out = &model.UserCouponListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	if in.UserName != "" {
		userId, err := dao.UserInfo.Ctx(ctx).Fields(dao.UserInfo.Columns().Id).Where(do.UserInfo{Name: in.UserName}).Value()
		if err != nil {
			return &model.UserCouponListOutput{}, err
		}
		m = m.Where(dao.UserCouponInfo.Columns().UserId, userId)
	}

	if in.CouponName != "" {
		couponId, err := dao.CouponInfo.Ctx(ctx).Fields(dao.CouponInfo.Columns().Id).Where(do.CouponInfo{Name: in.CouponName}).Value()
		if err != nil {
			return &model.UserCouponListOutput{}, err
		}
		m = m.Where(dao.UserCouponInfo.Columns().CouponId, couponId)
	}

	err = m.OrderDesc(dao.UserCouponInfo.Columns().Id).
		WithAll().
		Page(in.Page, in.Size).
		ScanAndCount(&out.List, &out.Total, true)
	return
}
