package permission

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sPermission struct {
}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) CreatePermission(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	permissionId, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return model.PermissionCreateOutput{}, err
	}
	return model.PermissionCreateOutput{
		PermissionId: uint(permissionId),
	}, nil
}

func (s *sPermission) UpdatePermission(ctx context.Context, in model.PermissionUpdateInput) (out model.PermissionUpdateOutput, err error) {
	err = dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.PermissionInfo.Ctx(ctx).Data(in).WherePri(in.PermissionId).Update()
		return err
	})

	return model.PermissionUpdateOutput{
		PermissionId: in.PermissionId,
	}, err
}

func (s *sPermission) DeletePermission(ctx context.Context, in model.PermissionDeleteInput) (err error) {
	return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.PermissionInfo.Ctx(ctx).Data(in).WherePri(in.PermissionId).Delete()
		return err
	})
}

func (s *sPermission) PermissionList(ctx context.Context, in model.PermissionListInput) (out *model.PermissionListOutput, err error) {
	m := dao.PermissionInfo.Ctx(ctx)

	out = &model.PermissionListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size).OrderDesc(dao.PermissionInfo.Columns().Id)

	// 执行查询，
	var list []*entity.PermissionInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}

	// 没有数据
	if len(list) == 0 {
		return out, nil
	}

	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}

	return
}
