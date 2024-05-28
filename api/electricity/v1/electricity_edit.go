package v1

import "github.com/gogf/gf/v2/frame/g"

// ElectricEditReq
//
// # 修改电费
//
// 修改电费, 用于修改电费. 传入谷电价, 峰电价, 时间.
//
// # 参数:
//   - Valley: float64, 谷电价
//   - Peak: float64, 峰电价
//   - TimePicker: string, 时间
type ElectricEditReq struct {
	g.Meta     `path:"/api/v1/electric" method:"Put" summary:"修改电费" tags:"电控制器"`
	Valley     float64 `json:"valley" v:"required|regex:^\\d+(\\.\\d{1,})$|(^[1-9]\\d*$)#请输入谷电价|谷电价只能为整数或小数" summary:"谷电价" example:"0.2"`
	Peak       float64 `json:"peak" v:"required|regex:^\\d+(\\.\\d{1,})$|(^[1-9]\\d*$)#请输入峰电价|峰电价只能为整数或小数" summary:"峰电价" example:"0.3"`
	TimePicker string  `json:"time_picker" v:"required#请输入时间" summary:"时间" example:"2024-05-01 00:00:00"`
}

// ElectricEditRes
//
// # 修改电费返回
//
// 修改电费返回.
type ElectricEditRes struct {
	g.Meta `mime:"application/json" dc:"修改电费返回"`
}
