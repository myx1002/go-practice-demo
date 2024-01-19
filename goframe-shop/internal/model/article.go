package model

import (
	"goframe-shop/internal/model/entity"
)

type ArticleSaveCommonInput struct {
	Title   string
	Desc    string
	PicUrl  string
	Detail  string
	IsAdmin string
	UserId  uint
}

type ArticleCreateInput struct {
	ArticleSaveCommonInput
}

type ArticleCreateOutput struct {
	ArticleId uint
}

type ArticleUpdateInput struct {
	ArticleId uint
	ArticleSaveCommonInput
}

type ArticleUpdateOutput struct {
	ArticleId uint
}

type ArticleDeleteInput struct {
	ArticleId uint
}

type ArticleDeleteOutput struct {
	ArticleId uint
}

type ArticleListInput struct {
	Page  int
	Size  int
	Title string
}

type ArticleListItem struct {
	entity.ArticleInfo
}

type ArticleListOutput struct {
	List  []ArticleListItem
	Page  int
	Size  int
	Total int
}
