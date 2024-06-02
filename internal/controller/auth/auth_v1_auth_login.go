package auth

import (
	"SmartPower/api/auth/v1"
	"SmartPower/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

// AuthLogin
//
// # 用户登录
//
// 用户登录请求, 登录成功后返回用户信息. 传入用户名/邮箱/手机号和密码.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.AuthLoginReq, 请求
//
// # 返回:
//   - res: *v1.AuthLoginRes, 返回
//   - err: error, 错误
func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	glog.Noticef(ctx, "[CONTROL] 执行 AuthLogin | 用户登录")
	// 根据 ctx 获取 Request
	getRequest := g.RequestFromCtx(ctx)
	// 检查用户是否已登陆
	err = service.Token().CheckTokenExist(ctx, getRequest.Header.Get("Authorization"))
	if err != nil {
		data, err := service.Auth().AuthLogin(ctx, req.User, req.Password)
		if err != nil {
			return nil, err
		} else {
			res = &v1.AuthLoginRes{
				UserCurrent: *data,
			}
			return res, nil
		}
	} else {
		return nil, berror.NewError(bcode.OperationFailed, "用户已登录")
	}
}
