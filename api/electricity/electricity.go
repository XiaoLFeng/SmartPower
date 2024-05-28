// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package electricity

import (
	"context"

	"SmartPower/api/electricity/v1"
)

type IElectricityV1 interface {
	ElectricAddRate(ctx context.Context, req *v1.ElectricAddRateReq) (res *v1.ElectricAddRateRes, err error)
	ElectricCreate(ctx context.Context, req *v1.ElectricCreateReq) (res *v1.ElectricCreateRes, err error)
	ElectricDelete(ctx context.Context, req *v1.ElectricDeleteReq) (res *v1.ElectricDeleteRes, err error)
	ElectricDeleteRate(ctx context.Context, req *v1.ElectricDeleteRateReq) (res *v1.ElectricDeleteRateRes, err error)
	ElectricEdit(ctx context.Context, req *v1.ElectricEditReq) (res *v1.ElectricEditRes, err error)
	ElectricRateEdit(ctx context.Context, req *v1.ElectricRateEditReq) (res *v1.ElectricRateEditRes, err error)
	ElectricGet(ctx context.Context, req *v1.ElectricGetReq) (res *v1.ElectricGetRes, err error)
	ElectricGetAll(ctx context.Context, req *v1.ElectricGetAllReq) (res *v1.ElectricGetAllRes, err error)
	ElectricGetRate(ctx context.Context, req *v1.ElectricGetRateReq) (res *v1.ElectricGetRateRes, err error)
	ElectricRegionCalculate(ctx context.Context, req *v1.ElectricRegionCalculateReq) (res *v1.ElectricRegionCalculateRes, err error)
}
