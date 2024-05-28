package delectric

import (
	"SmartPower/internal/model/dto/dcompany"
	"github.com/gogf/gf/v2/os/gtime"
)

// ElectricCompanyDTO
//
// # 电费公司 DTO
//
// 电费公司 DTO, 用于返回电费公司信息. 用于返回电费公司信息, 包含了电费公司的基本信息和电费公司的电费信息.
//
// # 参数:
//   - Company: dcompany.DCompany, 公司信息
//   - Electricity: ElectricCompanyEntityDTO, 电费信息
type ElectricCompanyDTO struct {
	Company     dcompany.DCompany        `json:"company"`
	Electricity ElectricCompanyEntityDTO `json:"electricity"`
}

type ElectricAllCompanyDTO struct {
	Company     dcompany.DCompany          `json:"company"`
	Total       int64                      `json:"total"`
	Electricity []ElectricCompanyEntityDTO `json:"electricity"`
}

// ElectricCompanyEntityDTO
//
// # 电费公司实体 DTO
//
// 电费公司实体 DTO, 用于返回电费公司信息. 用于返回电费公司信息, 包含了电费公司的基本信息和电费公司的电费信息.
//
// # 参数:
//   - CeUUID: string, 电费公司 UUID
//   - PeriodAt: string, 电费周期
//   - ValleyElectricity: float64, 谷电费
//   - ValleyElectricityBill: float64, 谷电费账单
//   - PeakElectricity: float64, 峰电费
//   - PeakElectricityBill: float64, 峰电费账单
//   - TotalElectricity: float64, 总电费
//   - TotalBill: float64, 总账单
//   - CreatedAt: *gtime.Time, 创建时间
//   - UpdatedAt: *gtime.Time, 更新时间
type ElectricCompanyEntityDTO struct {
	CeUUID                string      `json:"ce_uuid"`
	PeriodAt              string      `json:"period_at"`
	ValleyElectricity     float64     `json:"valley_electricity"`
	ValleyElectricityBill float64     `json:"valley_electricity_bill"`
	PeakElectricity       float64     `json:"peak_electricity"`
	PeakElectricityBill   float64     `json:"peak_electricity_bill"`
	TotalElectricity      float64     `json:"total_electricity"`
	TotalBill             float64     `json:"total_bill"`
	CreatedAt             *gtime.Time `json:"created_at"`
	UpdatedAt             *gtime.Time `json:"updated_at"`
}
