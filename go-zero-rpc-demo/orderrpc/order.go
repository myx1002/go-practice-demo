package main

import (
	"context"
	"flag"
	"fmt"

	"go-zero-rpc-demo/orderrpc/internal/config"
	"go-zero-rpc-demo/orderrpc/internal/server"
	"go-zero-rpc-demo/orderrpc/internal/svc"
	"go-zero-rpc-demo/orderrpc/order_pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order_pb.RegisterOrderInfoServer(grpcServer, server.NewOrderInfoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 添加服务端拦截器
	s.AddUnaryInterceptors(TestServerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func TestServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Printf("TestServerInterceptor =======> start \n")
	fmt.Printf("TestServerInterceptor-info =======> %v \n", info)
	fmt.Printf("TestServerInterceptor-req =======> %v \n", req)

	// handler即对应的logic的方法
	res, err := handler(ctx, req)

	fmt.Printf("TestServerInterceptor =======> end \n")
	if err != nil {
		return nil, err
	}

	return res, err
}
