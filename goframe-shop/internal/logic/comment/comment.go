package comment

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

type sComment struct {
}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) CommentAdd(ctx context.Context, in model.CommentAddInput) (out *model.CommentAddOutput, err error) {
	// 判断对应的收藏数据是否存在
	if in.Type == 1 {
		// 商品
		one, err := dao.GoodsInfo.Ctx(ctx).WherePri(in.ObjectId).One()
		if err != nil || one == nil {
			return nil, gerror.New("该商品不存在，评论失败")
		}
	} else {
		// 文章
		one, err := dao.ArticleInfo.Ctx(ctx).WherePri(in.ObjectId).One()
		if err != nil || one == nil {
			return nil, gerror.New("该文章不存在，评论失败")
		}
	}

	// 判断父级评论是否存在
	if in.ParentId > 0 {
		one, err := dao.CommentInfo.Ctx(ctx).WherePri(in.ParentId).One()
		if err != nil || one == nil {
			return nil, gerror.New("该条评论已消失，请重新评论")
		}
	}

	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CommentInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.CommentAddOutput{
		CommentId: uint(id),
	}, nil
}

func (s *sComment) CommentDelete(ctx context.Context, in model.CommentDeleteInput) (err error) {
	_, err = dao.CommentInfo.Ctx(ctx).
		Where(dao.CommentInfo.Columns().UserId, gconv.Uint(ctx.Value(consts.CtxUserId))).
		Where(dao.CommentInfo.Columns().Id, in.CommentId).Delete()
	return err
}

func (s *sComment) CommentList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error) {
	m := dao.CommentInfo.Ctx(ctx).Where(g.Map{
		dao.CollectionInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CollectionInfo.Columns().ObjectId: in.ObjectId,
		dao.CollectionInfo.Columns().Type:     in.Type,
	})

	out = &model.CommentListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	err = m.Page(in.Page, in.Size).WithAll().Scan(&out.List)
	return out, err
}
