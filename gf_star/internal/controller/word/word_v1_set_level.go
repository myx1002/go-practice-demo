package word

import (
	"context"

	"gf_star/api/word/v1"
)

func (c *ControllerV1) SetLevel(ctx context.Context, req *v1.SetLevelReq) (res *v1.SetLevelRes, err error) {

	uid, err := c.user.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	if err := c.word.SetLevel(ctx, req.Id, req.Level, uid); err != nil {
		return nil, err
	} else {
		return &v1.SetLevelRes{}, nil
	}
}
