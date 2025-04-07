// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"         orm:"id"         description:"ID"`   // ID
	Username  string      `json:"username"   orm:"username"   description:"用户名"`  // 用户名
	Nickname  string      `json:"nickname"   orm:"nickname"   description:"昵称"`   // 昵称
	Password  string      `json:"password"   orm:"password"   description:"密码"`   // 密码
	Avatar    string      `json:"avatar"     orm:"avatar"     description:"头像"`   // 头像
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"` // 创建时间
}
