package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionCreateUpdateBase 创建/修改内容基类
type PositionCreateUpdateBase struct {
	PicUrl    string
	Link      string
	Sort      uint
	GoodsName string
	GoodsId   uint
}

// PositionCreateInput 创建内容
type PositionCreateInput struct {
	PositionCreateUpdateBase
}

// PositionCreateOutput 创建内容返回结果
type PositionCreateOutput struct {
	PositionId uint `json:"position_id"`
}

// PositionUpdateInput 修改内容
type PositionUpdateInput struct {
	PositionCreateUpdateBase
	Id uint
}

// PositionGetListInput 获取内容列表
type PositionGetListInput struct {
	Page int  // 分页号码
	Size int  // 分页数量，最大50
	Sort uint // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type PositionGetListOutputItem struct {
	Id        uint        `json:"id"` // 自增ID
	PicUrl    string      `json:"pic_url"`
	Link      string      `json:"link"`
	Sort      uint        `json:"sort"`
	GoodsName string      `json:"goods_name"`
	GoodsId   uint        `json:"goods_id"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

// PositionGetListOutput 查询列表结果
type PositionGetListOutput struct {
	List  []PositionGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}
