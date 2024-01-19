package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsOptionsSaveCommon struct {
	PicUrl  string `json:"pic_url"    sm:"图片"       description:"图片"`
	Name    string `json:"name"    v:"required#商品规格名称不能为空"  sm:"商品规格名称" description:"商品规格名称"`
	Price   uint   `json:"price"     v:"min:0#商品规格价格不能为负数"     sm:"价格 单位分"   description:"价格 单位分"`
	GoodsId uint   `json:"goods_id"   v:"required#商品Id不能为空"      description:"商品id" sm:"商品id"`
	Stock   int    `json:"stock"            description:"库存" sm:"库存"`
}
type GoodsOptionsCreateReq struct {
	g.Meta `path:"/goods_options/add" method:"post" tags:"商品规格管理" sm:"新增商品规格"`
	GoodsOptionsSaveCommon
}

type GoodsOptionsCreateRes struct {
	GoodsOptionsId uint `json:"goods_options_id" sm:"商品规格Id"`
}

type GoodsOptionsUpdateReq struct {
	g.Meta         `path:"/goods_options/update" method:"post" tags:"商品规格管理" sm:"编辑商品规格"`
	GoodsOptionsId uint `json:"goods_options_id" v:"required#商品规格Id不能为空" sm:"商品规格Id"`
	GoodsOptionsSaveCommon
}

type GoodsOptionsUpdateRes struct {
	GoodsOptionsId uint `json:"goods_options_id" sm:"商品规格Id"`
}

type GoodsOptionsDeleteReq struct {
	g.Meta         `path:"/goods_options/delete" method:"delete" tags:"商品规格管理" sm:"删除商品规格"`
	GoodsOptionsId uint `json:"goods_options_id" v:"required#商品规格Id不能为空" sm:"商品规格Id"`
}

type GoodsOptionsDeleteRes struct {
	GoodsOptionsId uint `json:"goods_options_id" sm:"商品规格Id"`
}

type GoodsOptionsListReq struct {
	g.Meta `path:"/goods_options/list" method:"get" tags:"商品规格管理" sm:"商品规格列表"`
	Name   string `json:"name" sm:"商品规格名称" dc:"根据商品规格名称查找"`
	CommonPaginationReq
}

type GoodsOptionsListRes struct {
	CommonPaginationRes
}
