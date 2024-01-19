package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsSaveCommon struct {
	PicUrl           string `json:"pic_url"    sm:"图片"       description:"图片"`
	Name             string `json:"name"    v:"required#商品名称不能为空"  sm:"商品名称" description:"商品名称"`
	Price            int    `json:"price"     v:"min:0#商品价格不能为负数"     sm:"价格 单位分"   description:"价格 单位分"`
	Level1CategoryId int    `json:"level_1_category_id" description:"1级分类id" sm:"1级分类id"`
	Level2CategoryId int    `json:"level_2_category_id" description:"2级分类id" sm:"2级分类id"`
	Level3CategoryId int    `json:"level_3_category_id" description:"3级分类id" sm:"3级分类id"`
	Brand            string `json:"brand" v:"max-length:30#品牌名称不能超过30个字符"            description:"品牌"  sm:"品牌"`
	CouponId         int    `json:"coupon_id"         description:"优惠券id" sm:"优惠券id"`
	Stock            int    `json:"stock"            description:"库存" sm:"库存"`
	Sale             int    `json:"sale"             description:"销量" sm:"销量"`
	Tags             string `json:"tags"             description:"标签"  sm:"标签"`
	DetailInfo       string `json:"detail_info"       description:"商品详情" sm:"商品详情"`
}
type GoodsCreateReq struct {
	g.Meta `path:"/goods/add" method:"post" tags:"商品管理" sm:"新增商品"`
	GoodsSaveCommon
}

type GoodsCreateRes struct {
	GoodsId uint `json:"goods_id" sm:"优惠卷Id"`
}

type GoodsUpdateReq struct {
	g.Meta  `path:"/goods/update/{goods_id}" method:"post" tags:"商品管理" sm:"编辑商品"`
	GoodsId uint `json:"goods_id" v:"required#商品Id不能为空" sm:"商品Id"`
	GoodsSaveCommon
}

type GoodsUpdateRes struct {
	GoodsId uint `json:"goods_id" sm:"商品Id"`
}

type GoodsDeleteReq struct {
	g.Meta  `path:"/goods/delete" method:"delete" tags:"商品管理" sm:"删除商品"`
	GoodsId uint `json:"goods_id" v:"required#商品Id不能为空" sm:"商品Id"`
}

type GoodsDeleteRes struct {
	GoodsId uint `json:"goods_id" sm:"优惠卷Id"`
}

type GoodsListReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"商品管理" sm:"商品列表"`
	Name   string `json:"name" sm:"商品名称" dc:"根据商品名称查找"`
	CommonPaginationReq
}

type GoodsListRes struct {
	CommonPaginationRes
}
