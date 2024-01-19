package model

type PlaceAnOrderInput struct {
	Number           string
	UserId           int
	Remark           string
	ConsigneeName    string
	ConsigneePhone   string
	ConsigneeAddress string
	Price            int
	CouponPrice      int
	ActualPrice      int
	OrderGoodsInfo   []*OrderGoodsInfoItem
}

type OrderGoodsInfoItem struct {
	OrderId        int
	Remark         string
	GoodsId        int
	GoodsOptionsId int
	Count          int
	Price          int
	CouponPrice    int
	ActualPrice    int
}

type PlaceAnOrderOutput struct {
	OrderId int
}
