package user

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/user/v1"
)

// UserCurrent
//
// # 获取当前用户信息
//
// 获取当前用户信息, 用于获取当前登录用户的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//
// # 返回:
//   - res: *v1.UserCurrentRes, 获取当前用户信息返回
//   - err: error, 错误信息
func (c *ControllerV1) UserCurrent(ctx context.Context, _ *v1.UserCurrentReq) (res *v1.UserCurrentRes, err error) {
	glog.Noticef(ctx, "[CONTROLLER] 执行 UserCurrent | 获取当前用户信息")
	// 获取请求头
	getUser, err := service.User().GetUserCurrent(ctx)
	if err != nil {
		return nil, err
	} else {
		res = &v1.UserCurrentRes{
			UserDetailed: *getUser,
		}
		return res, nil
	}
}
