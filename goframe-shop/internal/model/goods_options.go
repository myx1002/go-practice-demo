package model

import (
	"goframe-shop/internal/model/entity"
)

type GoodsOptionsSaveCommonInput struct {
	PicUrl  string
	Name    string
	Price   int
	Stock   int
	GoodsId uint
}

type GoodsOptionsCreateInput struct {
	GoodsOptionsSaveCommonInput
}

type GoodsOptionsCreateOutput struct {
	GoodsOptionsId uint
}

type GoodsOptionsUpdateInput struct {
	GoodsOptionsId uint
	GoodsOptionsSaveCommonInput
}

type GoodsOptionsUpdateOutput struct {
	GoodsOptionsId uint
}

type GoodsOptionsDeleteInput struct {
	GoodsOptionsId uint
}

type GoodsOptionsDeleteOutput struct {
	GoodsOptionsId uint
}

type GoodsOptionsListInput struct {
	Page int
	Size int
	Name string
}

type GoodsOptionsListItem struct {
	entity.GoodsOptionsInfo
}

type GoodsOptionsListOutput struct {
	List  []GoodsOptionsListItem
	Page  int
	Size  int
	Total int
}
