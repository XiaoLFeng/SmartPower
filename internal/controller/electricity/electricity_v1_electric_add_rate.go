package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"

	"SmartPower/api/electricity/v1"
)

// ElectricAddRate
//
// # 添加月电价
//
// 添加月电价, 用于添加月电价. 传入谷电价, 峰电价, 时间.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricAddRateReq, 添加月电价请求
//
// # 返回:
//   - res: *v1.ElectricAddRateRes, 添加月电价返回
//   - err: error, 错误信息
func (c *ControllerV1) ElectricAddRate(
	ctx context.Context,
	req *v1.ElectricAddRateReq,
) (res *v1.ElectricAddRateRes, err error) {
	glog.Noticef(ctx, "[CONTROLLER] 执行 ElectricAddRate | 添加月电价")
	// 添加月电价
	err = service.Electric().RateAdd(ctx, req.ValleyRate, req.PeakRate, *gtime.NewFromStr(req.TimePicker))
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
