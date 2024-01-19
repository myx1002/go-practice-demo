package controller

import (
	"context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Role cRole

type cRole struct {
}

// 创建角色
func (c *cRole) CreateRole(ctx context.Context, req *backend.RoleReq) (res *backend.RoleRes, err error) {
	out, err := service.Role().CreateRole(ctx, model.RoleCreateInput{
		RoleCreateUpdateCommonInput: model.RoleCreateUpdateCommonInput{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleRes{
		RoleId: out.RoleId,
	}, nil
}

// 更新角色
func (c *cRole) UpdateRole(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	out, err := service.Role().UpdateRole(ctx, model.RoleUpdateInput{
		RoleCreateUpdateCommonInput: model.RoleCreateUpdateCommonInput{
			Name: req.Name,
			Desc: req.Desc,
		},
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleUpdateRes{
		RoleId: out.RoleId,
	}, nil
}

// 删除角色
func (c *cRole) DeleteRole(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().DeleteRole(ctx, req.Id)
	return
}

// 角色列表
func (c *cRole) RoleList(ctx context.Context, req *backend.RoleListReq) (res *backend.RoleListRes, err error) {
	getListRes, err := service.Role().RoleList(ctx, model.RoleListInput{Size: req.Size, Page: req.Page})
	if err != nil {
		return nil, err
	}
	return &backend.RoleListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// 更新角色-权限
func (c *cRole) CreateRolePermission(ctx context.Context, req *backend.RolePermissionCreateReq) (res *backend.RolePermissionCreateRes, err error) {
	out, err := service.Role().CreateRolePermission(ctx, model.RolePermissionCreateInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RolePermissionCreateRes{
		RoleId: out.RoleId,
	}, nil
}

// 删除角色-权限
func (c *cRole) DeleteRolePermission(ctx context.Context, req *backend.RolePermissionDeleteReq) (res *backend.RolePermissionDeleteRes, err error) {
	err = service.Role().DeleteRolePermission(ctx, model.RolePermissionDeleteInput{
		PermissionId: req.PermissionId,
		RoleId:       req.RoleId,
	})
	return
}
