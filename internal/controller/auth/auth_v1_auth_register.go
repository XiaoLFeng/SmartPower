package auth

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/glog"

	"SmartPower/api/auth/v1"
)

// AuthRegister
//
// # 注册
//
// 注册接口, 用于用户注册. 用户在这里可以进行注册，注册的时候需要用户提供基本的注册信息以及企业的认证信息，注册成功后会返回注册成功的信息.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.AuthRegisterReq, 请求
//
// # 返回:
//   - res: *v1.AuthRegisterRes, 返回
//   - err: error, 错误
func (c *ControllerV1) AuthRegister(ctx context.Context, req *v1.AuthRegisterReq) (res *v1.AuthRegisterRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 AuthRegister | 注册")
	err = service.Auth().UserRegister(ctx, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
