// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package electricity

import (
	"context"

	"SmartPower/api/electricity/v1"
)

type IElectricityV1 interface {
	ElectricCreate(ctx context.Context, req *v1.ElectricCreateReq) (res *v1.ElectricCreateRes, err error)
}