package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DataHeadReq struct {
	g.Meta `path:"/head/data" method:"get" tags:"头部卡片数据" dc:"今日订单量、日活和转化率"`
}

type DataHeadRes struct {
	TodayOrderCount int `json:"today_order_count" dc:"今日订单量"`
	DAU             int `json:"dau" dc:"日活量"`
	ConversionRate  int `json:"conversion_rate" dc:"转化率"`
}
