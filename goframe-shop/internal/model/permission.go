package model

type PermissionCreateInput struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type PermissionCreateOutput struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionUpdateInput struct {
	PermissionId uint   `json:"permission_id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
}

type PermissionUpdateOutput struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionDeleteInput struct {
	PermissionId uint `json:"permission_id"`
}

type PermissionListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}

type PermissionListOutputItem struct {
	Id   uint   `json:"permission_id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type PermissionListOutput struct {
	List  []PermissionListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}
