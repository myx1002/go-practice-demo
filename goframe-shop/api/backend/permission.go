package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PermissionCreateReq struct {
	g.Meta `path:"/permission/add" method:"post" tags:"权限管理" dc:"添加权限"`
	Name   string `json:"name" v:"required#权限名称不能为空" dc:"权限名称"`
	Path   string `json:"path" v:"required#权限路劲不能为空" dc:"权限路径"`
}

type PermissionCreateRes struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionDeleteReq struct {
	g.Meta       `path:"/permission/delete" method:"delete" tags:"权限管理" dc:"编辑权限"`
	PermissionId uint `json:"permission_id"      v:"min:1#请选择需要修改的权限" dc:"权限id"`
}

type PermissionDeleteRes struct {
}

type PermissionUpdateReq struct {
	g.Meta       `path:"/permission/update/{permission_id}" method:"post" tags:"权限管理" dc:"编辑权限"`
	PermissionId uint   `json:"permission_id"      v:"min:1#请选择需要修改的权限" dc:"权限id"`
	Name         string `json:"name" v:"required#权限名称不能为空" dc:"权限名称"`
	Path         string `json:"path" v:"required#权限路劲不能为空" dc:"权限路径"`
}

type PermissionUpdateRes struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionListReq struct {
	g.Meta `path:"/permission/list" method:"get" tags:"权限管理" dc:"权限列表"`
	CommonPaginationReq
}

type PermissionListRes struct {
	CommonPaginationRes
}
