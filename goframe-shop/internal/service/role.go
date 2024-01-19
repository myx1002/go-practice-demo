// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop/internal/model"
)

type (
	IRole interface {
		// 创建角色
		CreateRole(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error)
		// 更新角色
		UpdateRole(ctx context.Context, in model.RoleUpdateInput) (out model.RoleUpdateOutput, err error)
		// Delete 删除
		DeleteRole(ctx context.Context, id uint) error
		// 角色列表
		RoleList(ctx context.Context, in model.RoleListInput) (out *model.RoleListOutput, err error)
		CreateRolePermission(ctx context.Context, in model.RolePermissionCreateInput) (out model.RolePermissionCreateOutput, err error)
		DeleteRolePermission(ctx context.Context, in model.RolePermissionDeleteInput) error
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
