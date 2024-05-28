// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"SmartPower/api/user/v1"
)

type IUserV1 interface {
	UserCurrent(ctx context.Context, req *v1.UserCurrentReq) (res *v1.UserCurrentRes, err error)
	UserEditCompany(ctx context.Context, req *v1.UserEditCompanyReq) (res *v1.UserEditCompanyRes, err error)
	UserEditProfile(ctx context.Context, req *v1.UserEditProfileReq) (res *v1.UserEditProfileRes, err error)
}
