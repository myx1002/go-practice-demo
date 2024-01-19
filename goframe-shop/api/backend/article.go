package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleSaveCommon struct {
	Title   string `json:"title"  v:"required#标题不能为空|max-length:20#标题长度不能超过20个字符"   sm:"标题"       description:"标题"`
	Desc    string `json:"desc"    v:"max-length:200#摘要长度不能超过200个字符"  sm:"摘要" description:"摘要"`
	PicUrl  string `json:"pic_url"     v:"required#文章封面图不能为空"     sm:"文章封面图"   description:"文章封面图"`
	Detail  string `json:"detail"   v:"required#文章详情不能为空"      description:"文章详情" sm:"文章详情"`
	IsAdmin string `default:"1"    sm:"标识 1后台管理员发布 2前台用户发布"`
}
type ArticleCreateReq struct {
	g.Meta `path:"/article/add" method:"post" tags:"文章管理" sm:"新增文章"`
	ArticleSaveCommon
}

type ArticleCreateRes struct {
	ArticleId uint `json:"article_id" sm:"文章Id"`
}

type ArticleUpdateReq struct {
	g.Meta    `path:"/article/update" method:"post" tags:"文章管理" sm:"编辑文章"`
	ArticleId uint `json:"article_id" v:"required#文章Id不能为空" sm:"文章Id"`
	ArticleSaveCommon
}

type ArticleUpdateRes struct {
	ArticleId uint `json:"article_id" sm:"文章Id"`
}

type ArticleDeleteReq struct {
	g.Meta    `path:"/article/delete" method:"delete" tags:"文章管理" sm:"删除文章"`
	ArticleId uint `json:"article_id" v:"required#文章Id不能为空" sm:"文章Id"`
}

type ArticleDeleteRes struct {
	ArticleId uint `json:"article_id" sm:"文章Id"`
}

type ArticleListReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"文章管理" sm:"文章列表"`
	Title  string `json:"name" sm:"标题" dc:"根据标题查找"`
	CommonPaginationReq
}

type ArticleListRes struct {
	CommonPaginationRes
}
