package model

type CollectionAddInput struct {
	Type     int `json:"type" `
	ObjectId int `json:"object_id"`
	UserId   int `json:"user_id"`
}

type CollectionAddOutput struct {
	CollectionId int `json:"collection_id"`
}

type CollectionCancelInput struct {
	Type     int `json:"type"`
	ObjectId int `json:"object_id"`
}

type CollectionCancelOutput struct {
}

type CollectionListInput struct {
	Type int `json:"type"`
	Page int `json:"page"`
	Size int `json:"size"`
}

type CollectionListItem struct {
	Type         int    `json:"type"`
	ObjectId     int    `json:"object_id"`
	UserId       int    `json:"user_id"`
	ArticleId    int    `json:"article_id"`
	ArticleTitle string `json:"article_title"`
	GoodsId      int    `json:"goods_id"`
	GoodsName    string `json:"goods_name"`
	//ArticleInfo *CollectionArticleListItem `json:"article_info" orm:"with:id=object_id"`
	//GoodsInfo   *CollectionGoodsListItem   `json:"goods_info" orm:"with:id=object_id"`
}

//type CollectionArticleListItem struct {
//	gmeta.Meta `orm:"table:article_info"`
//	Id         int    `json:"id"        description:""`
//	UserId     int    `json:"userId"    description:"作者id"`
//	Title      string `json:"title"     description:"标题"`
//	Desc       string `json:"desc"      description:"摘要"`
//	PicUrl     string `json:"picUrl"    description:"封面图"`
//	IsAdmin    int    `json:"isAdmin"   description:"1后台管理员发布 2前台用户发布"`
//	Praise     int    `json:"praise"    description:"点赞数"`
//}

//type CollectionGoodsListItem struct {
//	gmeta.Meta `orm:"table:goods_info"`
//	Id         int    `json:"id"               description:""`
//	PicUrl     string `json:"picUrl"           description:"图片"`
//	Name       string `json:"name"             description:"商品名称"`
//	Price      int    `json:"price"            description:"价格 单位分"`
//}

type CollectionListOutput struct {
	List  []CollectionListItem
	Page  int
	Size  int
	Total int
}
