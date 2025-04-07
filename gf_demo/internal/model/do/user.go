// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} // ID
	Username  interface{} // 用户名
	Nickname  interface{} // 昵称
	Password  interface{} // 密码
	Avatar    interface{} // 头像
	CreatedAt *gtime.Time // 创建时间
}
