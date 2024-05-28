package electricity

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"

	"SmartPower/api/electricity/v1"
)

// ElectricRegionCalculate
//
// # 电费区域计算
//
// 电费区域计算, 用于计算电费. 传入区域, 开始时间, 结束时间.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ElectricRegionCalculateReq, 请求
//
// # 返回:
//   - res: *v1.ElectricRegionCalculateRes, 返回
//   - err: error, 错误
func (c *ControllerV1) ElectricRegionCalculate(
	ctx context.Context,
	req *v1.ElectricRegionCalculateReq,
) (res *v1.ElectricRegionCalculateRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 ElectricRegionCalculate | 电费区域计算")
	regionData, err := service.Electric().RegionCalculate(
		ctx,
		req.Region,
		*gtime.NewFromStr(req.StartTime),
		*gtime.NewFromStr(req.EndTime),
	)
	if err != nil {
		return nil, err
	} else {
		return &v1.ElectricRegionCalculateRes{
			CompanyRegionDTO: *regionData,
		}, nil
	}
}
