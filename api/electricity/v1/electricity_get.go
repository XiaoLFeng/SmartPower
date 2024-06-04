package v1

import (
	"SmartPower/internal/model/dto/delectric"
	"github.com/gogf/gf/v2/frame/g"
)

// ElectricGetReq
//
// # 获取电费
//
// 获取电费接口, 用于获取电费. 用户在这里可以进行获取电费，获取电费的时候需要用户提供时间参数，获取成功后会返回电费信息.
//
// # 参数:
//   - time_picker: string, 时间
type ElectricGetReq struct {
	g.Meta     `path:"/api/v1/electric" method:"Get" summary:"获取电费" tags:"电控制器"`
	CeUUID     string `json:"ce_uuid" summary:"企业电费主键" example:"018fe1fd-8abd-7f36-97b0-0bcf08daa8e2"`
	TimePicker string `json:"time_picker" summary:"时间" example:"2024-05-01 00:00:00"`
}

// ElectricGetRes
//
// # 获取电费返回
//
// 获取电费返回.
type ElectricGetRes struct {
	g.Meta `mime:"application/json"`
	delectric.ElectricCompanyDTO
}

// ElectricGetAllReq
//
// # 获取所有电费
//
// 获取所有电费接口, 用于获取所有电费. 用户在这里可以进行获取所有电费，获取成功后会返回所有电费信息.
type ElectricGetAllReq struct {
	g.Meta `path:"/api/v1/electric/all" method:"Get" summary:"获取所有电费" tags:"电控制器"`
}

// ElectricGetAllRes
//
// # 获取所有电费返回
//
// 获取所有电费返回.
type ElectricGetAllRes struct {
	g.Meta `mime:"application/json"`
	delectric.ElectricAllCompanyDTO
}
