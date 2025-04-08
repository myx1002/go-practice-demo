package user

import (
	"context"

	"gf_star/api/user/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := c.users.Login(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{Token: token}, nil
}
