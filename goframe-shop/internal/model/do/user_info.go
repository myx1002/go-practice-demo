// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta         `orm:"table:user_info, do:true"`
	Id             interface{} //
	Name           interface{} // 用户名
	Avatar         interface{} // 头像
	Password       interface{} //
	UserSalt       interface{} // 加密盐 生成密码用
	Sex            interface{} // 1男 2女
	Status         interface{} // 1正常 2拉黑冻结
	Sign           interface{} // 个性签名
	SecretQuestion interface{} // 密保问题
	SecretAnswer   interface{} // 密保问题的答案
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
}
