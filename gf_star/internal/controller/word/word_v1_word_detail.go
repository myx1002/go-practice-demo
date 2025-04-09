package word

import (
	"context"

	"gf_star/api/word/v1"
)

func (c *ControllerV1) WordDetail(ctx context.Context, req *v1.WordDetailReq) (res *v1.WordDetailRes, err error) {
	uid, err := c.user.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	if word, err := c.word.Detail(ctx, req.Id, uid); err != nil {
		return nil, err
	} else {
		return &v1.WordDetailRes{
			ChineseTranslation: word.ChineseTranslation,
			CreatedAt:          word.CreatedAt,
			Definition:         word.Definition,
			ExampleSentence:    word.ExampleSentence,
			Word:               word.Word,
			Id:                 word.Id,
			ProficiencyLevel:   v1.ProficiencyLevel(word.ProficiencyLevel),
			Pronunciation:      word.Pronunciation,
		}, nil
	}
}
