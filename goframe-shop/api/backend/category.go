package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CategoryCreateReq struct {
	g.Meta   `path:"/category/add" method:"post" tags:"商品分类" dc:"添加分类"`
	ParentId uint   `json:"parent_id"  dc:"父级分类id"`
	PicUrl   string `json:"pic_url"   dc:"分类icon"`
	Name     string `json:"name"  v:"required#分类名称不能为空" dc:"分类名称"`
	Level    uint   `json:"level" v:"min:1#层架最小为1|max:3#层级最大为3" dc:"分类层级"`
	Sort     uint   `json:"sort" dc:"分类序号"`
}

type CategoryCreateRes struct {
	CategoryId uint `json:"category_id"`
}

type CategoryUpdateReq struct {
	g.Meta     `path:"/category/update/{category_id}" method:"post" tags:"商品分类" dc:"编辑分类"`
	CategoryId uint   `json:"category_id" v:"required#分类id不能为空"  dc:"分类id"`
	ParentId   uint   `json:"parent_id"  dc:"父级分类id"`
	PicUrl     string `json:"pic_url"   dc:"分类icon"`
	Name       string `json:"name"  v:"required#分类名称不能为空" dc:"分类名称"`
	Level      uint   `json:"level" v:"min:1#层架最小为1|max:3#层级最大为3" dc:"分类层级"`
	Sort       uint   `json:"sort" dc:"分类序号"`
}

type CategoryUpdateRes struct {
	CategoryId uint `json:"category_id"`
}

type CategoryDeleteReq struct {
	g.Meta     `path:"/category/delete" method:"delete" tags:"商品分类" dc:"删除分类"`
	CategoryId uint `json:"category_id" v:"required#分类id不能为空"  dc:"分类id"`
}

type CategoryDeleteRes struct {
	CategoryId uint `json:"category_id"`
}

type CategoryListReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"商品分类" dc:"分类列表"`
	Name   string `json:"name"   dc:"分类名称"`
	CommonPaginationReq
}

type CategoryListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Total int         `json:"total" description:"数据总数"`
}
