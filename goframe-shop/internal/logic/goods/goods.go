package goods

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/logic/collection"
	"goframe-shop/internal/logic/praise"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sGoods struct {
}

func init() {
	service.RegisterGoods(New())
}

func New() *sGoods {
	return &sGoods{}
}

func (s *sGoods) GoodsCreate(ctx context.Context, in model.GoodsCreateInput) (out *model.GoodsCreateOutput, err error) {
	// 查询所选分类的商品是否已存在
	detail, err := dao.GoodsInfo.Ctx(ctx).Where(g.Map{
		dao.GoodsInfo.Columns().Name:             in.Name,
		dao.GoodsInfo.Columns().Level1CategoryId: in.Level1CategoryId,
		dao.GoodsInfo.Columns().Level2CategoryId: in.Level2CategoryId,
		dao.GoodsInfo.Columns().Level2CategoryId: in.Level2CategoryId,
	}).One()

	if err != nil {
		return nil, err
	}
	if detail != nil {
		return nil, gerror.New("该商品已存在，请勿重复添加！")
	}

	userCouponId, err := dao.GoodsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.GoodsCreateOutput{GoodsId: uint(userCouponId)}, nil
}

func (s *sGoods) GoodsUpdate(ctx context.Context, in model.GoodsUpdateInput) (out model.GoodsUpdateOutput, err error) {
	detail, err := dao.GoodsInfo.Ctx(ctx).
		WhereNot(dao.GoodsInfo.Columns().Id, in.GoodsId).
		Where(g.Map{
			dao.GoodsInfo.Columns().Name:             in.Name,
			dao.GoodsInfo.Columns().Level1CategoryId: in.Level1CategoryId,
			dao.GoodsInfo.Columns().Level2CategoryId: in.Level2CategoryId,
			dao.GoodsInfo.Columns().Level2CategoryId: in.Level2CategoryId,
		}).One()

	if err != nil {
		return model.GoodsUpdateOutput{}, err
	}

	if detail != nil {
		return model.GoodsUpdateOutput{}, gerror.New("该商品已存在，请重新编辑！")
	}

	err = dao.GoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.GoodsInfo.Ctx(ctx).WherePri(in.GoodsId).Data(in).Update()
		return err
	})
	return model.GoodsUpdateOutput{GoodsId: in.GoodsId}, err
}

func (s *sGoods) GoodsDelete(ctx context.Context, goodsId uint) (err error) {
	err = dao.GoodsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.GoodsInfo.Ctx(ctx).WherePri(goodsId).Delete()
		return err
	})

	return
}

func (s *sGoods) GoodsList(ctx context.Context, in model.GoodsListInput) (out *model.GoodsListOutput, err error) {
	m := dao.GoodsInfo.Ctx(ctx).OmitEmpty()
	out = &model.GoodsListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	if in.Name != "" {
		m = m.WhereLike(dao.GoodsInfo.Columns().Name, "%"+in.Name+"%")
	}

	err = m.OrderDesc(dao.GoodsInfo.Columns().Id).
		WithAll().
		Page(in.Page, in.Size).
		ScanAndCount(&out.List, &out.Total, true)
	return
}

func (s *sGoods) GoodsDetail(ctx context.Context, input model.GoodsDetailInput) (out *model.GoodsDetailOutput, err error) {
	err = dao.GoodsInfo.Ctx(ctx).WithAll().WherePri(input.GoodsId).Scan(&out)
	if err != nil {
		return nil, err
	}
	out.IsPraise, _ = praise.CheckObjectIsPraise(ctx, uint(input.GoodsId), 1)
	out.PraiseCount, _ = praise.ComputePraiseCountByObjectId(ctx, uint(input.GoodsId), 1)
	out.IsCollection, _ = collection.CheckObjectIsCollection(ctx, uint(input.GoodsId), 1)
	out.CollectionCount, _ = collection.ComputeCollectionCountByObjectId(ctx, uint(input.GoodsId), 1)
	return
}
