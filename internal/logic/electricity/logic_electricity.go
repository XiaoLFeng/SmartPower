package electricity

import (
	"SmartPower/internal/config/xerror"
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/dto/dcompany"
	"SmartPower/internal/model/dto/delectric"
	"SmartPower/internal/model/entity"
	"context"
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
	glog.Noticef(ctx, "[LOGIC] 执行 CreateElectricity | 创建电费")
	getYearMonth := timePicker.StartOfMonth().Format("Ym")
	getCompany, err := s.GetCompanyByHeader(ctx)
	if err != nil {
		return err
	}
	// 检查 timePicker 是否与之前统计重复
	var getElectricity *entity.XfCompaniesElectricity
	err = dao.XfCompaniesElectricity.Ctx(ctx).
		Where(do.XfCompaniesElectricity{PeriodAt: getYearMonth, Cods: getCompany.Cods}).
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

// GetElectricity
//
// # 获取电费
//
// 获取电费接口, 用于获取电费. 用户在这里可以进行获取电费，获取电费的时候需要用户提供获取电费的时间，获取成功后会返回获取成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - timer: gtime.Time, 时间
//
// # 返回:
//   - getElectricity: *delectric.ElectricCompanyDTO, 电费信息
//   - err: error, 错误信息
func (s *sElectric) GetElectricity(
	ctx context.Context,
	timer gtime.Time,
) (getElectricity *delectric.ElectricCompanyDTO, err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 ElectricGet | 获取电费")
	getYearMonth := timer.StartOfMonth().Format("Ym")
	getCompany, err := s.GetCompanyByHeader(ctx)
	if err != nil {
		return nil, err
	}
	// 获取电费
	var getElectric *entity.XfCompaniesElectricity
	err = dao.XfCompaniesElectricity.Ctx(ctx).
		Data(do.XfCompaniesElectricity{Ceuuid: getCompany.Cods, PeriodAt: getYearMonth}).
		Scan(&getElectric)
	if err != nil {
		return nil, xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if getElectric == nil {
		return nil, xerror.NewError(xerror.OperationFailed, "该月电费未创建")
	} else {
		return &delectric.ElectricCompanyDTO{
			Company: dcompany.DCompany{
				Cods:           getCompany.Cods,
				Name:           getCompany.Name,
				Address:        getCompany.Address,
				Representative: getCompany.Representative,
			},
			Electricity: delectric.ElectricCompanyEntityDTO{
				CeUUID:                getElectric.Ceuuid,
				PeriodAt:              getElectric.PeriodAt,
				ValleyElectricity:     getElectric.ValleyElectricity,
				ValleyElectricityBill: getElectric.ValleyElectricityBill,
				PeakElectricity:       getElectric.PeakElectricity,
				PeakElectricityBill:   getElectric.PeakElectricityBill,
				TotalElectricity:      getElectric.TotalElectricity,
				TotalBill:             getElectric.TotalBill,
				CreatedAt:             getElectric.CreatedAt,
				UpdatedAt:             getElectric.UpdatedAt,
			},
		}, nil
	}
}

// GetAllElectricity
//
// # 获取所有电费
//
// 获取所有电费接口, 用于获取所有电费. 用户在这里可以进行获取所有电费，获取成功后会返回获取成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//
// # 返回:
//   - electricity: delectric.ElectricAllCompanyDTO, 电费信息
//   - err: error, 错误信息
func (s *sElectric) GetAllElectricity(ctx context.Context) (electricity *delectric.ElectricAllCompanyDTO, err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 GetAllElectricity | 获取所有电费")
	getCompany, err := s.GetCompanyByHeader(ctx)
	if err != nil {
		return nil, err
	}
	// 获取电费
	var getElectric []*entity.XfCompaniesElectricity
	err = dao.XfCompaniesElectricity.Ctx(ctx).
		Where(do.XfCompaniesElectricity{Cods: getCompany.Cods}).
		OrderDesc("period_at").
		Scan(&getElectric)
	if err != nil {
		return nil, xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if getElectric == nil {
		electricity := &delectric.ElectricAllCompanyDTO{
			Company: dcompany.DCompany{
				Cods:           getCompany.Cods,
				Name:           getCompany.Name,
				Address:        getCompany.Address,
				Representative: getCompany.Representative,
			},
			Electricity: nil,
		}
		return electricity, nil
	} else {
		var electricDTO []delectric.ElectricCompanyEntityDTO
		for _, companiesElectricity := range getElectric {
			electricDTO = append(electricDTO, delectric.ElectricCompanyEntityDTO{
				CeUUID:                companiesElectricity.Ceuuid,
				PeriodAt:              companiesElectricity.PeriodAt,
				ValleyElectricity:     companiesElectricity.ValleyElectricity,
				ValleyElectricityBill: companiesElectricity.ValleyElectricityBill,
				PeakElectricity:       companiesElectricity.PeakElectricity,
				PeakElectricityBill:   companiesElectricity.PeakElectricityBill,
				TotalElectricity:      companiesElectricity.TotalElectricity,
				TotalBill:             companiesElectricity.TotalBill,
				CreatedAt:             companiesElectricity.CreatedAt,
				UpdatedAt:             companiesElectricity.UpdatedAt,
			})
		}
		electricity := &delectric.ElectricAllCompanyDTO{
			Company: dcompany.DCompany{
				Cods:           getCompany.Cods,
				Name:           getCompany.Name,
				Address:        getCompany.Address,
				Representative: getCompany.Representative,
			},
			Electricity: electricDTO,
			Total:       int64(len(electricDTO)),
		}
		return electricity, nil
	}
}

// EditElectricity
//
// # 编辑电费
//
// 编辑电费接口, 用于编辑电费. 用户在这里可以进行编辑电费，编辑电费的时候需要用户提供时间参数，编辑成功后会返回电费信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - valley: float64, 谷电费
//   - peak: float64, 峰电费
//   - timer: gtime.Time, 时间
//
// # 返回:
//   - err: error, 错误
func (s *sElectric) EditElectricity(ctx context.Context, valley float64, peak float64, timer gtime.Time) (err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 ElectricEdit | 编辑电费")
	getYearMonth := timer.StartOfMonth().Format("Ym")
	getCompany, err := s.GetCompanyByHeader(ctx)
	if err != nil {
		return err
	}
	// 检查指定月份电费是否存在
	var getElectric *entity.XfCompaniesElectricity
	err = dao.XfCompaniesElectricity.Ctx(ctx).
		Where(do.XfCompaniesElectricity{PeriodAt: getYearMonth, Cods: getCompany.Cods}).
		Scan(&getElectric)
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if getElectric == nil {
		return xerror.NewError(xerror.OperationFailed, "该月电费未创建")
	}
	// 修改后进行重新价格计算
	var getElectricRate *entity.XfElectricityRates
	err = dao.XfElectricityRates.Ctx(ctx).Where(do.XfElectricityRates{PeriodAt: getYearMonth}).Scan(&getElectricRate)
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if getElectricRate == nil {
		return xerror.NewError(xerror.OperationFailed, "该月电价未创建，请联系管理员")
	}
	// 计算电费
	valleyElectricityMoney := valley * getElectricRate.ValleyRate
	peakElectricityMoney := peak * getElectricRate.PeakRate
	// 更新电费
	_, err = dao.XfCompaniesElectricity.Ctx(ctx).
		Where(do.XfCompaniesElectricity{Ceuuid: getElectric.Ceuuid}).
		Update(do.XfCompaniesElectricity{
			ValleyElectricity:     valley,
			ValleyElectricityBill: valleyElectricityMoney,
			PeakElectricity:       peak,
			PeakElectricityBill:   peakElectricityMoney,
			TotalElectricity:      valley + peak,
			TotalBill:             valleyElectricityMoney + peakElectricityMoney,
		})
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	} else {
		return nil
	}
}

// DeleteElectricity
//
// # 删除电费
//
// 删除电费接口, 用于删除电费. 用户在这里可以进行删除电费，删除电费的时候需要用户提供电费的UUID，删除成功后会返回删除成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - CeUUID: string, 电费UUID
//
// # 返回:
//   - err: error, 错误
func (s *sElectric) DeleteElectricity(ctx context.Context, CeUUID string) (err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 ElectricDelete | 删除电费")
	// 获取用户信息
	getCompany, err := s.GetCompanyByHeader(ctx)
	if err != nil {
		return err
	}
	// 检查是否存在
	var getElectric *entity.XfCompaniesElectricity
	err = dao.XfCompaniesElectricity.Ctx(ctx).
		Where(do.XfCompaniesElectricity{Ceuuid: CeUUID, Cods: getCompany.Cods}).
		Scan(&getElectric)
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	}
	if getElectric == nil {
		return xerror.NewError(xerror.OperationFailed, "电费不存在")
	}
	// 删除电费
	_, err = dao.XfCompaniesElectricity.Ctx(ctx).Where(do.XfCompaniesElectricity{Ceuuid: CeUUID}).Delete()
	if err != nil {
		return xerror.NewErrorHasError(xerror.ServerInternalError, err)
	} else {
		return nil
	}
}
