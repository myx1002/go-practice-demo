package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type RoleCreateUpdateCommonInput struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type RoleCreateInput struct {
	RoleCreateUpdateCommonInput
}

type RoleCreateOutput struct {
	RoleId uint `json:"role_id"`
}

type RoleUpdateInput struct {
	RoleCreateUpdateCommonInput
	Id uint `json:"role_id"`
}

type RoleUpdateOutput struct {
	RoleId uint `json:"role_id"`
}

type RoleListInput struct {
	Page int
	Size int
}

type RoleListOutputItem struct {
	Id        uint        `json:"role_id"`
	Name      string      `json:"name"`
	Desc      string      `json:"desc"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type RoleListOutput struct {
	List  []RoleListOutputItem `json:"list" description:"列表"`
	Page  int                  `json:"page" description:"分页码"`
	Size  int                  `json:"size" description:"分页数量"`
	Total int                  `json:"total" description:"数据总数"`
}

type RolePermissionCreateInput struct {
	RoleId       uint `json:"role_id"`
	PermissionId uint `json:"permission_id"`
}

type RolePermissionCreateOutput struct {
	RoleId uint `json:"role_id"`
}

type RolePermissionDeleteInput struct {
	RoleId       uint `json:"role_id"`
	PermissionId uint `json:"permission_id"`
}
