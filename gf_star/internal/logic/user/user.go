package user

import (
	"context"
	"gf_star/internal/dao"
	"gf_star/internal/model/do"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Users struct {
}

func New() *Users {
	return &Users{}
}

type RegisterInput struct {
	Username string
	Password string
	Email    string
}

func (u *Users) Register(ctx context.Context, user RegisterInput) error {
	if err := u.CheckUserExist(ctx, user.Username); err != nil {
		return err
	}

	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: user.Username,
		Email:    user.Email,
		Password: u.encryptPassword(user.Password),
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) CheckUserExist(ctx context.Context, username string) error {
	count, err := dao.Users.Ctx(ctx).Where("username", username).Count()
	if err != nil {
		return err
	}

	if count > 0 {
		return gerror.New("该用户已存在")
	}

	return nil
}

func (u *Users) encryptPassword(password string) string {
	return gmd5.MustEncryptString(password)
}
