package user

import (
	"context"
	"gf_star/internal/dao"
	"gf_star/internal/model/do"
)

type Users struct {
}

func New() *Users {
	return &Users{}
}

func (u *Users) Register(ctx context.Context, user do.Users) error {
	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}
