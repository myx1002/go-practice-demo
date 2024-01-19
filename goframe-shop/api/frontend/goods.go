package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsListReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"前台商品" sm:"商品列表"`
	Name   string `json:"name" sm:"商品名称" dc:"根据商品名称查找"`
	CommonPaginationReq
}

type GoodsListRes struct {
	CommonPaginationRes
}

type GoodsDetailReq struct {
	g.Meta  `path:"/goods/detail/{goods_id}" method:"get" tags:"前台商品" sm:"商品详情"`
	GoodsId int `json:"goods_id"               description:""`
}

type GoodsDetailRes struct {
	Id                 int                    `json:"goods_id"               description:""`
	PicUrl             string                 `json:"pic_url"           description:"图片"`
	Name               string                 `json:"name"             description:"商品名称"`
	Price              int                    `json:"price"            description:"价格 单位分"`
	Level1CategoryId   int                    `json:"level1_category_id" description:"1级分类id"`
	Level1CategoryName int                    `json:"level1_category_name" description:"1级分类名称"`
	Level2CategoryId   int                    `json:"level2_category_id" description:"2级分类id"`
	Level2CategoryName int                    `json:"level2_category_name" description:"2级分类名称"`
	Level3CategoryId   int                    `json:"level3_category_id" description:"3级分类id"`
	Level3CategoryName int                    `json:"level3_category_name" description:"3级分类名称"`
	Brand              string                 `json:"brand"            description:"品牌"`
	Stock              int                    `json:"stock"            description:"库存"`
	Sale               int                    `json:"sale"             description:"销量"`
	Tags               string                 `json:"tags"             description:"标签"`
	DetailInfo         string                 `json:"detail_info"       description:"商品详情"`
	CreatedAt          string                 `json:"created_at"        description:""`
	GoodsOptions       []GoodsOptionsBaseItem `json:"goods_options" description:"规格信息"`
	IsCollection       bool                   `json:"is_collection" description:"是否收藏"`
	IsPraise           bool                   `json:"is_praise" description:"是否点赞"`
	PraiseCount        int                    `json:"praise_count" description:"点赞数量"`
	CollectionCount    int                    `json:"collection_count" description:"收藏数量"`
}

type GoodsOptionsBaseItem struct {
	Id      int    `json:"goods_options_id"               description:""`
	GoodsId int    `json:"goods_id"               description:""`
	PicUrl  string `json:"pic_url"        description:""`
	Name    string `json:"name"        description:""`
	Stock   int    `json:"stock"        description:""`
}
