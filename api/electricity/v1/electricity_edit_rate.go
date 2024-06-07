package v1

import "github.com/gogf/gf/v2/frame/g"

// ElectricRateEditReq
//
// # 修改电费单价
//
// 修改电费单价接口, 用于修改电费单价. 用户在这里可以进行修改电费单价，修改电费单价的时候需要用户提供电价ID以及新的电费单价，修改成功后会返回修改成功的信息.
//
// # 参数:
//   - rate_id: int, 电价ID
//   - valley_rate: float64, 谷电价
//   - peak_rate: float64, 峰电价
type ElectricRateEditReq struct {
	g.Meta     `path:"/api/v1/electric/rate" method:"Put" summary:"修改电费单价" tags:"电控制器"`
	RateID     int64   `json:"id" v:"required|电价ID不可缺少"`
	ValleyRate float64 `json:"valley_rate" v:"required|regex:^\\d+(\\.\\d{1,})$|(^[1-9]\\d*$)#请输入谷电价|谷电价只能为整数或小数" summary:"谷电价" example:"0.2"`
	PeakRate   float64 `json:"peak_rate" v:"required|regex:^\\d+(\\.\\d{1,})$|(^[1-9]\\d*$)#请输入峰电价|峰电价只能为整数或小数" summary:"峰电价" example:"0.3"`
}

// ElectricRateEditRes
//
// # 修改电费单价返回
//
// 修改电费单价返回.
type ElectricRateEditRes struct {
	g.Meta `mime:"application/json"`
}
