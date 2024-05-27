package auth

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/auth/v1"
)

// AuthChangePassword
//
// # 修改密码
//
// 修改密码接口, 用于用户修改密码. 用户在这里可以进行修改密码，修改密码的时候需要用户提供当前密码以及新密码，修改成功后会返回修改成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.AuthChangePasswordReq, 请求
//
// # 返回:
//   - res: *v1.AuthChangePasswordRes, 返回
//   - err: error, 错误
func (c *ControllerV1) AuthChangePassword(ctx context.Context, req *v1.AuthChangePasswordReq) (res *v1.AuthChangePasswordRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 AuthChangePassword | 修改密码")
	// 获取请求信息
	getRequest := g.RequestFromCtx(ctx)
	// 检查用户的登录信息
	getUser, err := service.Token().GetUserByToken(ctx, getRequest.Header.Get("Authorization"))
	if err == nil {
		err = service.Auth().ChangePassword(ctx, *getUser, req.NowPassword, req.NewPassword)
		if err != nil {
			return nil, err
		} else {
			return nil, nil
		}
	} else {
		return nil, err
	}
}
