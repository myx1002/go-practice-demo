package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RoleReq struct {
	g.Meta `path:"/role/add" method:"post" tags:"角色管理" dc:"添加角色"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleRes struct {
	RoleId uint `json:"role_id"`
}

// 删除角色
type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"delete" tags:"角色管理" dc:"删除角色"`
	Id     uint `json:"role_id" v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RoleDeleteRes struct{}

// 编辑角色
type RoleUpdateReq struct {
	g.Meta `path:"/role/update/{Id}" method:"post" tags:"角色管理" dc:"修改角色"`
	Id     uint   `json:"role_id"      v:"min:1#请选择需要修改的角色" dc:"角色id"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}
type RoleUpdateRes struct {
	RoleId uint `json:"role_id"`
}

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色管理" dc:"角色列表"`
	CommonPaginationReq
}

type RoleListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

// 删除角色
type RolePermissionDeleteReq struct {
	g.Meta       `path:"/role/permission/delete" method:"delete" tags:"角色管理" dc:"删除角色权限"`
	RoleId       uint `json:"role_id" v:"min:1#请选择需要删除的角色" dc:"角色id"`
	PermissionId uint `json:"permission_id" v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RolePermissionDeleteRes struct{}

// 编辑角色
type RolePermissionCreateReq struct {
	g.Meta       `path:"/role/permission/add" method:"post" tags:"角色管理" dc:"修改角色权限"`
	RoleId       uint `json:"role_id" v:"min:1#请选择需要删除的角色" dc:"角色id"`
	PermissionId uint `json:"permission_id" v:"min:1#请选择需要删除的角色" dc:"角色id"`
}
type RolePermissionCreateRes struct {
	RoleId uint `json:"role_id"`
}
