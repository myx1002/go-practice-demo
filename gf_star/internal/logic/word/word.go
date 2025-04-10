package word

import (
	"context"
	v1 "gf_star/api/word/v1"
	"gf_star/internal/dao"
	"gf_star/internal/model/do"
	"gf_star/internal/model/entity"
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

type UpdateInput struct {
	Id                 uint
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Word) Update(ctx context.Context, in UpdateInput) error {
	var cl = dao.Words.Columns()

	// 判断这个id的单词是否存在
	exist, err := dao.Words.Ctx(ctx).
		Where(cl.Uid, in.Uid).
		Where(cl.Id, in.Id).
		Exist()
	if err != nil {
		return err
	}

	if !exist {
		return gerror.New("单词不存在")
	}

	// 判断word是否重复
	exist, err = dao.Words.Ctx(ctx).
		Where(cl.Uid, in.Uid).
		Where(cl.Word, in.Word).
		WhereNot(cl.Id, in.Id).
		Exist()

	if err != nil {
		return err
	}

	if exist {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Where(cl.Id, in.Id).Update()

	if err != nil {
		return err
	}

	return nil
}

func (w *Word) CheckWordExists(ctx context.Context, word string, uid uint) error {
	var cl = dao.Words.Columns()

	count, err := dao.Words.Ctx(ctx).Data(g.Map{
		cl.Uid:  uid,
		cl.Word: word,
	}).Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return gerror.New("单词已存在")
	}

	return nil
}

func (w *Word) Create(ctx context.Context, in CreateInput) error {
	if err := w.CheckWordExists(ctx, in.Word, in.Uid); err != nil {
		return err
	}

	_, err := dao.Words.Ctx(ctx).Data(do.Words{
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

func (w *Word) Detail(ctx context.Context, id, uid uint) (word *entity.Words, err error) {
	err = dao.Words.Ctx(ctx).Where(g.Map{
		dao.Words.Columns().Id:  id,
		dao.Words.Columns().Uid: uid,
	}).Scan(&word)

	if err != nil {
		return nil, err
	}
	return
}

func (w *Word) Delete(ctx context.Context, id, uid uint) error {
	_, err := dao.Words.Ctx(ctx).Where(g.Map{
		dao.Words.Columns().Id:  id,
		dao.Words.Columns().Uid: uid,
	}).Delete()

	if err != nil {
		return err
	}

	return nil
}

func (w *Word) List(ctx context.Context, uid uint, req v1.WordListReq) (words []*entity.Words, total int, err error) {
	var cl = dao.Words.Columns()
	orm := dao.Words.Ctx(ctx).Where(cl.Uid, uid)

	if len(req.Word) != 0 {
		orm = orm.WhereLike(cl.Word, "%"+req.Word+"%")
	}

	err = orm.OrderDesc(cl.CreatedAt).OrderDesc(cl.Id).Page(req.Page, req.Size).ScanAndCount(&words, &total, false)

	if err != nil {
		return nil, 0, err
	}

	return
}

func (w *Word) RandomList(ctx context.Context, size int, uid uint) (words []*entity.Words, err error) {
	err = dao.Words.Ctx(ctx).Where("uid", uid).Limit(size).OrderRandom().Scan(&words)
	return
}

func (w *Word) SetLevel(ctx context.Context, id uint, level v1.ProficiencyLevel, uid uint) (err error) {
	exist, err := dao.Words.Ctx(ctx).Data(entity.Words{
		Id:  id,
		Uid: uid,
	}).Exist()

	if err != nil {
		return
	}

	if !exist {
		return gerror.New("单词不存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(g.Map{
		dao.Words.Columns().ProficiencyLevel: uint(level),
	}).Where(g.Map{
		dao.Words.Columns().Id:  id,
		dao.Words.Columns().Uid: uid,
	}).Update()

	return
}
