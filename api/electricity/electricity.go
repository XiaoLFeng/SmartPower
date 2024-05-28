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
	ElectricDeleteRate(ctx context.Context, req *v1.ElectricDeleteRateReq) (res *v1.ElectricDeleteRateRes, err error)
	ElectricRateEdit(ctx context.Context, req *v1.ElectricRateEditReq) (res *v1.ElectricRateEditRes, err error)
}
