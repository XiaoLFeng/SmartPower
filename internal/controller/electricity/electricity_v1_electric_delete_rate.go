package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/electricity/v1"
)

// ElectricDeleteRate
//
// # 删除电费单价
//
// 删除电费单价, 传入电价ID.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricDeleteRateReq, 请求
//
// # 返回:
//   - res: *v1.ElectricDeleteRateRes, 返回
//   - err: error, 错误
func (c *ControllerV1) ElectricDeleteRate(
	ctx context.Context,
	req *v1.ElectricDeleteRateReq,
) (res *v1.ElectricDeleteRateRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 ElectricDeleteRate | 删除电费单价")
	err = service.Electric().RateDelete(ctx, req.RateID)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
