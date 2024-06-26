package electricity

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/database/gdb"
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
	glog.Noticef(ctx, "[LOGIC] 执行 RateAdd | 添加月电价")
	getYearMonth := timer.StartOfMonth().Format("Ym")
	glog.Debugf(ctx, "获取到的年月: %s", getYearMonth)
	// 检查指定年月是否创建
	var checkRate *entity.XfElectricityRates
	err = dao.XfElectricityRates.Ctx(ctx).Where(do.XfElectricityRates{PeriodAt: getYearMonth}).Scan(&checkRate)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if checkRate != nil {
		return berror.NewError(bcode.OperationFailed, "已经存在该月电价")
	}
	// 创建电价
	_, err = dao.XfElectricityRates.Ctx(ctx).Data(do.XfElectricityRates{
		PeriodAt:   getYearMonth,
		ValleyRate: valley,
		PeakRate:   peak,
	}).Insert()
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return nil
}

// RateEdit
//
// # 编辑月电价
//
// 编辑月电价, 用于编辑月电价. 传入电价ID, 谷电价, 峰电价.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - rateID: int64, 电价ID
//   - valley: float64, 谷电价
//   - peak: float64, 峰电价
//
// # 返回:
//   - err: error, 错误信息
func (s *sElectric) RateEdit(ctx context.Context, rateID int64, valley float64, peak float64) (err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 RateEdit | 编辑月电价")
	// 检查电价 ID 是否存在
	var getRate *entity.XfElectricityRates
	err = dao.XfElectricityRates.Ctx(ctx).Where(do.XfElectricityRates{Id: rateID}).Scan(&getRate)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getRate == nil {
		return berror.NewError(bcode.OperationFailed, "电价 ID 不存在")
	}
	// 修改电价信息
	_, err = dao.XfElectricityRates.Ctx(ctx).Data(do.XfElectricityRates{
		ValleyRate: valley,
		PeakRate:   peak,
	}).Where(do.XfElectricityRates{Id: rateID}).Update()
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return nil
}

// RateDelete
//
// # 删除月电价
//
// 删除月电价, 用于删除月电价. 传入电价ID.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - rateID: int64, 电价ID
//
// # 返回:
//   - err: error, 错误信息
func (s *sElectric) RateDelete(ctx context.Context, rateID int64) (err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 RateDelete | 删除月电价")
	// 检查电价 ID 是否存在
	var getRate *entity.XfElectricityRates
	err = dao.XfElectricityRates.Ctx(ctx).Where(do.XfElectricityRates{Id: rateID}).Scan(&getRate)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if getRate == nil {
		return berror.NewError(bcode.OperationFailed, "电价 ID 不存在")
	}
	// 删除电价信息
	err = dao.XfElectricityRates.Transaction(ctx, func(_ context.Context, tx gdb.TX) error {
		// 删除指定月份的企业计费信息
		_, err = dao.XfCompaniesElectricity.Ctx(ctx).
			Where(do.XfCompaniesElectricity{PeriodAt: getRate.PeriodAt}).
			Delete()
		if err != nil {
			return err
		}
		// 删除电价信息
		_, err := dao.XfElectricityRates.Ctx(ctx).
			Where(do.XfElectricityRates{Id: rateID}).
			Delete()
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

// RateGet
//
// # 获取月电价
//
// 获取月电价, 用于获取月电价. 传入电价ID.
//
// # 参数:
//   - ctx: context.Context, 上下文
//
// # 返回:
//   - err: error, 错误信息
func (s *sElectric) RateGet(ctx context.Context) (rate []*entity.XfElectricityRates, err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 RateGet | 获取月电价")
	var getRate []*entity.XfElectricityRates
	err = dao.XfElectricityRates.Ctx(ctx).OrderDesc("period_at").Scan(&getRate)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return getRate, nil
}
