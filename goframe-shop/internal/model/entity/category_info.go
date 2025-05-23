// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryInfo is the golang structure for table category_info.
type CategoryInfo struct {
	Id        int         `json:"id"        description:""`
	ParentId  int         `json:"parentId"  description:"父级id"`
	Name      string      `json:"name"      description:""`
	PicUrl    string      `json:"picUrl"    description:"icon"`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	Level     int         `json:"level"     description:"等级 默认1级分类"`
	Sort      int         `json:"sort"      description:""`
}
