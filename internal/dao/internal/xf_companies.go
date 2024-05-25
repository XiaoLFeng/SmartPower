// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// XfCompaniesDao is the data access object for table xf_companies.
type XfCompaniesDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns XfCompaniesColumns // columns contains all the column names of Table for convenient usage.
}

// XfCompaniesColumns defines and stores column names for table xf_companies.
type XfCompaniesColumns struct {
	Cods           string // 企业唯一信用社会代码
	Name           string // 企业名字
	Address        string // 企业所在地址
	Representative string // 法定代表人
	Uuid           string // 指定用户
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// xfCompaniesColumns holds the columns for table xf_companies.
var xfCompaniesColumns = XfCompaniesColumns{
	Cods:           "cods",
	Name:           "name",
	Address:        "address",
	Representative: "representative",
	Uuid:           "uuid",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewXfCompaniesDao creates and returns a new DAO object for table data access.
func NewXfCompaniesDao() *XfCompaniesDao {
	return &XfCompaniesDao{
		group:   "default",
		table:   "xf_companies",
		columns: xfCompaniesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *XfCompaniesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *XfCompaniesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *XfCompaniesDao) Columns() XfCompaniesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *XfCompaniesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *XfCompaniesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *XfCompaniesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
