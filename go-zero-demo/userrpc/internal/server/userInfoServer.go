// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero-demo/userrpc/internal/logic"
	"go-zero-demo/userrpc/internal/svc"
	"go-zero-demo/userrpc/pb"
)

type UserInfoServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserInfoServer
}

func NewUserInfoServer(svcCtx *svc.ServiceContext) *UserInfoServer {
	return &UserInfoServer{
		svcCtx: svcCtx,
	}
}

// 定义一个 UserInfo 一元 rpc 方法，请求体和响应体必填。
func (s *UserInfoServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}