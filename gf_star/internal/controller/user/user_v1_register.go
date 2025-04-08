package user

import (
	"context"
	"gf_star/api/user/v1"
	"gf_star/internal/logic/user"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = c.users.Register(ctx, user.RegisterInput{
		Username: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	})
	return nil, err
}
