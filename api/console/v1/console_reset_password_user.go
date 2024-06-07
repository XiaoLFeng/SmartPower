package v1

import (
	"SmartPower/internal/model/dto/duser"
	"github.com/gogf/gf/v2/frame/g"
)

// ConsoleUserResetPasswordReq
//
// # 重置用户密码
//
// 重置用户密码接口, 用于重置用户密码. 用户在这里可以进行重置用户密码，重置用户密码的时候需要用户提供用户UUID，重置成功后会返回重置成功的信息.
//
// # 参数:
//   - UUID: string, 用户UUID
type ConsoleUserResetPasswordReq struct {
	g.Meta `path:"/api/v1/console/password/reset" method:"Post" summary:"重置用户密码" tags:"管理控制器"`
	UUID   string `json:"uuid" v:"required|regex:^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$#请输入用户UUID|请输入正确的用户UUID"`
}

// ConsoleUserResetPasswordRes
//
// # 重置用户密码返回
//
// 重置用户密码返回.
type ConsoleUserResetPasswordRes struct {
	g.Meta `mime:"application/json"`
	duser.DUserForConsoleCreate
}
