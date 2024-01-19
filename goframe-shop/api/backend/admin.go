package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminReq struct {
	g.Meta   `path:"/admin/add" tags:"管理员" method:"post" summary:"添加管理员"`
	Name     string `json:"name" v:"required#管理员名称不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  uint   `json:"role_ids"     dc:"角色ids" `
	IsAdmin  uint   `json:"is_admin"    v:"required#是否超级管理员不能为空" dc:"是否超级管理员"`
}

type AdminRes struct {
	AdminId uint `json:"admin_id"`
}

type AdminDeleteReq struct {
	g.Meta `path:"/admin/delete" method:"delete" tags:"管理员" summary:"删除管理员"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"管理员id"`
}
type AdminDeleteRes struct{}

type AdminUpdateReq struct {
	g.Meta   `path:"/admin/update/{Id}" method:"post" tags:"管理员" summary:"修改管理员"`
	Id       uint   `json:"id"      v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	Name     string `json:"name" v:"required#管理员名称不能为空" dc:"用户名"`
	Password string `json:"password"    v:"required#密码不能为空" dc:"密码"`
	RoleIds  uint   `json:"role_ids"     dc:"角色ids"`
	IsAdmin  uint   `json:"is_admin"    v:"required#是否超级管理员不能为空" dc:"是否超级管理员"`
}
type AdminUpdateRes struct{}

type AdminGetListCommonReq struct {
	g.Meta `path:"/admin/list" method:"get" tags:"管理员" summary:"管理员列表"`
	CommonPaginationReq
}

type AdminGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type AdminGetInfoReq struct {
	g.Meta `path:"/admin/info" method:"get" tags:"管理员" summary:"管理员详情"`
}

// for jwt
//type AdminGetInfoRes struct {
//	Id          int    `json:"id"`
//	IdentityKey string `json:"identity_key"`
//	Payload     string `json:"payload"`
//}

// for gtoken
type AdminGetInfoRes struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	RoleIds string `json:"role_ids"`
	IsAdmin uint   `json:"is_admin"`
}
