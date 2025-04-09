package word

import (
	"context"
	v1 "gf_star/api/word/v1"
	"gf_star/internal/dao"
	"gf_star/internal/model/do"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type Word struct {
}

func New() *Word {
	return &Word{}
}

type CreateInput struct {
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Word) Create(ctx context.Context, in CreateInput) error {
	var cl = dao.Words.Columns()

	count, err := dao.Words.Ctx(ctx).Data(g.Map{
		cl.Uid:  in.Uid,
		cl.Word: in.Word,
	}).Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Uid:                in.Uid,
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Insert()

	if err != nil {
		return err
	}

	return nil
}
