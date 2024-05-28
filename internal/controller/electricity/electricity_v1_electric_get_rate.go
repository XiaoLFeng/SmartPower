package electricity

import (
	"SmartPower/internal/model/dto/delectric"
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/electricity/v1"
)

func (c *ControllerV1) ElectricGetRate(ctx context.Context, req *v1.ElectricGetRateReq) (res *v1.ElectricGetRateRes, err error) {
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
