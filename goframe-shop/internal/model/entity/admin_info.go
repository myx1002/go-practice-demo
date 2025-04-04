// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure for table admin_info.
type AdminInfo struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:"用户名"`
	Password  string      `json:"password"  description:"密码"`
	RoleIds   string      `json:"roleIds"   description:"角色ids"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
	UserSalt  string      `json:"userSalt"  description:"加密盐"`
	IsAdmin   int         `json:"isAdmin"   description:"是否超级管理员"`
}
