// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Emp is the golang structure of table emp for DAO operations like Where/Data.
type Emp struct {
	g.Meta `orm:"table:emp, do:true"`
	Id     interface{} // ID
	DeptId interface{} // 所属部门
	Name   interface{} // 姓名
	Gender interface{} // 性别: 0=男 1=女
	Phone  interface{} // 联系电话
	Email  interface{} // 邮箱
	Avatar interface{} // 照片
}
