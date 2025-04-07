// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Hobby is the golang structure of table hobby for DAO operations like Where/Data.
type Hobby struct {
	g.Meta `orm:"table:hobby, do:true"`
	Id     interface{} // ID
	EmpId  interface{} // EmpID
	Hobby  interface{} // 爱好
}
