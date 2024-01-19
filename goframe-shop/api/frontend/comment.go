package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CommentAddReq struct {
	g.Meta   `path:"/comment/add" method:"post"  tags:"评论管理" sm:"新增评论"`
	ParentId uint   `json:"parent_id" default:"0" sm:"父级评论id"`
	ObjectId uint   `json:"object_id" v:"required#评论对象Id不能为空" sm:"对象Id"`
	Type     int    `json:"type" v:"required#评论类型不能为空|in:1,2#点赞类型有误" sm:"评论类型"`
	Content  string `json:"content" v:"required#评论内容不能为空" sm:"评论内容"`
}

type CommentAddRes struct {
	CommentId uint `json:"comment_id"`
}

type CommentDeleteReq struct {
	g.Meta    `path:"/comment/delete" method:"delete"  tags:"评论管理" sm:"删除评论"`
	CommentId uint `json:"comment_id"`
}

type CommentDeleteRes struct {
}

// 评论列表
type CommentListReq struct {
	g.Meta   `path:"/comment/list" method:"get" tags:"评论管理" sm:"评论列表"`
	ObjectId uint `json:"object_id" v:"required#评论对象Id不能为空" sm:"对象Id"`
	Type     int  `json:"type" v:"required#评论类型不能为空|in:1,2#点赞类型有误" sm:"评论类型"`
	CommonPaginationReq
}

type CommentListRes struct {
	CommonPaginationRes
}
