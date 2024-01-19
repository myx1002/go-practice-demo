// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure of table position_info for DAO operations like Where/Data.
type PositionInfo struct {
	g.Meta    `orm:"table:position_info, do:true"`
	Id        interface{} // 自增id
	PicUrl    interface{} // 轮播图地址
	Link      interface{} // 商品链接
	Sort      interface{} // 排序
	GoodsName interface{} // 商品名称
	GoodsId   interface{} // 商品id
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
