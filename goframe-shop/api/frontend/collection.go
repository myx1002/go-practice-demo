package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CollectionAddReq struct {
	g.Meta   `path:"/collection/add" method:"post" tags:"用户收藏" sm:"添加收藏"`
	Type     int `json:"type" v:"required#收藏类型不能为空|in:1,2#收藏类型有误" sm:"收藏类型"`
	ObjectId int `json:"object_id" v:"required#收藏对象Id不能为空" sm:"收藏对象Id"`
}

type CollectionAddRes struct {
	CollectionId int `json:"collection_id" sm:"收藏id"`
}

type CollectionCancelReq struct {
	g.Meta   `path:"/collection/cancel" method:"post" tags:"用户收藏" sm:"取消收藏"`
	Type     int `json:"type" v:"required#收藏类型不能为空|in:1,2#收藏类型有误" sm:"收藏类型"`
	ObjectId int `json:"object_id" v:"required#收藏对象Id不能为空" sm:"收藏对象Id"`
}

type CollectionCancelRes struct {
}

type CollectionListReq struct {
	g.Meta `path:"/collection/list" method:"get" tags:"用户收藏" sm:"收藏列表"`
	CommonPaginationReq
	Type int `json:"type" v:"required#收藏类型不能为空|in:0,1,2#收藏类型有误" sm:"收藏类型"`
}

type CollectionListRes struct {
	CommonPaginationRes
}
