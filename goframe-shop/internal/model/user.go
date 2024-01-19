package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterUserInput struct {
	Name           string
	Avatar         string
	Password       string
	UserSalt       string
	Sex            int
	Status         int
	Sign           string
	SecretQuestion string
	SecretAnswer   string
}

type RegisterUserOutput struct {
	UserId uint
}

type ChangeUserPasswordInput struct {
	Password     string
	SecretAnswer string
	UserSalt     string
}

type ChangeUserPasswordOutput struct {
	UserId uint
}

type CommonUserItem struct {
	g.Meta `orm:"table:user_info"`
	Id     uint
	Name   string
	Avatar string
	Sex    int
	Sign   string
}
