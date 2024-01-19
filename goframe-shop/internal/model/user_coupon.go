package model

type UserCouponCreateInput struct {
	UserId   uint
	CouponId uint
}

type UserCouponCreateOutput struct {
	UserCouponId uint
}

type UserCouponUpdateInput struct {
	UserCouponId uint
	UserId       uint
	CouponId     uint
	Status       uint
}

type UserCouponUpdateOutput struct {
	UserCouponId uint
}

type UserCouponDeleteInput struct {
	UserCouponId uint
}

type UserCouponDeleteOutput struct {
	UserCouponId uint
}

type UserCouponListInput struct {
	Page       int
	Size       int
	UserName   string
	CouponName string
}

type UserCouponListItem struct {
	Id         uint            `json:"user_coupon_id"`
	UserId     uint            `json:"user_id"`
	CouponId   uint            `json:"coupon_id"`
	Status     uint            `json:"status"`
	CouponInfo *CouponListItem `json:"coupon_info" orm:"with:id=coupon_id"`
	// todo 关联用户信息
}

type UserCouponListOutput struct {
	List  []UserCouponListItem
	Page  int
	Size  int
	Total int
}
