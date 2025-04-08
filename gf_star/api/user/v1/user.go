package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"users/register" method:"post" tags:"User" summary:"用户注册"`
	UserName string `v:"required|length:6,30#请输入用户名|用户名长度为:min到:max位" json:"username"`
	Password string `v:"required|length:6,30#请输入密码|密码长度为:min到:max位" json:"password"`
	Email    string `v:"required|email#请输入邮箱|邮箱格式错误" json:"email"`
}

type RegisterRes struct {
}
