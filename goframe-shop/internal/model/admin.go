package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminCreateUpdateBase 创建/修改内容基类
type AdminCreateUpdateBase struct {
	Id       uint
	Name     string
	Password string
	RoleIds  uint
	IsAdmin  uint
	UserSalt string
}

// AdminCreateInput 创建内容
type AdminCreateInput struct {
	AdminCreateUpdateBase
}

// AdminCreateOutput 创建内容返回结果
type AdminCreateOutput struct {
	AdminId uint `json:"admin_id"`
}

// AdminUpdateInput 修改内容
type AdminUpdateInput struct {
	AdminCreateUpdateBase
	Id uint
}

// AdminGetListInput 获取内容列表
type AdminGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

type AdminGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	Name      string      `json:"name"`
	RoleIds   uint        `json:"role_ids"`
	IsAdmin   uint        `json:"is_admin"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

// AdminGetListOutput 查询列表结果
type AdminGetListOutput struct {
	List  []AdminGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}
