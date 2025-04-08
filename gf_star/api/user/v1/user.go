package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"users/register" method:"post" tags:"User" sm:"注册" tags:"用户"`
	UserName string `v:"required|length:6,30#请输入用户名|用户名长度为:min到:max位" json:"username" dc:"用户名"`
	Password string `v:"required|length:6,30#请输入密码|密码长度为:min到:max位" json:"password" dc:"密码"`
	Email    string `v:"required|email#请输入邮箱|邮箱格式错误" json:"email" dc:"邮箱"`
}

type RegisterRes struct {
}
