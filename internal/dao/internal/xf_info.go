// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// XfInfoDao is the data access object for table xf_info.
type XfInfoDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns XfInfoColumns // columns contains all the column names of Table for convenient usage.
}

// XfInfoColumns defines and stores column names for table xf_info.
type XfInfoColumns struct {
	SystemId  string // 系统主键
	Keyword   string // 键
	Value     string // 值
	UpdatedAt string // 更新时间
}

// xfInfoColumns holds the columns for table xf_info.
var xfInfoColumns = XfInfoColumns{
	SystemId:  "system_id",
	Keyword:   "keyword",
	Value:     "value",
	UpdatedAt: "updated_at",
}

// NewXfInfoDao creates and returns a new DAO object for table data access.
func NewXfInfoDao() *XfInfoDao {
	return &XfInfoDao{
		group:   "default",
		table:   "xf_info",
		columns: xfInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *XfInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *XfInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *XfInfoDao) Columns() XfInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *XfInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *XfInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *XfInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
