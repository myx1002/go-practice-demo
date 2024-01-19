package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CouponCreateReq struct {
	g.Meta     `path:"/coupon/add" method:"post" tags:"优惠卷管理" sm:"新增优惠卷"`
	Name       string `json:"name" v:"required#优惠卷名称不能为空" sm:"优惠卷名称"`
	Price      uint   `json:"price" v:"required#优惠卷金额不能为空" sm:"优惠卷金额"`
	GoodsIds   string `json:"goods_ids"  sm:"优惠卷关联商品"`
	CategoryId uint   `json:"category_id"  sm:"优惠卷关联商品分类"`
}

type CouponCreateRes struct {
	CouponId uint `json:"coupon_id" sm:"优惠卷Id"`
}

type CouponUpdateReq struct {
	g.Meta     `path:"/coupon/update/{coupon_id}" method:"post" tags:"优惠卷管理" sm:"编辑优惠卷"`
	CouponId   uint   `json:"coupon_id" v:"required#优惠卷Id不能为空" sm:"优惠卷Id"`
	Name       string `json:"name" v:"required#优惠卷名称不能为空" sm:"优惠卷名称"`
	Price      uint   `json:"price" v:"required#优惠卷金额不能为空" sm:"优惠卷金额"`
	GoodsIds   string `json:"goods_ids"  sm:"优惠卷关联商品"`
	CategoryId uint   `json:"category_id"  sm:"优惠卷关联商品分类"`
}

type CouponUpdateRes struct {
	CouponId uint `json:"coupon_id" sm:"优惠卷Id"`
}

type CouponDeleteReq struct {
	g.Meta   `path:"/coupon/delete" method:"delete" tags:"优惠卷管理" sm:"删除优惠卷"`
	CouponId uint `json:"coupon_id" v:"required#优惠卷Id不能为空" sm:"优惠卷Id"`
}

type CouponDeleteRes struct {
	CouponId uint `json:"coupon_id" sm:"优惠卷Id"`
}

type CouponListReq struct {
	g.Meta `path:"/coupon/list" method:"get" tags:"优惠卷管理" sm:"优惠卷列表"`
	Name   string `json:"name" sm:"优惠卷名称"`
	CommonPaginationReq
}

type CouponListRes struct {
	CommonPaginationRes
}
