// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package console

import (
	"context"

	"SmartPower/api/console/v1"
)

type IConsoleV1 interface {
	ConsoleUserAdd(ctx context.Context, req *v1.ConsoleUserAddReq) (res *v1.ConsoleUserAddRes, err error)
	ConsoleUserDelete(ctx context.Context, req *v1.ConsoleUserDeleteReq) (res *v1.ConsoleUserDeleteRes, err error)
	ConsoleUserEdit(ctx context.Context, req *v1.ConsoleUserEditReq) (res *v1.ConsoleUserEditRes, err error)
	ConsoleUserResetPassword(ctx context.Context, req *v1.ConsoleUserResetPasswordReq) (res *v1.ConsoleUserResetPasswordRes, err error)
}
