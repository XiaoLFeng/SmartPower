package console

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"SmartPower/api/console/v1"
)

// ConsoleUserEdit
//
// # 编辑控制台用户
//
// 编辑控制台用户, 用于编辑控制台用户. 传入用户ID, 用户名, 邮箱, 手机号, 公司名称, 公司编码, 公司地址, 法人代表, 是否为管理员.
//
// # 参数:
//   - ctx: context.Context, 上下文
//   - req: *v1.ConsoleUserEditReq, 编辑控制台用户请求
//
// # 返回:
//   - res: *v1.ConsoleUserEditRes, 编辑控制台用户返回
//   - err: error, 错误信息
func (c *ControllerV1) ConsoleUserEdit(ctx context.Context, req *v1.ConsoleUserEditReq) (res *v1.ConsoleUserEditRes, err error) {
	g.Log().Noticef(ctx, "[CONTROL] 执行 ConsoleUserEdit | 编辑控制台用户")
	err = service.Console().UserEdit(ctx, req)
	if err != nil {
		return nil, err
	} else {
		return nil, nil
	}
}
