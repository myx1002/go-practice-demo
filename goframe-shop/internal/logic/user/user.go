package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) RegisterUser(ctx context.Context, in model.RegisterUserInput) (out *model.RegisterUserOutput, err error) {
	// 检查该用户名是否已存在
	one, err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Name, in.Name).One()
	if err != nil {
		return nil, err
	}

	if one != nil {
		return nil, gerror.New("该用户名已存在，请重新输入")
	}

	// 生成盐值
	UserSalt := grand.S(10)
	in.Password = utility.EncrypPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt

	userId, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.RegisterUserOutput{UserId: uint(userId)}, err
}

func (s *sUser) ChangeUserPassword(ctx context.Context, in model.ChangeUserPasswordInput) (out model.ChangeUserPasswordOutput, err error) {
	userInfo := entity.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return model.ChangeUserPasswordOutput{}, err
	}

	if in.SecretAnswer != gconv.String(userInfo.SecretAnswer) {
		return model.ChangeUserPasswordOutput{}, gerror.New("密保答案不正确，请重新输入！")
	}

	UserSalt := grand.S(10)
	in.Password = utility.EncrypPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	if err != nil {
		return model.ChangeUserPasswordOutput{}, err
	}

	return model.ChangeUserPasswordOutput{UserId: userId}, nil
}
