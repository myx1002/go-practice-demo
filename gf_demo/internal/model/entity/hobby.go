// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Hobby is the golang structure for table hobby.
type Hobby struct {
	Id    uint   `json:"id"     orm:"id"     description:"ID"`    // ID
	EmpId uint   `json:"emp_id" orm:"emp_id" description:"EmpID"` // EmpID
	Hobby string `json:"hobby"  orm:"hobby"  description:"爱好"`    // 爱好
}
