// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderGoodsInfoDao is the data access object for table order_goods_info.
type OrderGoodsInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns OrderGoodsInfoColumns // columns contains all the column names of Table for convenient usage.
}

// OrderGoodsInfoColumns defines and stores column names for table order_goods_info.
type OrderGoodsInfoColumns struct {
	Id             string // 商品维度的订单表
	OrderId        string // 关联的主订单表
	GoodsId        string // 商品id
	GoodsOptionsId string // 商品规格id
	Count          string // 商品数量
	PayType        string // 支付方式 1微信 2支付宝 3云闪付
	Remark         string // 备注
	Status         string // 订单状态 0待支付 1已支付 3已确认收货
	Price          string // 订单金额 单位分
	CouponPrice    string // 优惠券金额 单位分
	ActualPrice    string // 实际支付金额 单位分
	PayAt          string // 支付时间
	CreatedAt      string //
	UpdatedAt      string //
}

// orderGoodsInfoColumns holds the columns for table order_goods_info.
var orderGoodsInfoColumns = OrderGoodsInfoColumns{
	Id:             "id",
	OrderId:        "order_id",
	GoodsId:        "goods_id",
	GoodsOptionsId: "goods_options_id",
	Count:          "count",
	PayType:        "pay_type",
	Remark:         "remark",
	Status:         "status",
	Price:          "price",
	CouponPrice:    "coupon_price",
	ActualPrice:    "actual_price",
	PayAt:          "pay_at",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewOrderGoodsInfoDao creates and returns a new DAO object for table data access.
func NewOrderGoodsInfoDao() *OrderGoodsInfoDao {
	return &OrderGoodsInfoDao{
		group:   "default",
		table:   "order_goods_info",
		columns: orderGoodsInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderGoodsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderGoodsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderGoodsInfoDao) Columns() OrderGoodsInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderGoodsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderGoodsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderGoodsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
