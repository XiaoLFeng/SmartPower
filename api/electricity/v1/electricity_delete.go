package v1

import "github.com/gogf/gf/v2/frame/g"

// ElectricDeleteReq
//
// # 删除电费
//
// 删除电费单价, 传入电价ID.
//
// # 参数:
//   - CeUUID: string, 电费单价UUID
type ElectricDeleteReq struct {
	g.Meta `path:"/api/v1/electric" method:"Delete" summary:"删除电费" tags:"电控制器"`
	CeUUID string `json:"ce_uuid" v:"required|regex:^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$#请输入电费单价UUID|UUID格式不正确"`
}

// ElectricDeleteRes
//
// # 删除电费返回
//
// 删除电费返回.
type ElectricDeleteRes struct {
	g.Meta `mime:"application/json"`
}
