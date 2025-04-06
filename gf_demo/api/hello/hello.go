// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct {
	g.Meta   `path:"/hello/{version}/params" method:"all" tags:"Hello" summary:"Hello Params"`
	UserName string `f:"user_name" json:"user_name" d:"李四"`
	Age      int    `json:"age"`
	Password string `json:"password"`
	Version  string `json:"version"`
}

type ParamsRes struct {
}
