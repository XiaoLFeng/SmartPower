package electricity

import (
	"SmartPower/internal/model/dto/delectric"
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/electricity/v1"
)

// ElectricGetRate
//
// # 获取电费单价
//
// 获取电费单价接口, 用于获取电费单价. 用户在这里可以进行获取电费单价，获取电费单价的时候不需要用户提供任何参数，获取成功后会返回电费单价信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricGetRateReq, 请求
//
// # 返回:
//   - res: *v1.ElectricGetRateRes, 返回
//   - err: error, 错误
func (c *ControllerV1) ElectricGetRate(ctx context.Context, _ *v1.ElectricGetRateReq) (res *v1.ElectricGetRateRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 ElectricGetRate | 获取电费单价")
	rate, err := service.Electric().RateGet(ctx)
	if err != nil {
		return nil, err
	} else {
		res := &v1.ElectricGetRateRes{
			ElectricRateDTO: delectric.ElectricRateDTO{
				Rate:  rate,
				Total: int64(len(rate)),
			},
		}
		return res, nil
	}
}
