package console

import (
	"SmartPower/api/console/v1"
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// ConsoleUserAdd
//
// # 添加控制台用户
//
// 添加控制台用户, 用于添加控制台用户. 传入用户名, 密码, 角色.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ConsoleUserAddReq, 添加控制台用户请求
//
// # 返回:
//   - res: *v1.ConsoleUserAddRes, 添加控制台用户返回
//   - err: error, 错误信息
func (c *ControllerV1) ConsoleUserAdd(ctx context.Context, req *v1.ConsoleUserAddReq) (res *v1.ConsoleUserAddRes, err error) {
	g.Log().Noticef(ctx, "[CONTROL] 执行 ConsoleUserAdd | 添加控制台用户")
	// 获取用户进行注册
	getUser, err := service.Console().UserAdd(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return &v1.ConsoleUserAddRes{
			DUserForConsoleCreate: *getUser,
		}, nil
	}
}
