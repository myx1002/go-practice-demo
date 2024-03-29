// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CollectionInfoDao is the data access object for table collection_info.
type CollectionInfoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns CollectionInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CollectionInfoColumns defines and stores column names for table collection_info.
type CollectionInfoColumns struct {
	Id        string //
	UserId    string //
	ObjectId  string //
	Type      string // 收藏类型：1商品 2文章
	CreatedAt string //
	UpdatedAt string //
}

// collectionInfoColumns holds the columns for table collection_info.
var collectionInfoColumns = CollectionInfoColumns{
	Id:        "id",
	UserId:    "user_id",
	ObjectId:  "object_id",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewCollectionInfoDao creates and returns a new DAO object for table data access.
func NewCollectionInfoDao() *CollectionInfoDao {
	return &CollectionInfoDao{
		group:   "default",
		table:   "collection_info",
		columns: collectionInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CollectionInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CollectionInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CollectionInfoDao) Columns() CollectionInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CollectionInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CollectionInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CollectionInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
