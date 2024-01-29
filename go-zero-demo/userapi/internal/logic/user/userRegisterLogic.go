package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"go-zero-demo/userapi/internal/svc"
	"go-zero-demo/userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// todo: add your logic here and delete this line
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, req.Phone)

	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &types.UserRegisterResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func getJwtToken(secretKey string, iat, seconds int64, userPhone string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_phone"] = userPhone
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
