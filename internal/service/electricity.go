// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
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
