package electricity

import (
	"SmartPower/internal/config/xerror"
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/entity"
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

// CreateElectricity
//
// # 创建电费
//
// 创建电费接口, 用于创建电费. 用户在这里可以进行创建电费，创建电费的时候需要用户提供电费的相关信息，创建成功后会返回创建成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - valley: float32, 谷电费
//   - peak: float32, 峰电费
//
// # 返回:
//   - err: error, 错误信息
func (s *sElectric) CreateElectricity(
	ctx context.Context,
	valley float64,
	peak float64,
	timePicker gtime.Time,
) (err error) {
	glog.Noticef(ctx, "[CONTROLLER] 执行 CreateElectricity | 创建电费")
	getYearMonth := timePicker.StartOfMonth().Format("200601")
	// 获取用户所在的企业
	getRequest := g.RequestFromCtx(ctx)
	getToken := getRequest.Header.Get("Authorization")
	userInfo, err := service.Token().GetUserByToken(ctx, getToken)
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if userInfo == nil {
		return xerror.NewError(xerror.OperationFailed, "用户不存在")
	}
	// 根据用户查询企业
	var getCompany *entity.XfCompanies
	err = dao.XfCompanies.Ctx(ctx).Where(do.XfCompanies{Uuid: userInfo.Uuid}).Scan(&getCompany)
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if getCompany == nil {
		return xerror.NewError(xerror.OperationFailed, "企业不存在")
	}
	// 检查 timePicker 是否与之前统计重复
	var getElectricity *entity.XfCompaniesElectricity
	err = dao.XfCompaniesElectricity.Ctx(ctx).
		Where(do.XfCompaniesElectricity{PeriodAt: &timePicker, Cods: getCompany.Cods}).
		Scan(&getElectricity)
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if getElectricity == nil {
		// 获取指定月份电价
		var getElectricRate *entity.XfElectricityRates
		err := dao.XfElectricityRates.Ctx(ctx).Where(do.XfElectricityRates{PeriodAt: getYearMonth}).Scan(&getElectricRate)
		if err != nil {
			return xerror.NewErrorHasError(xerror.ServerInternalError, err)
		}
		if getElectricRate == nil {
			return xerror.NewError(xerror.OperationFailed, "该月电价未创建，请联系管理员")
		}
		// 计算电费
		valleyElectricityMoney := valley * getElectricRate.ValleyRate
		peakElectricityMoney := peak * getElectricRate.PeakRate
		// 创建电费
		getUUID, _ := uuid.NewV7()
		_, err = dao.XfCompaniesElectricity.Ctx(ctx).Data(do.XfCompaniesElectricity{
			Ceuuid:                getUUID.String(),
			Cods:                  getCompany.Cods,
			PeriodAt:              getYearMonth,
			ValleyElectricity:     valley,
			ValleyElectricityBill: valleyElectricityMoney,
			PeakElectricity:       peak,
			PeakElectricityBill:   peakElectricityMoney,
			TotalElectricity:      valley + peak,
			TotalBill:             valleyElectricityMoney + peakElectricityMoney,
		}).Insert()
		if err != nil {
			return xerror.NewErrorHasError(xerror.ServerInternalError, err)
		} else {
			return nil
		}
	} else {
		return xerror.NewError(xerror.OperationFailed, "该月电表已被创建")
	}
}
