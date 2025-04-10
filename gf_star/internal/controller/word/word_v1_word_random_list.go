package word

import (
	"context"

	"gf_star/api/word/v1"
)

func (c *ControllerV1) WordRandomList(ctx context.Context, req *v1.WordRandomListReq) (res *v1.WordRandomListRes, err error) {
	uid, err := c.user.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	if words, err := c.word.RandomList(ctx, req.Size, uid); err != nil {
		return nil, err
	} else {
		var list []v1.WordDetailRes
		for _, word := range words {
			list = append(list, v1.WordDetailRes{
				ChineseTranslation: word.ChineseTranslation,
				CreatedAt:          word.CreatedAt,
				UpdatedAt:          word.UpdatedAt,
				Definition:         word.Definition,
				ExampleSentence:    word.ExampleSentence,
				Word:               word.Word,
				Id:                 word.Id,
				ProficiencyLevel:   v1.ProficiencyLevel(word.ProficiencyLevel),
				Pronunciation:      word.Pronunciation,
			})
		}
		return &v1.WordRandomListRes{
			List: list,
		}, nil
	}
}
