package controller

import (
	"context"

	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

var Permission cPermission

type cPermission struct {
}

func (c *cPermission) CreatePermission(ctx context.Context, req *backend.PermissionCreateReq) (res *backend.PermissionCreateRes, err error) {
	out, err := service.Permission().CreatePermission(ctx, model.PermissionCreateInput{
		Name: req.Name,
		Path: req.Path,
	})
	if err != nil {
		return nil, err
	}

	return &backend.PermissionCreateRes{
		PermissionId: out.PermissionId,
	}, nil
}

func (c *cPermission) UpdatePermission(ctx context.Context, req *backend.PermissionUpdateReq) (res *backend.PermissionUpdateRes, err error) {
	out, err := service.Permission().UpdatePermission(ctx, model.PermissionUpdateInput{
		PermissionId: req.PermissionId,
		Name:         req.Name,
		Path:         req.Path,
	})
	if err != nil {
		return nil, err
	}

	return &backend.PermissionUpdateRes{
		PermissionId: out.PermissionId,
	}, nil
}

func (c *cPermission) DeletePermission(ctx context.Context, req *backend.PermissionDeleteReq) (res *backend.PermissionDeleteRes, err error) {
	err = service.Permission().DeletePermission(ctx, model.PermissionDeleteInput{
		PermissionId: req.PermissionId,
	})
	return
}

func (c *cPermission) PermissionList(ctx context.Context, req *backend.PermissionListReq) (res *backend.PermissionListRes, err error) {
	out, err := service.Permission().PermissionList(ctx, model.PermissionListInput{
		Size: req.Size,
		Page: req.Page,
	})
	return &backend.PermissionListRes{
		CommonPaginationRes: backend.CommonPaginationRes{
			List:  out.List,
			Total: out.Total,
			Size:  out.Size,
			Page:  out.Page,
		},
	}, err
}
