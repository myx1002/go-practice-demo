package word

import (
	"context"

	"gf_star/api/word/v1"
)

func (c *ControllerV1) WordDelete(ctx context.Context, req *v1.WordDeleteReq) (res *v1.WordDeleteRes, err error) {
	uid, err := c.user.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	if err := c.word.Delete(ctx, req.Id, uid); err != nil {
		return nil, err
	}

	return
}
