package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/rotation/add" tags:"轮播图" method:"post" summary:"添加轮播图"`
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   uint   `json:"sort"    v:"min:1#排序值必须大于0"      dc:"序号"`
}
type RotationRes struct {
	RotationId uint `json:"rotation_id"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"轮播图id"`
}
type RotationDeleteRes struct{}

type RotationUpdateReq struct {
	g.Meta `path:"/rotation/update/{Id}" method:"post" tags:"轮播图" summary:"修改轮播图"`
	Id     uint   `json:"id"      v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   uint   `json:"sort"    v:"min:1#排序值必须大于0"      dc:"序号"`
}
type RotationUpdateRes struct{}

type RotationGetListCommonReq struct {
	g.Meta `path:"/rotation/list" method:"get" tags:"轮播图" summary:"轮播图列表"`
	Sort   uint `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type RotationGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
