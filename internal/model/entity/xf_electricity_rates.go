// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// XfElectricityRates is the golang structure for table xf_electricity_rates.
type XfElectricityRates struct {
	Id         int64       `json:"id"         orm:"id"          ` // 主键
	PeriodAt   string      `json:"periodAt"   orm:"period_at"   ` // 周期月度表
	ValleyRate float64     `json:"valleyRate" orm:"valley_rate" ` // 谷值率
	PeakRate   float64     `json:"peakRate"   orm:"peak_rate"   ` // 峰值率
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  ` // 修改时间
}
