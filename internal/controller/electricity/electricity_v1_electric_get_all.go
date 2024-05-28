package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/electricity/v1"
)

// ElectricGetAll
//
// # 获取电费
//
// 获取电费接口, 用于获取电费. 用户在这里可以进行获取电费，获取电费的时候不需要用户提供任何参数，获取成功后会返回电费信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//
// # 返回:
//   - res: *v1.ElectricGetAllRes, 返回
//   - err: error, 错误
func (c *ControllerV1) ElectricGetAll(
	ctx context.Context,
	_ *v1.ElectricGetAllReq,
) (res *v1.ElectricGetAllRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 ElectricGetAll | 获取电费")
	electricity, err := service.Electric().GetAllElectricity(ctx)
	if err != nil {
		return nil, err
	} else {
		res := &v1.ElectricGetAllRes{
			ElectricAllCompanyDTO: *electricity,
		}
		return res, nil
	}
}
