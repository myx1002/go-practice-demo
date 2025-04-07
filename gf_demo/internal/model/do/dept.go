// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Dept is the golang structure of table dept for DAO operations like Where/Data.
type Dept struct {
	g.Meta `orm:"table:dept, do:true"`
	Id     interface{} // ID
	Pid    interface{} // 上级部门ID
	Name   interface{} // 部门名称
	Leader interface{} // 部门领导
	Phone  interface{} // 联系电话
}
