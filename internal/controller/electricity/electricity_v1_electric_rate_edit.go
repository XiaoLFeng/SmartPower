package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/electricity/v1"
)

// ElectricRateEdit
//
// # 编辑电费单价
//
// 编辑电费单价, 传入电价ID和电价信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricRateEditReq, 请求
//
// # 返回:
//   - res: *v1.ElectricRateEditRes, 返回
//   - err: error, 错误
func (c *ControllerV1) ElectricRateEdit(
	ctx context.Context,
	req *v1.ElectricRateEditReq,
) (res *v1.ElectricRateEditRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 ElectricRateEdit | 编辑电费单价")
	err = service.Electric().RateEdit(ctx, req.RateID, req.ValleyRate, req.PeakRate)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
