// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure of table file_info for DAO operations like Where/Data.
type FileInfo struct {
	g.Meta       `orm:"table:file_info, do:true"`
	Id           interface{} // 文件id
	Name         interface{} // 文件名称
	RealFileName interface{} //
	Src          interface{} // 文件路径
	Url          interface{} // 访问地址
	UserId       interface{} // 用户id
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
