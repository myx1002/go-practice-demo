package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PositionReq struct {
	g.Meta    `path:"/position/add" tags:"手工位管理" method:"post" summary:"添加手工位"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort      uint   `json:"sort"    v:"min:1#排序值必须大于0"      dc:"序号"`
	GoodsName string `json:"goods_name"    v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id"    v:"required#商品id不能为空" dc:"商品id"`
}

type PositionRes struct {
	PositionId uint `json:"position_id"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/position/delete" method:"delete" tags:"手工位管理" summary:"删除手工位"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"手工位id"`
}
type PositionDeleteRes struct{}

type PositionUpdateReq struct {
	g.Meta    `path:"/position/update/{Id}" method:"post" tags:"手工位管理" summary:"修改手工位"`
	Id        uint   `json:"id"      v:"min:1#请选择需要修改的轮播图" dc:"轮播图Id"`
	PicUrl    string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link      string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort      uint   `json:"sort"    v:"min:1#排序值必须大于0"      dc:"序号"`
	GoodsName string `json:"goods_name"    v:"required#商品名称不能为空" dc:"商品名称"`
	GoodsId   uint   `json:"goods_id"    v:"required#商品id不能为空" dc:"商品id"`
}
type PositionUpdateRes struct{}

type PositionGetListCommonReq struct {
	g.Meta `path:"/position/list" method:"get" tags:"手工位管理" summary:"手工位列表"`
	Sort   uint `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}

type PositionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
