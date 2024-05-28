package v1

import (
	"SmartPower/internal/model/dto/dcompany"
	"github.com/gogf/gf/v2/frame/g"
)

// ElectricRegionCalculateReq
//
// # 电费区域计算
//
// 电费区域计算, 用于计算电费. 传入区域, 开始时间, 结束时间.
//
// # 参数:
//   - Region: string, 区域
//   - StartTime: string, 开始时间
//   - EndTime: string, 结束时间
type ElectricRegionCalculateReq struct {
	g.Meta    `path:"/api/v1/electric/region" method:"Get" summary:"电费区域计算" tags:"电控制器"`
	Region    string `json:"region" v:"required#请输入区域" summary:"区域" example:"深圳"`
	StartTime string `json:"start_time" v:"required#请输入开始时间" summary:"开始时间" example:"2024-05-01 00:00:00"`
	EndTime   string `json:"end_time" v:"required#请输入结束时间" summary:"结束时间" example:"2024-05-01 00:00:00"`
}

// ElectricRegionCalculateRes
//
// # 电费区域计算返回
//
// 电费区域计算返回.
type ElectricRegionCalculateRes struct {
	g.Meta `mime:"application/json"`
	dcompany.CompanyRegionDTO
}
