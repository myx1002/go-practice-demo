package word

import (
	"context"
	"gf_star/api/word/v1"
)

func (c *ControllerV1) WordList(ctx context.Context, req *v1.WordListReq) (res *v1.WordListRes, err error) {
	uid, err := c.user.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	if words, total, err := c.word.List(ctx, uid, *req); err != nil {
		return nil, err
	} else {
		var list []v1.WordDetailRes
		if len(words) == 0 {
			return &v1.WordListRes{
				List:  []v1.WordDetailRes{},
				Total: 0,
			}, nil
		}
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
		res = &v1.WordListRes{
			List:  list,
			Total: total,
		}
	}

	return
}
