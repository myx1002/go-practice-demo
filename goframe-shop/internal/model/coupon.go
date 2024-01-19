package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

type CouponCreateInput struct {
	Name       string
	Price      uint
	GoodsIds   string
	CategoryId uint
}

type CouponCreateOutput struct {
	CouponId uint
}

type CouponUpdateInput struct {
	CouponId   uint
	Name       string
	Price      uint
	GoodsIds   string
	CategoryId uint
}

type CouponUpdateOutput struct {
	CouponId uint
}

type CouponDeleteInput struct {
	CouponId uint
}

type CouponDeleteOutput struct {
	CouponId uint
}

type CouponListInput struct {
	Page int
	Size int
	Name string
}

type CouponListItem struct {
	gmeta.Meta `orm:"table:coupon_info"`
	Id         uint   `json:"coupon_id" v:"required#优惠卷Id不能为空" sm:"优惠卷Id"`
	Name       string `json:"name" v:"required#优惠卷名称不能为空" sm:"优惠卷名称"`
	Price      uint   `json:"price" v:"required#优惠卷金额不能为空" sm:"优惠卷金额"`
	GoodsIds   string `json:"goods_ids"  sm:"优惠卷关联商品"`
	CategoryId uint   `json:"category_id"  sm:"优惠卷关联商品分类"`
	CreatedAt  string `json:"created_at"  sm:"优惠卷关联商品分类"`
}

type CouponListOutput struct {
	List  []CouponListItem
	Page  int
	Size  int
	Total int
}
