// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FileInfoDao is the data access object for table file_info.
type FileInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns FileInfoColumns // columns contains all the column names of Table for convenient usage.
}

// FileInfoColumns defines and stores column names for table file_info.
type FileInfoColumns struct {
	Id           string // 文件id
	Name         string // 文件名称
	RealFileName string //
	Src          string // 文件路径
	Url          string // 访问地址
	UserId       string // 用户id
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
}

// fileInfoColumns holds the columns for table file_info.
var fileInfoColumns = FileInfoColumns{
	Id:           "id",
	Name:         "name",
	RealFileName: "real_file_name",
	Src:          "src",
	Url:          "url",
	UserId:       "user_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewFileInfoDao creates and returns a new DAO object for table data access.
func NewFileInfoDao() *FileInfoDao {
	return &FileInfoDao{
		group:   "default",
		table:   "file_info",
		columns: fileInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FileInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FileInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FileInfoDao) Columns() FileInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FileInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FileInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FileInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
