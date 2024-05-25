// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// XfElectricityRatesDao is the data access object for table xf_electricity_rates.
type XfElectricityRatesDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns XfElectricityRatesColumns // columns contains all the column names of Table for convenient usage.
}

// XfElectricityRatesColumns defines and stores column names for table xf_electricity_rates.
type XfElectricityRatesColumns struct {
	Id         string // 主键
	PeriodAt   string // 周期月度表
	ValleyRate string // 谷值率
	PeakRate   string // 峰值率
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
}

// xfElectricityRatesColumns holds the columns for table xf_electricity_rates.
var xfElectricityRatesColumns = XfElectricityRatesColumns{
	Id:         "id",
	PeriodAt:   "period_at",
	ValleyRate: "valley_rate",
	PeakRate:   "peak_rate",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewXfElectricityRatesDao creates and returns a new DAO object for table data access.
func NewXfElectricityRatesDao() *XfElectricityRatesDao {
	return &XfElectricityRatesDao{
		group:   "default",
		table:   "xf_electricity_rates",
		columns: xfElectricityRatesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *XfElectricityRatesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *XfElectricityRatesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *XfElectricityRatesDao) Columns() XfElectricityRatesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *XfElectricityRatesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *XfElectricityRatesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *XfElectricityRatesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
