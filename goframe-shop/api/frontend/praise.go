package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 点赞
type PraiseAddReq struct {
	g.Meta   `path:"/praise/add" method:"post" tags:"用户点赞" sm:"添加点赞"`
	Type     int `json:"type" v:"required#点赞类型不能为空|in:1,2#点赞类型有误" sm:"点赞类型"`
	ObjectId int `json:"object_id" v:"required#点赞对象Id不能为空" sm:"点赞对象Id"`
}

type PraiseAddRes struct {
	PraiseId int `json:"praise_id" sm:"添加id"`
}

// 取消点赞
type PraiseCancelReq struct {
	g.Meta   `path:"/praise/cancel" method:"post" tags:"用户点赞" sm:"取消点赞"`
	Type     int `json:"type" v:"required#点赞类型不能为空|in:1,2#点赞类型有误" sm:"点赞类型"`
	ObjectId int `json:"object_id" v:"required#点赞对象Id不能为空" sm:"点赞对象Id"`
}

type PraiseCancelRes struct {
}

// 点赞列表
type PraiseListReq struct {
	g.Meta `path:"/praise/list" method:"get" tags:"用户点赞" sm:"点赞列表"`
	Type   int `json:"type" v:"required#点赞类型不能为空|in:0,1,2#点赞类型有误" sm:"点赞类型"`
	CommonPaginationReq
}

type PraiseListRes struct {
	CommonPaginationRes
}
