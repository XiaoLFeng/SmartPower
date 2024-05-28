package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/electricity/v1"
)

// ElectricDelete
//
// # 删除电费
//
// 删除电费单价, 传入电价ID.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricDeleteReq, 请求
//
// # 返回:
//   - res: *v1.ElectricDeleteRes, 返回
//   - err: error, 错误
func (c *ControllerV1) ElectricDelete(
	ctx context.Context,
	req *v1.ElectricDeleteReq,
) (res *v1.ElectricDeleteRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 ElectricDelete | 删除电费")
	err = service.Electric().DeleteElectricity(ctx, req.CeUUID)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
