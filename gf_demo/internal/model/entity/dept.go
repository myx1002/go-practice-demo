// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Dept is the golang structure for table dept.
type Dept struct {
	Id     uint   `json:"id"     orm:"id"     description:"ID"`     // ID
	Pid    uint   `json:"pid"    orm:"pid"    description:"上级部门ID"` // 上级部门ID
	Name   string `json:"name"   orm:"name"   description:"部门名称"`   // 部门名称
	Leader string `json:"leader" orm:"leader" description:"部门领导"`   // 部门领导
	Phone  string `json:"phone"  orm:"phone"  description:"联系电话"`   // 联系电话
}
