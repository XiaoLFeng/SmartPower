package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"

	"SmartPower/api/electricity/v1"
)

// ElectricEdit
//
// # 编辑电费
//
// 编辑电费接口, 用于编辑电费. 用户在这里可以进行编辑电费，编辑电费的时候需要用户提供时间参数，编辑成功后会返回电费信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricEditReq, 请求
//
// # 返回:
//   - res: *v1.ElectricEditRes, 返回
//   - err: error, 错误
func (c *ControllerV1) ElectricEdit(ctx context.Context, req *v1.ElectricEditReq) (res *v1.ElectricEditRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 ElectricEdit | 编辑电费")
	err = service.Electric().EditElectricity(ctx, req.Valley, req.Peak, *gtime.NewFromStr(req.TimePicker))
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
