package model

type DataHeadOutput struct {
	TodayOrderCount int `json:"today_order_count" dc:"今日订单量"`
	DAU             int `json:"dau" dc:"日活量"`
	ConversionRate  int `json:"conversion_rate" dc:"转化率"`
}
