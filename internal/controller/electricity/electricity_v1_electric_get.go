package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"

	"SmartPower/api/electricity/v1"
)

// ElectricGet
//
// # 获取电费
//
// 获取电费接口, 用于获取电费. 用户在这里可以进行获取电费，获取电费的时候不需要用户提供任何参数，获取成功后会返回电费信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricGetReq, 请求
//
// # 返回:
//   - res: *v1.ElectricGetRes, 返回
//   - err: error, 错误
func (c *ControllerV1) ElectricGet(
	ctx context.Context,
	req *v1.ElectricGetReq,
) (res *v1.ElectricGetRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 ElectricGet | 获取电费")
	getElectricity, err := service.Electric().GetElectricity(ctx, *gtime.NewFromStr(req.TimePicker))
	if err != nil {
		return nil, err
	} else {
		res := &v1.ElectricGetRes{
			ElectricCompanyDTO: *getElectricity,
		}
		return res, nil
	}
}
