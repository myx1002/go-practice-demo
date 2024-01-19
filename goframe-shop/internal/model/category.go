package model

type CategoryCreateInput struct {
	ParentId uint
	PicUrl   string
	Name     string
	Level    uint
	Sort     uint
}

type CategoryCreateOutput struct {
	CategoryId uint
}

type CategoryUpdateInput struct {
	CategoryId uint
	ParentId   uint
	PicUrl     string
	Name       string
	Level      uint
	Sort       uint
}

type CategoryUpdateOutput struct {
	CategoryId uint
}

type CategoryDeleteInput struct {
	CategoryId uint
}

type CategoryListItem struct {
	Id       uint   `json:"category_id"`
	ParentId uint   `json:"parent_id"`
	PicUrl   string `json:"pic_url"`
	Name     string `json:"name"`
	Level    uint   `json:"level"`
	Sort     uint   `json:"sort"`
}

type CategoryListInput struct {
	Page int  // 分页号码
	Size int  // 分页数量，最大50
	Name uint // 分类名称
}

type CategoryListOutput struct {
	List  []CategoryListItem
	Size  int // 分页数量，最大50
	Page  int // 分页号码
	Total int // 总数
}
