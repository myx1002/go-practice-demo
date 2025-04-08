package user

import (
	"context"
	"gf_star/internal/model/do"

	"gf_star/api/user/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = c.users.Register(ctx, do.Users{
		Username: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	})
	return nil, err
}
