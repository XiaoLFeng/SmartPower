package console

import (
	"SmartPower/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"SmartPower/api/console/v1"
)

// ConsoleUserGetList
//
// # 获取用户列表
//
// 获取用户列表接口, 用于获取用户列表. 用户在这里可以进行获取用户列表，获取用户列表的时候需要用户提供用户列表的信息，获取成功后会返回获取成功的信息.
//
// # 参数:
//   - ctx: context.Context
//
// # 返回:
//   - *v1.ConsoleUserGetListRes
//   - error
func (c *ControllerV1) ConsoleUserGetList(ctx context.Context, _ *v1.ConsoleUserGetListReq) (res *v1.ConsoleUserGetListRes, err error) {
	g.Log().Noticef(ctx, "[CONTROL] 执行 ConsoleUserGetList | 获取用户列表")
	getUserDetail, err := service.Console().UserGetList(ctx)
	if err != nil {
		return nil, err
	} else {
		return &v1.ConsoleUserGetListRes{
			List: getUserDetail,
		}, nil
	}
}
