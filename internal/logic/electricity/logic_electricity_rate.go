package electricity

import (
	"SmartPower/internal/config/xerror"
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

// RateAdd
//
// # 添加月电价
//
// 添加月电价, 用于添加月电价. 传入谷电价, 峰电价, 时间.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - valley: float64, 谷电价
//   - peak: float64, 峰电价
//   - timer: gtime.Time, 时间
//
// # 返回:
//   - err: error, 错误信息
func (s *sElectric) RateAdd(ctx context.Context, valley float64, peak float64, timer gtime.Time) (err error) {
	glog.Noticef(ctx, "[CONTROLLER] 执行 RateAdd | 添加月电价")
	getYearMonth := timer.StartOfMonth().Format("200601")
	glog.Debugf(ctx, "获取到的年月: %s", getYearMonth)
	// 检查指定年月是否创建
	var checkRate *entity.XfElectricityRates
	err = dao.XfElectricityRates.Ctx(ctx).Where(do.XfElectricityRates{PeriodAt: &timer}).Scan(&checkRate)
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if checkRate != nil {
		return xerror.NewError(xerror.OperationFailed, "已经存在该月电价")
	}
	// 创建电价
	_, err = dao.XfElectricityRates.Ctx(ctx).Data(do.XfElectricityRates{
		PeriodAt:   getYearMonth,
		ValleyRate: valley,
		PeakRate:   peak,
	}).Insert()
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	return nil
}
