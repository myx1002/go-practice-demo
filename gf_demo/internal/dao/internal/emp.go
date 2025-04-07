// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmpDao is the data access object for the table emp.
type EmpDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  EmpColumns         // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// EmpColumns defines and stores column names for the table emp.
type EmpColumns struct {
	Id     string // ID
	DeptId string // 所属部门
	Name   string // 姓名
	Gender string // 性别: 0=男 1=女
	Phone  string // 联系电话
	Email  string // 邮箱
	Avatar string // 照片
}

// empColumns holds the columns for the table emp.
var empColumns = EmpColumns{
	Id:     "id",
	DeptId: "dept_id",
	Name:   "name",
	Gender: "gender",
	Phone:  "phone",
	Email:  "email",
	Avatar: "avatar",
}

// NewEmpDao creates and returns a new DAO object for table data access.
func NewEmpDao(handlers ...gdb.ModelHandler) *EmpDao {
	return &EmpDao{
		group:    "default",
		table:    "emp",
		columns:  empColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EmpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EmpDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EmpDao) Columns() EmpColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EmpDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EmpDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *EmpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
