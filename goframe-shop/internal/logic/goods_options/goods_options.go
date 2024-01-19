package goods

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sGoodsOptions struct {
}

func init() {
	service.RegisterGoodsOptions(New())
}

func New() *sGoodsOptions {
	return &sGoodsOptions{}
}

func (s *sGoodsOptions) GoodsOptionsCreate(ctx context.Context, in model.GoodsOptionsCreateInput) (out *model.GoodsOptionsCreateOutput, err error) {
	detail, err := dao.GoodsOptionsInfo.Ctx(ctx).Where(g.Map{
		dao.GoodsOptionsInfo.Columns().Name:    in.Name,
		dao.GoodsOptionsInfo.Columns().GoodsId: in.GoodsId,
	}).One()

	if err != nil {
		return nil, err
	}
	if detail != nil {
		return nil, gerror.New("该商品规格已存在，请勿重复添加！")
	}

	id, err := dao.GoodsOptionsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.GoodsOptionsCreateOutput{GoodsOptionsId: uint(id)}, nil
}

func (s *sGoodsOptions) GoodsOptionsUpdate(ctx context.Context, in model.GoodsOptionsUpdateInput) (out model.GoodsOptionsUpdateOutput, err error) {
	detail, err := dao.GoodsOptionsInfo.Ctx(ctx).
		WhereNot(dao.GoodsOptionsInfo.Columns().Id, in.GoodsOptionsId).
		Where(g.Map{
			dao.GoodsOptionsInfo.Columns().Name:    in.Name,
			dao.GoodsOptionsInfo.Columns().GoodsId: in.GoodsId,
		}).One()

	if err != nil {
		return model.GoodsOptionsUpdateOutput{}, err
	}

	if detail != nil {
		return model.GoodsOptionsUpdateOutput{}, gerror.New("该商品规格已存在，请重新编辑！")
	}

	err = dao.GoodsOptionsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.GoodsOptionsInfo.Ctx(ctx).WherePri(in.GoodsOptionsId).Data(in).Update()
		return err
	})
	return model.GoodsOptionsUpdateOutput{GoodsOptionsId: in.GoodsOptionsId}, err
}

func (s *sGoodsOptions) GoodsOptionsDelete(ctx context.Context, goodsId uint) (err error) {
	err = dao.GoodsOptionsInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.GoodsOptionsInfo.Ctx(ctx).WherePri(goodsId).Delete()
		return err
	})

	return
}

func (s *sGoodsOptions) GoodsOptionsList(ctx context.Context, in model.GoodsOptionsListInput) (out *model.GoodsOptionsListOutput, err error) {
	m := dao.GoodsOptionsInfo.Ctx(ctx).OmitEmpty()
	out = &model.GoodsOptionsListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	if in.Name != "" {
		m = m.WhereLike(dao.GoodsOptionsInfo.Columns().Name, "%"+in.Name+"%")
	}

	err = m.OrderDesc(dao.GoodsOptionsInfo.Columns().Id).
		WithAll().
		Page(in.Page, in.Size).
		ScanAndCount(&out.List, &out.Total, true)
	return
}
