// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	Id               int         `json:"id"               description:""`
	Number           string      `json:"number"           description:"订单编号"`
	UserId           int         `json:"userId"           description:"用户id"`
	PayType          int         `json:"payType"          description:"支付方式 1微信 2支付宝 3云闪付"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:""`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:""`
	Remark           string      `json:"remark"           description:"备注"`
	PayAt            *gtime.Time `json:"payAt"            description:"支付时间"`
	Status           int         `json:"status"           description:"订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价"`
	ConsigneeName    string      `json:"consigneeName"    description:"收货人姓名"`
	ConsigneePhone   string      `json:"consigneePhone"   description:"收货人手机号"`
	ConsigneeAddress string      `json:"consigneeAddress" description:"收货人详细地址"`
	Price            int         `json:"price"            description:"订单金额 单位分"`
	CouponPrice      int         `json:"couponPrice"      description:"优惠券金额 单位分"`
	ActualPrice      int         `json:"actualPrice"      description:"实际支付金额 单位分"`
}
