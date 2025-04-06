package user

import "github.com/gogf/gf/v2/frame/g"

type UserDetailReq struct {
	g.Meta `path:"/user/detail" method:"get" tags:"User" summary:"获取用户详情"`
	Id     int `json:"id"`
}

type UserDetailRes struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
