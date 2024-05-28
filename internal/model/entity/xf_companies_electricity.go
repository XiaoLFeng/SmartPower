// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// XfCompaniesElectricity is the golang structure for table xf_companies_electricity.
type XfCompaniesElectricity struct {
	Ceuuid                string      `json:"ceuuid"                orm:"ceuuid"                  ` // 企业电费主键
	Cods                  string      `json:"cods"                  orm:"cods"                    ` //
	PeriodAt              string      `json:"periodAt"              orm:"period_at"               ` // 计费月份周期
	ValleyElectricity     float64     `json:"valleyElectricity"     orm:"valley_electricity"      ` // 谷用电
	ValleyElectricityBill float64     `json:"valleyElectricityBill" orm:"valley_electricity_bill" ` // 谷电价
	PeakElectricity       float64     `json:"peakElectricity"       orm:"peak_electricity"        ` // 峰用电
	PeakElectricityBill   float64     `json:"peakElectricityBill"   orm:"peak_electricity_bill"   ` // 峰电价
	TotalElectricity      float64     `json:"totalElectricity"      orm:"total_electricity"       ` // 合计电量
	TotalBill             float64     `json:"totalBill"             orm:"total_bill"              ` // 合计电费
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"              ` // 创建时间
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"              ` // 更新时间
}
