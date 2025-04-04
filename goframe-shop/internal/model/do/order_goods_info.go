// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderGoodsInfo is the golang structure of table order_goods_info for DAO operations like Where/Data.
type OrderGoodsInfo struct {
	g.Meta         `orm:"table:order_goods_info, do:true"`
	Id             interface{} // 商品维度的订单表
	OrderId        interface{} // 关联的主订单表
	GoodsId        interface{} // 商品id
	GoodsOptionsId interface{} // 商品规格id
	Count          interface{} // 商品数量
	PayType        interface{} // 支付方式 1微信 2支付宝 3云闪付
	Remark         interface{} // 备注
	Status         interface{} // 订单状态 0待支付 1已支付 3已确认收货
	Price          interface{} // 订单金额 单位分
	CouponPrice    interface{} // 优惠券金额 单位分
	ActualPrice    interface{} // 实际支付金额 单位分
	PayAt          *gtime.Time // 支付时间
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
