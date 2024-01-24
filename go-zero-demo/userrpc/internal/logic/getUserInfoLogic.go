package logic

import (
	"context"

	"go-zero-demo/userrpc/internal/svc"
	"go-zero-demo/userrpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义一个 UserInfo 一元 rpc 方法，请求体和响应体必填。
func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	users := map[int64]string{
		1: "小明 form rpc",
		2: "小李 form rpc",
		3: "小贝 form rpc",
	}

	userName := "unknown form rpc"
	if name, ok := users[in.Id]; ok {
		userName = name
	}
	return &pb.GetUserInfoResp{Id: in.Id, Name: userName}, nil
}
