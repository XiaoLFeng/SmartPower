package delectric

import "SmartPower/internal/model/entity"

// ElectricRateDTO
//
// # 电费单价
//
// 电费单价数据传输对象.
//
// # 返回:
//   - total: int64, 总数
//   - rate_id: []*entity.XfElectricityRates, 电价信息
type ElectricRateDTO struct {
	Total int64                        `json:"total"`
	Rate  []*entity.XfElectricityRates `json:"rate"`
}
