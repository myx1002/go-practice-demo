package account

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"proxima/app/user/internal/model/entity"
)

func Register(ctx context.Context) (id int, err error) {
	return 1, nil
}

func Login(ctx context.Context) (token string, err error) {
	return "I am token", nil
}

// Info get user info
func Info(ctx context.Context, token string) (user *entity.Users, err error) {
	return &entity.Users{
		Id:        1,
		Username:  "oldme",
		Password:  "123456",
		Email:     "tyyn1022@gmail.com",
		CreatedAt: gtime.New("2024-12-05 22:00:00"),
		UpdatedAt: gtime.New("2024-12-05 22:00:00"),
	}, nil
}
