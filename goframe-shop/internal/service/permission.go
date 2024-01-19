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
	IPermission interface {
		CreatePermission(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error)
		UpdatePermission(ctx context.Context, in model.PermissionUpdateInput) (out model.PermissionUpdateOutput, err error)
		DeletePermission(ctx context.Context, in model.PermissionDeleteInput) (err error)
		PermissionList(ctx context.Context, in model.PermissionListInput) (out *model.PermissionListOutput, err error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
