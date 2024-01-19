package model

type PraiseAddInput struct {
	Type     int `json:"type" `
	ObjectId int `json:"object_id"`
	UserId   int `json:"user_id"`
}

type PraiseAddOutput struct {
	PraiseId int `json:"collection_id"`
}

type PraiseCancelInput struct {
	Type     int `json:"type"`
	ObjectId int `json:"object_id"`
}

type PraiseListInput struct {
	Type int `json:"type"`
	Page int `json:"Page"`
	Size int `json:"size"`
}

type PraiseListItem struct {
	Type         int    `json:"type"`
	ObjectId     int    `json:"object_id"`
	UserId       int    `json:"user_id"`
	ArticleId    int    `json:"article_id"`
	ArticleTitle string `json:"article_title"`
	GoodsId      int    `json:"goods_id"`
	GoodsName    string `json:"goods_name"`
}

type PraiseListOutput struct {
	List  []PraiseListItem
	Page  int
	Size  int
	Total int
}
