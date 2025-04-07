// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct {
	g.Meta   `path:"/hello/{version}/params" method:"all" tags:"Hello" summary:"Hello Params"`
	UserName string `f:"user_name" d:"李四"`
	Age      int
	Password string
	Version  string
}

type ResponsReq struct {
	g.Meta   `path:"/hello/{version}/respons" method:"all" tags:"Hello" summary:"Hello Params"`
	UserName string `f:"user_name" d:"李四"`
	Age      int
	Password string
	Version  string
}

type ParamsRes struct {
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
	Version  string `json:"version"`
}

type ValidateReq struct {
	g.Meta `method:"all"`
	Name   string `v:"required|length:6,16#必填|长度为6-16"`
	Age    int    `v:"required"`
	Phone  string `v:"phone#手机号码格式有误"`
}

type ValidateRes struct {
	Data interface{} `json:"data"`
}
