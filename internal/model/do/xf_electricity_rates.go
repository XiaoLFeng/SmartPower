// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// XfElectricityRates is the golang structure of table xf_electricity_rates for DAO operations like Where/Data.
type XfElectricityRates struct {
	g.Meta     `orm:"table:xf_electricity_rates, do:true"`
	Id         interface{} // 主键
	PeriodAt   interface{} // 周期月度表
	ValleyRate interface{} // 谷值率
	PeakRate   interface{} // 峰值率
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
}
