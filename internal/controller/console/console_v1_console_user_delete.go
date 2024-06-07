package console

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"SmartPower/api/console/v1"
)

// ConsoleUserDelete
//
// # 删除控制台用户
//
// 删除控制台用户, 用于删除控制台用户. 传入用户ID.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ConsoleUserDeleteReq, 删除控制台用户请求
//
// # 返回:
//   - res: *v1.ConsoleUserDeleteRes, 删除控制台用户返回
//   - err: error, 错误信息
func (c *ControllerV1) ConsoleUserDelete(ctx context.Context, req *v1.ConsoleUserDeleteReq) (res *v1.ConsoleUserDeleteRes, err error) {
	g.Log().Noticef(ctx, "[CONTROL] 执行 ConsoleUserDelete | 删除控制台用户")
	// 获取用户进行删除
	err = service.Console().UserDelete(ctx, req.UUID)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
