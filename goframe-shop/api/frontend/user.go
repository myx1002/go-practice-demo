package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterUserReq struct {
	g.Meta         `path:"/register" method:"post"  tags:"用户管理" description:"用户注册"`
	Name           string `json:"name"    v:"required#用户名不能为空|max-length:10#用户名不能超过10个字符"     description:"用户名"`
	Avatar         string `json:"avatar"       description:"头像"`
	Password       string `json:"password"   v:"password"   description:""`
	Sex            int    `json:"sex"   default:"1"      description:"1男 2女"`
	Sign           string `json:"sign"         description:"个性签名"`
	SecretQuestion string `json:"secret_question" v:"max-length:32#密保问题不能超过32个字符"  description:"密保问题"`
	SecretAnswer   string `json:"secret_answer" v:"max-length:32#密保答案不能超过32个字符" description:"密保答案"`
}

type RegisterUserRes struct {
	UserId uint `json:"user_id" ds:"用户Id"`
}

type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户管理" description:"用户信息"`
}

type UserInfoRes struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar" `
	Status int    `json:"status"`
	Sign   string `json:"sign" `
}

type ChangeUserPasswordReq struct {
	g.Meta       `path:"/user/change_password" method:"post" tags:"用户管理" description:"修改密码"`
	Password     string `json:"password"   v:"password"   description:""`
	SecretAnswer string `json:"secret_answer" v:"max-length:32#密保答案不能超过32个字符" description:"密保答案"`
}

type ChangeUserPasswordRes struct {
	UserId uint `json:"user_id"`
}
