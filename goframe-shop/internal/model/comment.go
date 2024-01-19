package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

type CommentAddInput struct {
	ParentId uint
	ObjectId uint
	Type     uint8
	UserId   uint
	Content  string
}

type CommentAddOutput struct {
	CommentId uint
}

type CommentDeleteInput struct {
	CommentId uint
}

type CommentListInput struct {
	ObjectId uint
	Type     uint8
	UserId   uint
	Page     int
	Size     int
}

type CommentListItem struct {
	gmeta.Meta `orm:"table:comment_info"`
	Id         uint
	ParentId   uint
	ObjectId   uint
	Type       uint8
	UserId     uint
	Content    string
	CreatedAt  string
	UserInfo   *CommonUserItem `orm:"with:id=user_id"`
}

type CommentListOutput struct {
	List  []CommentListItem
	Page  int
	Size  int
	Total int
}
