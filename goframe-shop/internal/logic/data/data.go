package data

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
)

type sData struct{}

// 注册服务
func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             grand.Intn(1000),
		ConversionRate:  grand.Intn(30),
	}, nil
}

// 获取当日订单数量
func todayOrderCount(ctx context.Context) int {
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
