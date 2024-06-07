package console

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"SmartPower/api/console/v1"
)

// ConsoleUserResetPassword
//
// # 重置控制台用户密码
//
// 重置控制台用户密码, 用于重置控制台用户密码. 传入用户ID.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ConsoleUserResetPasswordReq, 重置控制台用户密码请求
//
// # 返回:
//   - res: *v1.ConsoleUserResetPasswordRes, 重置控制台用户密码返回
//   - err: error, 错误信息
func (c *ControllerV1) ConsoleUserResetPassword(ctx context.Context, req *v1.ConsoleUserResetPasswordReq) (res *v1.ConsoleUserResetPasswordRes, err error) {
	g.Log().Noticef(ctx, "[CONTROL] 执行 ConsoleUserResetPassword | 重置控制台用户密码")
	getUser, err := service.Console().UserResetPassword(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return &v1.ConsoleUserResetPasswordRes{
			DUserForConsoleCreate: *getUser,
		}, nil
	}
}
