package order

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

type sOrder struct {
}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

func (s *sOrder) PlaceOrder(ctx context.Context, in model.PlaceAnOrderInput) (out *model.PlaceAnOrderOutput, err error) {
	in.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	in.Number = utility.GenerateOrderNumber()
	out = &model.PlaceAnOrderOutput{}
	err = dao.OrderInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 插入订单
		id, err := dao.OrderInfo.Ctx(ctx).InsertAndGetId(in)
		if err != nil {
			return err
		}

		// 插入订单商品信息
		for _, info := range in.OrderGoodsInfo {
			info.OrderId = gconv.Int(id)
			_, err := dao.OrderGoodsInfo.Ctx(ctx).Insert(info)
			if err != nil {
				return err
			}
		}

		for _, info := range in.OrderGoodsInfo {
			// 商品增加销量
			_, err = dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Increment(dao.GoodsInfo.Columns().Sale, info.Count)
			if err != nil {
				return err
			}
			// 商品减少库存
			_, err = dao.GoodsInfo.Ctx(ctx).WherePri(info.GoodsId).Decrement(dao.GoodsInfo.Columns().Stock, info.Count)
			if err != nil {
				return err
			}
			// 商品规格减少库存
			_, err = dao.GoodsOptionsInfo.Ctx(ctx).WherePri(info.GoodsOptionsId).Decrement(dao.GoodsOptionsInfo.Columns().Stock, info.Count)
			if err != nil {
				return err
			}
		}
		out.OrderId = gconv.Int(id)
		return nil
	})
	return out, err
}
