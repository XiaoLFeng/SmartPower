package v1

import "github.com/gogf/gf/v2/frame/g"

// ElectricAddRateReq
//
// # 添加月电价
//
// 添加月电价, 用于添加月电价. 传入谷电价, 峰电价, 时间.
//
// # 参数:
//   - ValleyRate: float64, 谷电价
//   - PeakRate: float64, 峰电价
//   - TimePicker: string, 时间
type ElectricAddRateReq struct {
	g.Meta     `path:"/api/v1/electric/rate" method:"Post" summary:"添加月电价" tags:"电控制器"`
	ValleyRate float64 `json:"valley_rate" v:"required#请输入谷电价" summary:"谷电价" example:"0.2"`
	PeakRate   float64 `json:"peak_rate" v:"required#请输入峰电价" summary:"峰电价" example:"0.3"`
	TimePicker string  `json:"time_picker" v:"required#请输入时间" summary:"时间" example:"2024-05-01 00:00:00"`
}

// ElectricAddRateRes
//
// # 添加月电价返回
//
// 添加月电价返回.
type ElectricAddRateRes struct {
	g.Meta `mime:"application/json" dc:"添加月电价返回"`
}
