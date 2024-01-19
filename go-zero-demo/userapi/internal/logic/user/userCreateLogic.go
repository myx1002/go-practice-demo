package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/userapi/internal/svc"
	"go-zero-demo/userapi/internal/types"
	"go-zero-demo/userapi/model"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	if err := l.svcCtx.UserModel.TransCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		user := &model.User{}
		user.Name = req.Name
		userResult, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
		if err != nil {
			return err
		}
		userId, _ := userResult.LastInsertId()

		fmt.Println(userId)
		userData := &model.UserData{
			UserId: userId,
			Data:   "xxxxxxxx",
		}

		_, err = l.svcCtx.UserDataModel.TransInsert(ctx, session, userData)

		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return &types.UserCreateResp{Flag: true}, nil
}
