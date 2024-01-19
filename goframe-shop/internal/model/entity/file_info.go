// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure for table file_info.
type FileInfo struct {
	Id           int         `json:"id"           description:"文件id"`
	Name         string      `json:"name"         description:"文件名称"`
	RealFileName string      `json:"realFileName" description:""`
	Src          string      `json:"src"          description:"文件路径"`
	Url          string      `json:"url"          description:"访问地址"`
	UserId       int         `json:"userId"       description:"用户id"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:"更新时间"`
	DeletedAt    *gtime.Time `json:"deletedAt"    description:"删除时间"`
}
