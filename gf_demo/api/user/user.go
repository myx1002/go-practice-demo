package user

import "github.com/gogf/gf/v2/net/ghttp"

type User struct {
}

func New() *User {
	return &User{}
}

func (u *User) Add(req *ghttp.Request) {
	req.Response.Write("添加用户!")
}

func (u *User) Update(req *ghttp.Request) {
	req.Response.Write("更新用户!")
}

func (u *User) Delete(req *ghttp.Request) {
	req.Response.Write("删除用户!")
}

func (u *User) Get(req *ghttp.Request) {
	req.Response.Write("查询一个用户!")
}
