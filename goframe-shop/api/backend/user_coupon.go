package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserCouponCreateReq struct {
	g.Meta   `path:"/user/coupon/add" method:"post" tags:"用户优惠卷管理" sm:"新增用户优惠卷"`
	UserId   uint `json:"user_id" v:"required#用户Id不能为空" sm:"用户Id"`
	CouponId uint `json:"coupon_id" v:"required#优惠卷Id不能为空" sm:"优惠卷Id"`
}

type UserCouponCreateRes struct {
	UserCouponId uint `json:"user_coupon_id" sm:"优惠卷Id"`
}

type UserCouponUpdateReq struct {
	g.Meta       `path:"/user/coupon/update/{user_coupon_id}" method:"post" tags:"用户优惠卷管理" sm:"编辑用户优惠卷"`
	UserCouponId uint `json:"user_coupon_id" v:"required#用户优惠卷Id不能为空" sm:"用户优惠卷Id"`
	UserId       uint `json:"user_id" v:"required#用户Id不能为空" sm:"用户Id"`
	CouponId     uint `json:"coupon_id" v:"required#优惠卷Id不能为空" sm:"优惠卷Id"`
	Status       uint `json:"status" v:"required#优惠卷状态不能为空|in:1,2,3#优惠卷状态有误" sm:"优惠卷状态"`
}

type UserCouponUpdateRes struct {
	UserCouponId uint `json:"user_coupon_id" sm:"用户优惠卷Id"`
}

type UserCouponDeleteReq struct {
	g.Meta       `path:"/user/coupon/delete" method:"delete" tags:"用户优惠卷管理" sm:"删除用户优惠卷"`
	UserCouponId uint `json:"user_coupon_id" v:"required#用户优惠卷Id不能为空" sm:"用户优惠卷Id"`
}

type UserCouponDeleteRes struct {
	UserCouponId uint `json:"user_coupon_id" sm:"优惠卷Id"`
}

type UserCouponListReq struct {
	g.Meta     `path:"/user/coupon/list" method:"get" tags:"用户优惠卷管理" sm:"用户优惠卷列表"`
	UserName   string `json:"user_name" sm:"用户名称" dc:"根据用户名称查找优惠卷"`
	CouponName string `json:"coupon_name" sm:"优惠卷名称" dc:"根据优惠卷名称查找优惠卷"`
	CommonPaginationReq
}

type UserCouponListRes struct {
	CommonPaginationRes
}
