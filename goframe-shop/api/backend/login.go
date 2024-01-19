package backend

import (
	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop/internal/model/entity"
)

// for jwt
//type LoginDoReq struct {
//	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
//	Name     string `json:"na" v:"required#请输入账号"   dc:"账号"`
//	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
//}
//type LoginDoRes struct {
//	Token  string    `json:"token"`
//	Expire time.Time `json:"expire"`
//}
//
//type RefreshTokenReq struct {
//	g.Meta `path:"/refresh_token" method:"post"`
//}
//
//type RefreshTokenRes struct {
//	Token  string    `json:"token"`
//	Expire time.Time `json:"expire"`
//}

type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     int                     `json:"is_admin"`
	RoleIds     string                  `json:"role_ids"`
	Permissions []entity.PermissionInfo `json:"permissions"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type LogoutRes struct {
}
