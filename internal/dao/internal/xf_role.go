// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// XfRoleDao is the data access object for table xf_role.
type XfRoleDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns XfRoleColumns // columns contains all the column names of Table for convenient usage.
}

// XfRoleColumns defines and stores column names for table xf_role.
type XfRoleColumns struct {
	Ruuid       string // 角色表主键
	Name        string // 角色名
	DisplayName string // 展示名称
	Description string // 描述
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// xfRoleColumns holds the columns for table xf_role.
var xfRoleColumns = XfRoleColumns{
	Ruuid:       "ruuid",
	Name:        "name",
	DisplayName: "display_name",
	Description: "description",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewXfRoleDao creates and returns a new DAO object for table data access.
func NewXfRoleDao() *XfRoleDao {
	return &XfRoleDao{
		group:   "default",
		table:   "xf_role",
		columns: xfRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *XfRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *XfRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *XfRoleDao) Columns() XfRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *XfRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *XfRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *XfRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
