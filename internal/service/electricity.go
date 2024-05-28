// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"SmartPower/internal/model/dto/delectric"
	"SmartPower/internal/model/entity"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	IElectric interface {
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
		CreateElectricity(ctx context.Context, valley float64, peak float64, timePicker gtime.Time) (err error)
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
		GetElectricity(ctx context.Context, timer gtime.Time) (getElectricity *delectric.ElectricCompanyDTO, err error)
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
		GetAllElectricity(ctx context.Context) (electricity *delectric.ElectricAllCompanyDTO, err error)
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
		EditElectricity(ctx context.Context, valley float64, peak float64, timer gtime.Time) (err error)
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
		DeleteElectricity(ctx context.Context, CeUUID string) (err error)
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
		RateAdd(ctx context.Context, valley float64, peak float64, timer gtime.Time) (err error)
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
		RateEdit(ctx context.Context, rateID int64, valley float64, peak float64) (err error)
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
		RateDelete(ctx context.Context, rateID int64) (err error)
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
		RateGet(ctx context.Context) (rate []*entity.XfElectricityRates, err error)
	}
)

var (
	localElectric IElectric
)

func Electric() IElectric {
	if localElectric == nil {
		panic("implement not found for interface IElectric, forgot register?")
	}
	return localElectric
}

func RegisterElectric(i IElectric) {
	localElectric = i
}
