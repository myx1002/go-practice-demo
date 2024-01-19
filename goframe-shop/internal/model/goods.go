package model

import (
	"goframe-shop/internal/model/do"
	"goframe-shop/internal/model/entity"
)

type GoodsSaveCommonInput struct {
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	CouponId         int
	Stock            int
	Sale             int
	Tags             string
	DetailInfo       string
}

type GoodsCreateInput struct {
	GoodsSaveCommonInput
}

type GoodsCreateOutput struct {
	GoodsId uint
}

type GoodsUpdateInput struct {
	GoodsId uint
	GoodsSaveCommonInput
}

type GoodsUpdateOutput struct {
	GoodsId uint
}

type GoodsDeleteInput struct {
	GoodsId uint
}

type GoodsDeleteOutput struct {
	GoodsId uint
}

type GoodsListInput struct {
	Page int
	Size int
	Name string
}

type GoodsListItem struct {
	entity.GoodsInfo
}

type GoodsListOutput struct {
	List  []GoodsListItem
	Page  int
	Size  int
	Total int
}

type GoodsDetailInput struct {
	GoodsId int
}

type GoodsDetailOutput struct {
	Id                 int
	PicUrl             string
	Name               string
	Price              int
	Level1CategoryId   int
	Level1CategoryName int
	Level2CategoryId   int
	Level2CategoryName int
	Level3CategoryId   int
	Level3CategoryName int
	Brand              string
	Stock              int
	Sale               int
	Tags               string
	DetailInfo         string
	CreatedAt          string
	GoodsOptions       []*do.GoodsOptionsInfo `orm:"with:goods_id=id"`
	IsCollection       bool
	PraiseCount        int
	IsPraise           bool
	CollectionCount    int
}
