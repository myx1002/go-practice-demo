package role

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sRole struct {
}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// 创建角色
func (s *sRole) CreateRole(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	roleId, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}

	return model.RoleCreateOutput{
		RoleId: uint(roleId),
	}, nil
}

// 更新角色
func (s *sRole) UpdateRole(ctx context.Context, in model.RoleUpdateInput) (out model.RoleUpdateOutput, err error) {
	err = dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}

		_, err := dao.RoleInfo.Ctx(ctx).WherePri(in.Id).Update(in)
		return err
	})

	return model.RoleUpdateOutput{RoleId: in.Id}, err
}

// Delete 删除
func (s *sRole) DeleteRole(ctx context.Context, id uint) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容-软删除
		// 如果要硬删除，可加.Unscoped()来忽略时间特性
		_, err := dao.RoleInfo.Ctx(ctx).WherePri(id).Delete()
		return err
	})
}

// 角色列表
func (s *sRole) RoleList(ctx context.Context, in model.RoleListInput) (out *model.RoleListOutput, err error) {
	var (
		m = dao.RoleInfo.Ctx(ctx)
	)
	out = &model.RoleListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	// 排序方式
	listModel := m.Page(in.Page, in.Size).OrderDesc(dao.RoleInfo.Columns().Id)

	// 执行查询
	var list []*entity.RoleInfo
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
	// Rotation
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (s *sRole) CreateRolePermission(ctx context.Context, in model.RolePermissionCreateInput) (out model.RolePermissionCreateOutput, err error) {
	_, err = dao.RolePermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}

	return model.RolePermissionCreateOutput{
		RoleId: in.RoleId,
	}, nil
}

func (s *sRole) DeleteRolePermission(ctx context.Context, in model.RolePermissionDeleteInput) error {
	return dao.RoleInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容-软删除
		// 如果要硬删除，可加.Unscoped()来忽略时间特性
		_, err := dao.RolePermissionInfo.Ctx(ctx).Unscoped().Where(g.Map{
			dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
			dao.RolePermissionInfo.Columns().PermissionId: in.PermissionId,
		}).Delete()
		return err
	})
}
