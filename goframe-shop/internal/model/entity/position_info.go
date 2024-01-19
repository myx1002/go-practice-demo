// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure for table position_info.
type PositionInfo struct {
	Id        int         `json:"id"        description:"自增id"`
	PicUrl    string      `json:"picUrl"    description:"轮播图地址"`
	Link      string      `json:"link"      description:"商品链接"`
	Sort      uint        `json:"sort"      description:"排序"`
	GoodsName string      `json:"goodsName" description:"商品名称"`
	GoodsId   uint        `json:"goodsId"   description:"商品id"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
