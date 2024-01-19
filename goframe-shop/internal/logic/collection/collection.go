package collection

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

type sCollection struct {
}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) CollectionAdd(ctx context.Context, in model.CollectionAddInput) (out model.CollectionAddOutput, err error) {
	// 判断对应的收藏数据是否存在
	if in.Type == 1 {
		// 商品
		one, err := dao.GoodsInfo.Ctx(ctx).WherePri(in.ObjectId).One()
		if err != nil || one == nil {
			return model.CollectionAddOutput{}, gerror.New("该商品不存在，收藏失败")
		}
	} else {
		// 文章
		one, err := dao.ArticleInfo.Ctx(ctx).WherePri(in.ObjectId).One()
		if err != nil || one == nil {
			return model.CollectionAddOutput{}, gerror.New("该文章不存在，收藏失败")
		}
	}

	in.UserId = gconv.Int(ctx.Value(consts.CtxUserId))
	save, err := dao.CollectionInfo.Ctx(ctx).Where(g.Map{
		dao.CollectionInfo.Columns().UserId:   in.UserId,
		dao.CollectionInfo.Columns().ObjectId: in.ObjectId,
		dao.CollectionInfo.Columns().Type:     in.Type,
	}).Save(in)
	if err != nil {
		return model.CollectionAddOutput{}, err
	}

	id, err := save.LastInsertId()
	if err != nil {
		return model.CollectionAddOutput{}, err
	}

	return model.CollectionAddOutput{CollectionId: int(id)}, nil
}

func (s *sCollection) CollectionCancel(ctx context.Context, in model.CollectionCancelInput) (out model.CollectionCancelInput, err error) {
	_, err = dao.CollectionInfo.Ctx(ctx).Where(g.Map{
		dao.CollectionInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CollectionInfo.Columns().ObjectId: in.ObjectId,
		dao.CollectionInfo.Columns().Type:     in.Type,
	}).Delete()

	return out, err
}

func (s *sCollection) CollectionList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error) {
	m := dao.CollectionInfo.Ctx(ctx).Where(g.Map{
		dao.CollectionInfo.Columns().UserId: ctx.Value(consts.CtxUserId),
	})

	if in.Type != 0 {
		m = m.Where(dao.CollectionInfo.Columns().Type, in.Type)
	}

	out = &model.CollectionListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	err = m.Page(in.Page, in.Size).
		LeftJoin("article_info", "ar", "ar.id=object_id and type=2").
		LeftJoin("goods_info", "go", "go.id=object_id and type=1").
		Fields("collection_info.*,ar.id as article_id,ar.title as article_title,go.id as goods_id,go.name as goods_name").Scan(&out.List)
	return out, err
}

// 统计文章或商品的被收藏的用户数
func ComputeCollectionCountByObjectId(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.CollectionInfo.Columns().ObjectId: objectId,
		dao.CollectionInfo.Columns().Type:     collectionType,
	}

	count, err = dao.CollectionInfo.Ctx(ctx).Where(condition).Count()
	return
}

// 检查用户是否收藏该文章或商品
func CheckObjectIsCollection(ctx context.Context, objectId uint, collectionType uint8) (bool, error) {
	userId := ctx.Value(consts.CtxUserId)
	if userId == nil {
		return false, nil
	}
	count, err := dao.CollectionInfo.Ctx(ctx).Where(g.Map{
		dao.CollectionInfo.Columns().ObjectId: objectId,
		dao.CollectionInfo.Columns().Type:     collectionType,
		dao.CollectionInfo.Columns().UserId:   userId,
	}).Count()

	if err != nil {
		return false, err
	}

	return count > 0, err
}
