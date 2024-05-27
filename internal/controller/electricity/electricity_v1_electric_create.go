package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"

	"SmartPower/api/electricity/v1"
)

// ElectricCreate
//
// # 创建电费
//
// 创建电费接口, 用于创建电费. 用户在这里可以进行创建电费，创建电费的时候需要用户提供电费的相关信息，创建成功后会返回创建成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricCreateReq, 创建电费请求
//
// # 返回:
//   - res: *v1.ElectricCreateRes, 创建电费返回
//   - err: error, 错误信息
func (c *ControllerV1) ElectricCreate(ctx context.Context, req *v1.ElectricCreateReq) (res *v1.ElectricCreateRes, err error) {
	glog.Noticef(ctx, "[CONTROLLER] 执行 ElectricCreate | 创建电费")
	// 创建电费
	err = service.Electric().CreateElectricity(ctx, req.ValleyElectricity, req.PeakElectricity, *gtime.NewFromStr(req.TimePicker))
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
