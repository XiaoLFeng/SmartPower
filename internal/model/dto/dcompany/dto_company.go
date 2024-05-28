package dcompany

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DCompany
//
// # 公司
//
// 公司信息.
//
// # 方法
//   - Cods: string, 公司代码
//   - Name: string, 公司名称
//   - Address: string, 公司地址
//   - Representative: string, 公司代表
type DCompany struct {
	Cods           string `json:"cods"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Representative string `json:"representative"`
}

// CompanyRegionDTO
//
// # 电费区域计算
//
// 电费区域计算, 用于计算电费. 传入区域, 开始时间, 结束时间.
//
// # 参数:
//   - Region: string, 区域
//   - StartTime: *gtime.Time, 开始时间
//   - EndTime: *gtime.Time, 结束时间
//   - TotalValleyElectricity: float64, 总谷电量
//   - TotalPeakElectricity: float64, 总峰电量
//   - TotalElectricity: float64, 总电量
//   - TotalValleyBill: float64, 总谷电费
//   - TotalPeakBill: float64, 总峰电费
//   - TotalBill: float64, 总电费
//   - Company: []*CompanyCalculateDTO, 企业
type CompanyRegionDTO struct {
	Region                 string                 `json:"region"`
	StartTime              *gtime.Time            `json:"start_time"`
	EndTime                *gtime.Time            `json:"end_time"`
	TotalValleyElectricity float64                `json:"total_valley_electricity"`
	TotalPeakElectricity   float64                `json:"total_peak_electricity"`
	TotalElectricity       float64                `json:"total_electricity"`
	TotalValleyBill        float64                `json:"total_valley_bill"`
	TotalPeakBill          float64                `json:"total_peak_bill"`
	TotalBill              float64                `json:"total_bill"`
	Company                []*CompanyCalculateDTO `json:"company"`
}

// CompanyCalculateDTO
//
// # 企业计算
//
// 企业计算, 用于计算企业电费. 传入开始时间, 结束时间, 区域, 总谷电量, 总峰电量, 总电量, 总谷电费, 总峰电费, 总电费, 分电费.
//
// # 参数:
//   - Address: string, 区域
//   - TotalValleyElectricity: float64, 总谷电量
//   - TotalPeakElectricity: float64, 总峰电量
//   - TotalElectricity: float64, 总电量
//   - TotalValleyBill: float64, 总谷电费
//   - TotalPeakBill: float64, 总峰电费
//   - TotalBill: float64, 总电费
//   - Rate: []*delectric.ElectricCompanyEntityDTO, 分电费
type CompanyCalculateDTO struct {
	Company                DCompany                    `json:"company"`
	TotalValleyElectricity float64                     `json:"total_valley_electricity"`
	TotalPeakElectricity   float64                     `json:"total_peak_electricity"`
	TotalElectricity       float64                     `json:"total_electricity"`
	TotalValleyBill        float64                     `json:"total_valley_bill"`
	TotalPeakBill          float64                     `json:"total_peak_bill"`
	TotalBill              float64                     `json:"total_bill"`
	Rate                   []*ElectricCompanyEntityDTO `json:"rate"`
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
