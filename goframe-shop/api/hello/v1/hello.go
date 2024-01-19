package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// g.Meta 元数据
// 地址：https://goframe.org/pages/viewpage.action?pageId=47703679
// 操作方法：https://goframe.org/pages/viewpage.action?pageId=7297204
type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
