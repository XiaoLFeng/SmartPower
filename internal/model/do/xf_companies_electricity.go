// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// XfCompaniesElectricity is the golang structure of table xf_companies_electricity for DAO operations like Where/Data.
type XfCompaniesElectricity struct {
	g.Meta                `orm:"table:xf_companies_electricity, do:true"`
	Ceuuid                interface{} // 企业电费主键
	Cods                  interface{} //
	PeriodAt              *gtime.Time // 计费月份周期
	ValleyElectricity     interface{} // 谷用电
	ValleyElectricityBill interface{} // 谷电价
	PeakElectricity       interface{} // 峰用电
	PeakElectricityBill   interface{} // 峰电价
	TotalElectricity      interface{} // 合计电量
	TotalBill             interface{} // 合计电费
	CreatedAt             *gtime.Time // 创建时间
	UpdatedAt             *gtime.Time // 更新时间
}
