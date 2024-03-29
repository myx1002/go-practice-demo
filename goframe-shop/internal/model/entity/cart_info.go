// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CartInfo is the golang structure for table cart_info.
type CartInfo struct {
	Id        int         `json:"id"        description:"购物车表"`
	UserId    int         `json:"userId"    description:""`
	GoodsId   int         `json:"goodsId"   description:""`
	Count     int         `json:"count"     description:"商品数量"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
