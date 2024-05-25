// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// XfUserDao is the data access object for table xf_user.
type XfUserDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns XfUserColumns // columns contains all the column names of Table for convenient usage.
}

// XfUserColumns defines and stores column names for table xf_user.
type XfUserColumns struct {
	Uuid      string // 用户 id
	Username  string // 用户名
	Email     string // 邮箱
	Phone     string // 手机号
	Password  string // 用户密码
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// xfUserColumns holds the columns for table xf_user.
var xfUserColumns = XfUserColumns{
	Uuid:      "uuid",
	Username:  "username",
	Email:     "email",
	Phone:     "phone",
	Password:  "password",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewXfUserDao creates and returns a new DAO object for table data access.
func NewXfUserDao() *XfUserDao {
	return &XfUserDao{
		group:   "default",
		table:   "xf_user",
		columns: xfUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *XfUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *XfUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *XfUserDao) Columns() XfUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *XfUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *XfUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *XfUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
