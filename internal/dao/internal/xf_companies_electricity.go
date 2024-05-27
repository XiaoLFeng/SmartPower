// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// XfCompaniesElectricityDao is the data access object for table xf_companies_electricity.
type XfCompaniesElectricityDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns XfCompaniesElectricityColumns // columns contains all the column names of Table for convenient usage.
}

// XfCompaniesElectricityColumns defines and stores column names for table xf_companies_electricity.
type XfCompaniesElectricityColumns struct {
	Ceuuid                string // 企业电费主键
	Cods                  string //
	PeriodAt              string // 计费月份周期
	ValleyElectricity     string // 谷用电
	ValleyElectricityBill string // 谷电价
	PeakElectricity       string // 峰用电
	PeakElectricityBill   string // 峰电价
	TotalElectricity      string // 合计电量
	TotalBill             string // 合计电费
	CreatedAt             string // 创建时间
	UpdatedAt             string // 更新时间
}

// xfCompaniesElectricityColumns holds the columns for table xf_companies_electricity.
var xfCompaniesElectricityColumns = XfCompaniesElectricityColumns{
	Ceuuid:                "ceuuid",
	Cods:                  "cods",
	PeriodAt:              "period_at",
	ValleyElectricity:     "valley_electricity",
	ValleyElectricityBill: "valley_electricity_bill",
	PeakElectricity:       "peak_electricity",
	PeakElectricityBill:   "peak_electricity_bill",
	TotalElectricity:      "total_electricity",
	TotalBill:             "total_bill",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewXfCompaniesElectricityDao creates and returns a new DAO object for table data access.
func NewXfCompaniesElectricityDao() *XfCompaniesElectricityDao {
	return &XfCompaniesElectricityDao{
		group:   "default",
		table:   "xf_companies_electricity",
		columns: xfCompaniesElectricityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *XfCompaniesElectricityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *XfCompaniesElectricityDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *XfCompaniesElectricityDao) Columns() XfCompaniesElectricityColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *XfCompaniesElectricityDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *XfCompaniesElectricityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *XfCompaniesElectricityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
