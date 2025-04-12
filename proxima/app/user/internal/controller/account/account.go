package account

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "proxima/app/user/api/account/v1"
	"proxima/app/user/api/pbentity"
	"proxima/app/user/internal/logic/account"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedAccountServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAccountServer(s.Server, &Controller{})
}

func (*Controller) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	uid, err := account.Register(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserRegisterRes{
		Id: uint32(uid),
	}, nil
}

func (*Controller) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	token, err := account.Login(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserLoginRes{
		Token: token,
	}, nil
}

func (*Controller) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	info, err := account.Info(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoRes{
		User: &pbentity.Users{
			Id:        uint32(info.Id),
			Username:  info.Username,
			Email:     info.Email,
			CreatedAt: timestamppb.New(info.CreatedAt.Time),
			UpdatedAt: timestamppb.New(info.UpdatedAt.Time),
		},
	}, nil
}
