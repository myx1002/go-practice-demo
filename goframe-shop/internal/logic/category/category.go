package category

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sCategory struct {
}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

func (s *sCategory) CategoryCreate(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {
	one, err := dao.CategoryInfo.Ctx(ctx).Where(g.Map{
		dao.CategoryInfo.Columns().Name:     in.Name,
		dao.CategoryInfo.Columns().ParentId: in.ParentId,
		dao.CategoryInfo.Columns().Level:    in.Level,
	}).One()
	if err != nil {
		return model.CategoryCreateOutput{}, err
	}

	if one != nil {
		return model.CategoryCreateOutput{}, gerror.New("该分类已存在，请勿重复添加！")
	}

	categoryId, err := dao.CategoryInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.CategoryCreateOutput{}, err
	}

	return model.CategoryCreateOutput{
		CategoryId: uint(categoryId),
	}, err
}

func (s *sCategory) CategoryUpdate(ctx context.Context, in model.CategoryUpdateInput) (out model.CategoryUpdateOutput, err error) {
	// 应该写在logic层，这里只是通用的更新方法
	//one, err := dao.CategoryInfo.Ctx(ctx).Where(g.Map{
	//	dao.CategoryInfo.Columns().Name:     in.Name,
	//	dao.CategoryInfo.Columns().ParentId: in.ParentId,
	//	dao.CategoryInfo.Columns().Level:    in.Level,
	//}).WhereNot(dao.CategoryInfo.Columns().Id, in.CategoryId).One()
	//if err != nil {
	//	return model.CategoryUpdateOutput{}, err
	//}
	//
	//if one != nil {
	//	return model.CategoryUpdateOutput{}, gerror.New("该分类已存在，请重新编辑！")
	//}
	err = dao.CategoryInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.CategoryInfo.Ctx(ctx).Data(in).WherePri(in.CategoryId).Update()
		return err
	})

	return model.CategoryUpdateOutput{CategoryId: in.CategoryId}, err

}

func (s *sCategory) CategoryDelete(ctx context.Context, CategoryId uint) error {
	return dao.CategoryInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.PositionInfo.Ctx(ctx).Where(g.Map{
			dao.PositionInfo.Columns().Id: CategoryId,
		}).Delete()
		return err
	})
}

func (s *sCategory) CategoryList(ctx context.Context, in model.CategoryListInput) (out *model.CategoryListOutput, err error) {
	m := dao.CategoryInfo.Ctx(ctx)
	out = &model.CategoryListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size).OrderDesc(dao.CategoryInfo.Columns().Id)
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	out.Total, err = m.Count()
	return out, err
}
