// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	Id             int         `json:"id"             description:""`
	Name           string      `json:"name"           description:"用户名"`
	Avatar         string      `json:"avatar"         description:"头像"`
	Password       string      `json:"password"       description:""`
	UserSalt       string      `json:"userSalt"       description:"加密盐 生成密码用"`
	Sex            int         `json:"sex"            description:"1男 2女"`
	Status         int         `json:"status"         description:"1正常 2拉黑冻结"`
	Sign           string      `json:"sign"           description:"个性签名"`
	SecretQuestion string      `json:"secretQuestion" description:"密保问题"`
	SecretAnswer   string      `json:"secretAnswer"   description:"密保问题的答案"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:""`
}
