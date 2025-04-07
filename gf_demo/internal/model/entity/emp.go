// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Emp is the golang structure for table emp.
type Emp struct {
	Id     uint   `json:"id"      orm:"id"      description:"ID"`          // ID
	DeptId uint   `json:"dept_id" orm:"dept_id" description:"所属部门"`        // 所属部门
	Name   string `json:"name"    orm:"name"    description:"姓名"`          // 姓名
	Gender int    `json:"gender"  orm:"gender"  description:"性别: 0=男 1=女"` // 性别: 0=男 1=女
	Phone  string `json:"phone"   orm:"phone"   description:"联系电话"`        // 联系电话
	Email  string `json:"email"   orm:"email"   description:"邮箱"`          // 邮箱
	Avatar string `json:"avatar"  orm:"avatar"  description:"照片"`          // 照片
}
