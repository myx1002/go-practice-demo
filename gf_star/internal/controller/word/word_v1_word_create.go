package word

import (
	"context"
	"gf_star/internal/logic/word"

	"gf_star/api/word/v1"
)

func (c *ControllerV1) WordCreate(ctx context.Context, req *v1.WordCreateReq) (res *v1.WordCreateRes, err error) {
	uid, err := c.user.GetUid(ctx)

	if err != nil {
		return nil, err
	}

	if err := c.word.Create(ctx, word.CreateInput{
		ChineseTranslation: req.ChineseTranslation,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ProficiencyLevel:   req.ProficiencyLevel,
		Pronunciation:      req.Pronunciation,
		Uid:                uid,
		Word:               req.Word,
	}); err != nil {
		return nil, err
	}

	return
}
