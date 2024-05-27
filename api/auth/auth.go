// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"SmartPower/api/auth/v1"
)

type IAuthV1 interface {
	AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error)
	AuthChangePassword(ctx context.Context, req *v1.AuthChangePasswordReq) (res *v1.AuthChangePasswordRes, err error)
	AuthRegister(ctx context.Context, req *v1.AuthRegisterReq) (res *v1.AuthRegisterRes, err error)
}
