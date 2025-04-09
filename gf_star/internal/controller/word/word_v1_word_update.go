package word

import (
	"context"
	"gf_star/internal/logic/word"

	"gf_star/api/word/v1"
)

func (c *ControllerV1) WordUpdate(ctx context.Context, req *v1.WordUpdateReq) (res *v1.WordUpdateRes, err error) {
	uid, err := c.user.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.word.Update(ctx, word.UpdateInput{
		ChineseTranslation: req.ChineseTranslation,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		Id:                 req.Id,
		ProficiencyLevel:   req.ProficiencyLevel,
		Pronunciation:      req.Pronunciation,
		Uid:                uid,
		Word:               req.Word,
	}); err != nil {
		return nil, err
	}

	return
}
