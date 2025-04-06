package user

import (
	"context"
	"gf_demo/api/user"
	"github.com/gogf/gf/v2/net/ghttp"
)

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

func (u *User) Get(ctx context.Context, req *user.UserDetailReq) (res *user.UserDetailRes, err error) {
	res = &user.UserDetailRes{
		Age:  18,
		Name: "张三",
	}
	return
}
