package v1

import "github.com/gogf/gf/v2/frame/g"

// ElectricCreateReq
//
// # 创建电表
//
// 创建电表, 用于创建电表. 传入谷电, 峰电, 时间.
//
// # 参数:
//   - ValleyElectricity: float64, 谷电
//   - PeakElectricity: float64, 峰电
//   - TimePicker: string, 时间
type ElectricCreateReq struct {
	g.Meta            `path:"/api/v1/electric" method:"Post" summary:"创建电表" tags:"电控制器"`
	ValleyElectricity float64 `json:"valley_electricity" v:"required#请输入谷电" summary:"谷电" example:"20"`
	PeakElectricity   float64 `json:"peak_electricity" v:"required#请输入峰电" summary:"峰电" example:"18"`
	TimePicker        string  `json:"time_picker" v:"required#请输入时间" summary:"时间" example:"2024-05-01 00:00:00"`
}

// ElectricCreateRes
//
// # 创建电表返回
//
// 创建电表返回.
type ElectricCreateRes struct {
	g.Meta `mime:"application/json" dc:"创建电表返回"`
}
