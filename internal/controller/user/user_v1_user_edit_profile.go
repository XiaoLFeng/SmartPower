package user

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/user/v1"
)

// UserEditProfile
//
// # 编辑用户信息
//
// 编辑用户信息. 用户在这里可以进行编辑用户信息，编辑用户信息的时候需要用户提供邮箱, 手机号，编辑成功后会返回编辑成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.UserEditProfileReq, 编辑用户信息请求
//
// # 返回:
//   - res: *v1.UserEditProfileRes, 编辑用户信息返回
//   - err: error, 错误信息
func (c *ControllerV1) UserEditProfile(
	ctx context.Context,
	req *v1.UserEditProfileReq,
) (res *v1.UserEditProfileRes, err error) {
	glog.Noticef(ctx, "[CONTROLLER] 执行 UserEditProfile | 编辑用户信息")
	err = service.User().UserEditProfile(ctx, req.Phone, req.Email)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
