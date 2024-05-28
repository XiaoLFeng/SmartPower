package v1

import (
	"SmartPower/internal/model/dto/delectric"
	"github.com/gogf/gf/v2/frame/g"
)

// ElectricGetRateReq
//
// # 获取电费单价
//
// 获取电费单价接口, 用于获取电费单价. 用户在这里可以进行获取电费单价，获取电费单价的时候不需要用户提供任何参数，获取成功后会返回电费单价信息.
type ElectricGetRateReq struct {
	g.Meta `path:"/api/v1/electric/rate" method:"Get" summary:"获取电费单价" tags:"电控制器"`
}

// ElectricGetRateRes
//
// # 获取电费单价返回
//
// 获取电费单价返回.
type ElectricGetRateRes struct {
	g.Meta `mime:"application/json" dc:"获取电费单价返回"`
	delectric.ElectricRateDTO
}
