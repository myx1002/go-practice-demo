// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package user

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"proxima/app/gateway/api/user"
	"proxima/app/gateway/utility"
	v1 "proxima/app/user/api/account/v1"
)

// 业务网关相当于gRPC客户端，各个微服务相当于gRPC服务端。我们将在控制器属性里，定义gRPC client，供后续使用。
// 定义client由grpcx.Client.MustNewGrpcClientConn(service, opts...)完成。
type ControllerV1 struct {
	AccountClient v1.AccountClient
}

func NewV1() user.IUserV1 {
	var conn = grpcx.Client.MustNewGrpcClientConn("user", grpcx.Client.ChainUnary(
		utility.GrpcClientTimeout,
	))
	return &ControllerV1{
		AccountClient: v1.NewAccountClient(conn),
	}
}
