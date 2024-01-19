package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "goframe-shop/api/hello/v1"
)

// 返回指针类型是为了降低代码对error处理的复杂度、统一项目中对空查询结果处理逻辑
func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
