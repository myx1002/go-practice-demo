package svc

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"go-zero-demo/orderrpc/orderinfo"
	"go-zero-demo/userapi/internal/config"
	"go-zero-demo/userapi/internal/middleware"
	"go-zero-demo/userapi/model"
)

// ServiceContext 初始化配置
type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
	UserModel      model.UserModel
	UserDataModel  model.UserDataModel
	OrderRpcClient orderinfo.OrderInfo // rpc客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserModel:      model.NewUserModel(sqlConn, c.Cache),
		UserDataModel:  model.NewUserDataModel(sqlConn, c.Cache),
		OrderRpcClient: orderinfo.NewOrderInfo(zrpc.MustNewClient(c.OrderRpcConf, zrpc.WithUnaryClientInterceptor(TestClientInterceptor))), // 引入rpc客户端到service上下文
	}
}

// rpc的客户端拦截器，在调用rpc接口的时候会触发
// 作用：可以在不同服务之间通过metadata传值
func TestClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("TestClientInterceptor =======> start \n")
	fmt.Printf("TestClientInterceptor-method =======> %v \n", method)
	fmt.Printf("TestClientInterceptor-req =======> %v \n", req)
	fmt.Printf("TestClientInterceptor-reply =======> %v \n", reply)
	fmt.Printf("TestClientInterceptor-cc =======> %v \n", cc)

	// 模拟metadata
	md := metadata.New(map[string]string{"md_test_name_key": "md_test_name_val"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	// handler即对应的logic的方法
	err := invoker(ctx, method, req, reply, cc, opts...)

	fmt.Printf("TestClientInterceptor =======> end \n")
	if err != nil {
		return err
	}
	return nil
}
