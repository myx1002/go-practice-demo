package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PlaceAnOrderReq struct {
	g.Meta           `path:"/order/save" method:"post" tags:"订单管理" sm:"用户下单"`
	Remark           string                   `json:"remark"     v:"max-length:50#备注信息不能超过50个字符"      description:"备注"`
	ConsigneeName    string                   `json:"consignee_name"  v:"required#收货人姓名不能为空|max-length:10#收货人姓名不能超过10个字符"  description:"收货人姓名"`
	ConsigneePhone   string                   `json:"consignee_phone"  v:"required#收货人手机号不能为空|phone#收货人手机号格式有误" description:"收货人手机号"`
	ConsigneeAddress string                   `json:"consignee_address" v:"required#收货人详细地址不能为空|max-length:50#收货人详细地址不能超过50个字符" description:"收货人详细地址"`
	Price            int                      `json:"price"        v:"required#订单金额不能为空"    description:"订单金额 单位分"`
	CouponPrice      int                      `json:"coupon_price"  v:"required#优惠券金额不能为空"    description:"优惠券金额 单位分"`
	ActualPrice      int                      `json:"actual_price"  v:"required#实际支付金额不能为空"    description:"实际支付金额 单位分"`
	OrderGoodsInfo   []*OrderGoodsInfoReqItem `json:"order_goods_info"`
}

type OrderGoodsInfoReqItem struct {
	Remark         string `json:"remark"     v:"max-length:50#备注信息不能超过50个字符"      description:"备注"`
	GoodsId        int    `json:"goods_id"        v:"required#商品ID不能为空"    description:"商品ID"`
	GoodsOptionsId int    `json:"goods_options_id"        v:"required#商品规格ID不能为空"    description:"商品规格ID"`
	Count          int    `json:"count"        v:"required#数量不能为空|min:1#数量不能少于1件"    description:"数量"`
	Price          int    `json:"price"        v:"required#订单金额不能为空"    description:"订单金额 单位分"`
	CouponPrice    int    `json:"coupon_price"  v:"required#优惠券金额不能为空"    description:"优惠券金额 单位分"`
	ActualPrice    int    `json:"actual_price"  v:"required#实际支付金额不能为空"    description:"实际支付金额 单位分"`
}

type PlaceAnOrderRes struct {
	OrderId uint `json:"order_id"`
}
