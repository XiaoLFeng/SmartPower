package v1

import "github.com/gogf/gf/v2/frame/g"

// ElectricDeleteRateReq
//
// # 删除电费单价
//
// 删除电费单价, 传入电价ID.
//
// # 参数:
//   - RateID: int, 电价ID
type ElectricDeleteRateReq struct {
	g.Meta `path:"/api/v1/electric/rate" method:"Delete" summary:"删除电费单价" tags:"电控制器"`
	RateID int64 `json:"id" v:"required|电价ID不可缺少"`
}

// ElectricDeleteRateRes
//
// # 删除电费单价返回
//
// 删除电费单价返回.
type ElectricDeleteRateRes struct {
	g.Meta `mime:"application/json"`
}
