package praise

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sPraise struct {
}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

func (s *sPraise) PraiseAdd(ctx context.Context, in model.PraiseAddInput) (out model.PraiseAddOutput, err error) {
	// 判断对应的收藏数据是否存在
	if in.Type == 1 {
		// 商品
		one, err := dao.GoodsInfo.Ctx(ctx).WherePri(in.ObjectId).One()
		if err != nil || one == nil {
			return model.PraiseAddOutput{}, gerror.New("该商品不存在，点赞失败")
		}
	} else {
		// 文章
		one, err := dao.ArticleInfo.Ctx(ctx).WherePri(in.ObjectId).One()
		if err != nil || one == nil {
			return model.PraiseAddOutput{}, gerror.New("该文章不存在，点赞失败")
		}
	}
	in.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	save, err := dao.PraiseInfo.Ctx(ctx).Where(g.Map{
		dao.PraiseInfo.Columns().UserId:   in.UserId,
		dao.PraiseInfo.Columns().ObjectId: in.ObjectId,
		dao.PraiseInfo.Columns().Type:     in.Type,
	}).Data(in).Save()

	id, err := save.LastInsertId()
	if err != nil {
		return model.PraiseAddOutput{}, err
	}

	return model.PraiseAddOutput{PraiseId: int(id)}, nil
}

func (s *sPraise) PraiseCancel(ctx context.Context, in model.PraiseCancelInput) (err error) {
	_, err = dao.PraiseInfo.Ctx(ctx).Where(g.Map{
		dao.PraiseInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.PraiseInfo.Columns().ObjectId: in.ObjectId,
		dao.PraiseInfo.Columns().Type:     in.Type,
	}).Delete()

	return err
}

func (s *sPraise) PraiseList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error) {
	m := dao.PraiseInfo.Ctx(ctx).Where(dao.PraiseInfo.Columns().UserId, ctx.Value(consts.CtxUserId))
	out = &model.PraiseListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	if in.Type != 0 {
		m = m.Where(dao.PraiseInfo.Columns().Type, in.Type)
	}

	out.Total, err = m.
		LeftJoin("article_info", "ar", "ar.id=object_id and type=2").
		LeftJoin("goods_info", "go", "go.id=object_id and type=1").
		Count()
	if err != nil {
		return out, err
	}

	err = m.Page(in.Page, in.Size).
		LeftJoin("article_info", "ar", "ar.id=object_id and type=2").
		LeftJoin("goods_info", "go", "go.id=object_id and type=1").
		Fields("praise_info.*,ar.id as article_id,ar.title as article_title,go.id as goods_id,go.name as goods_name").Scan(&out.List)
	return out, err
}

// 统计文章或商品的被点赞的用户数
func ComputePraiseCountByObjectId(ctx context.Context, objectId uint, praiseType uint8) (count int, err error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().ObjectId: objectId,
		dao.PraiseInfo.Columns().Type:     praiseType,
	}

	count, err = dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	return
}

// 检查用户是否点赞该文章或商品
func CheckObjectIsPraise(ctx context.Context, objectId uint, praiseType uint8) (bool, error) {
	userId := ctx.Value(consts.CtxUserId)
	if userId == nil {
		return false, nil
	}
	count, err := dao.PraiseInfo.Ctx(ctx).Where(g.Map{
		dao.PraiseInfo.Columns().ObjectId: objectId,
		dao.PraiseInfo.Columns().Type:     praiseType,
		dao.PraiseInfo.Columns().UserId:   userId,
	}).Count()

	if err != nil {
		return false, err
	}

	return count > 0, err
}
