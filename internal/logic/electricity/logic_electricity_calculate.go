package electricity

import (
	"SmartPower/internal/dao"
	"SmartPower/internal/model/do"
	"SmartPower/internal/model/dto/dcompany"
	"SmartPower/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

// RegionCalculate
//
// # 电费区域计算
//
// 电费区域计算, 用于计算电费. 传入区域, 开始时间, 结束时间.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - region: string, 区域
//   - startTime: gtime.Time, 开始时间
//   - endTime: gtime.Time, 结束时间
//
// # 返回:
//   - err: error, 错误
func (s *sElectric) RegionCalculate(
	ctx context.Context,
	region string,
	startTime gtime.Time,
	endTime gtime.Time,
) (calculate *dcompany.CompanyRegionDTO, err error) {
	glog.Noticef(ctx, "[LOGIC] 执行 RegionCalculate | 电费区域计算")
	// 根据地域搜索企业
	var getCompanies []*entity.XfCompanies
	err = dao.XfCompanies.Ctx(ctx).
		WhereLike("address", "%"+region+"%").
		WhereOrLike("name", "%"+region+"%").
		Scan(&getCompanies)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if len(getCompanies) == 0 {
		return nil, berror.NewError(bcode.NotExist, "该地区没有企业")
	}
	// 计算电费
	var regionDTO *dcompany.CompanyRegionDTO
	var calculateDTO []*dcompany.CompanyCalculateDTO

	var maxTotalValleyElectricity float64 = 0
	var maxTotalPeakElectricity float64 = 0
	var maxTotalValleyBill float64 = 0
	var maxTotalPeakBill float64 = 0

	for _, company := range getCompanies {
		var getElectricity []*entity.XfCompaniesElectricity
		err := dao.XfCompaniesElectricity.Ctx(ctx).
			Where(do.XfCompaniesElectricity{Cods: company.Cods}).
			Where("period_at >= ?", startTime.StartOfMonth().Format("Ym")).
			Where("period_at <= ?", endTime.EndOfMonth().Format("Ym")).
			Scan(&getElectricity)
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
		var totalValleyElectricity float64 = 0
		var totalPeakElectricity float64 = 0
		var totalValleyBill float64 = 0
		var totalPeakBill float64 = 0
		var rateDTO []*dcompany.ElectricCompanyEntityDTO
		if len(getElectricity) != 0 {
			// 循环处理
			for _, electricity := range getElectricity {
				totalValleyElectricity += electricity.ValleyElectricity
				totalPeakElectricity += electricity.PeakElectricity
				totalValleyBill += electricity.ValleyElectricityBill
				totalPeakBill += electricity.PeakElectricityBill
				rateDTO = append(rateDTO, &dcompany.ElectricCompanyEntityDTO{
					CeUUID:                electricity.Ceuuid,
					PeriodAt:              electricity.PeriodAt,
					ValleyElectricity:     electricity.ValleyElectricity,
					ValleyElectricityBill: electricity.ValleyElectricityBill,
					PeakElectricity:       electricity.PeakElectricity,
					PeakElectricityBill:   electricity.PeakElectricityBill,
					TotalElectricity:      electricity.TotalElectricity,
					TotalBill:             electricity.TotalBill,
					CreatedAt:             electricity.CreatedAt,
					UpdatedAt:             electricity.UpdatedAt,
				})
			}
		}
		// 数据总结
		calculateDTO = append(calculateDTO, &dcompany.CompanyCalculateDTO{
			Company: dcompany.DCompany{
				Cods:           company.Cods,
				Name:           company.Name,
				Address:        company.Address,
				Representative: company.Representative,
			},
			TotalValleyElectricity: totalValleyElectricity,
			TotalPeakElectricity:   totalPeakElectricity,
			TotalElectricity:       totalValleyElectricity + totalPeakElectricity,
			TotalValleyBill:        totalValleyBill,
			TotalPeakBill:          totalPeakBill,
			TotalBill:              totalValleyBill + totalPeakBill,
			Rate:                   rateDTO,
		})

		// 计算最大值
		maxTotalPeakBill += totalPeakBill
		maxTotalValleyBill += totalValleyBill
		maxTotalPeakElectricity += totalPeakElectricity
		maxTotalValleyElectricity += totalValleyElectricity
	}

	// 计算总电费
	regionDTO = &dcompany.CompanyRegionDTO{
		Region:                 region,
		StartTime:              gtime.NewFromTime(startTime.StartOfMonth().Time),
		EndTime:                gtime.NewFromTime(endTime.EndOfMonth().Time),
		TotalValleyElectricity: maxTotalValleyElectricity,
		TotalPeakElectricity:   maxTotalPeakElectricity,
		TotalElectricity:       maxTotalValleyElectricity + maxTotalPeakElectricity,
		TotalValleyBill:        maxTotalValleyBill,
		TotalPeakBill:          maxTotalPeakBill,
		TotalBill:              maxTotalValleyBill + maxTotalPeakBill,
		Company:                calculateDTO,
	}
	return regionDTO, nil
}
