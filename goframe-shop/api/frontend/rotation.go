package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 列表请求
type RotationGetListCommonReq struct {
	g.Meta `path:"/frontend/rotation/list" method:"get" tags:"轮播图" sm:"轮播图"`
	Sort   uint `json:"sort" in:"query" sm:"排序"` // in是数据输入方式，对应query、body或者header
	CommonPaginationReq
}

// 列表返回
type RotationGetListCommonRes struct {
	List  interface{} `json:"list"  ds:"列表"`
	Page  int         `json:"page"  ds:"分页码"`
	Size  int         `json:"size"  ds:"分页数量"`
	Total int         `json:"total" ds:"数据总数"`
}
