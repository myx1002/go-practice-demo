package goods

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sArticle struct {
}

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

func (s *sArticle) ArticleCreate(ctx context.Context, in model.ArticleCreateInput) (out *model.ArticleCreateOutput, err error) {
	id, err := dao.ArticleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.ArticleCreateOutput{ArticleId: uint(id)}, nil
}

func (s *sArticle) ArticleUpdate(ctx context.Context, in model.ArticleUpdateInput) (out model.ArticleUpdateOutput, err error) {
	err = dao.ArticleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.ArticleInfo.Ctx(ctx).WherePri(in.ArticleId).Data(in).Update()
		return err
	})
	return model.ArticleUpdateOutput{ArticleId: in.ArticleId}, err
}

func (s *sArticle) ArticleDelete(ctx context.Context, goodsId uint) (err error) {
	err = dao.ArticleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.ArticleInfo.Ctx(ctx).WherePri(goodsId).Delete()
		return err
	})

	return
}

func (s *sArticle) ArticleList(ctx context.Context, in model.ArticleListInput) (out *model.ArticleListOutput, err error) {
	m := dao.ArticleInfo.Ctx(ctx).OmitEmpty()
	out = &model.ArticleListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	if in.Title != "" {
		m = m.WhereLike(dao.ArticleInfo.Columns().Title, "%"+in.Title+"%")
	}

	err = m.OrderDesc(dao.ArticleInfo.Columns().Id).
		WithAll().
		Page(in.Page, in.Size).
		ScanAndCount(&out.List, &out.Total, true)
	return
}
