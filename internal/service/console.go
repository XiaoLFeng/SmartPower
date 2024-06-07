// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "SmartPower/api/console/v1"
	"SmartPower/internal/model/dto/duser"
	"context"
)

type (
	IConsole interface {
		// UserAdd
		//
		// # 添加控制台用户
		//
		// 添加控制台用户, 用于添加控制台用户. 传入用户名, 密码, 角色.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - getData: *v1.ConsoleUserAddReq, 添加控制台用户请求
		//
		// # 返回:
		//   - err: error, 错误
		UserAdd(ctx context.Context, getData *v1.ConsoleUserAddReq) (user *duser.DUserForConsoleCreate, err error)
		// UserResetPassword
		//
		// # 重置控制台用户密码
		//
		// 重置控制台用户密码, 用于重置控制台用户密码. 传入用户ID.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - req: v1.ConsoleUserResetPasswordReq, 重置控制台用户密码请求
		//
		// # 返回:
		//   - err: error, 错误
		UserResetPassword(ctx context.Context, req *v1.ConsoleUserResetPasswordReq) (User *duser.DUserForConsoleCreate, err error)
		// UserDelete
		//
		// # 删除控制台用户
		//
		// 删除控制台用户, 用于删除控制台用户. 传入用户ID.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - getUUID: string, 用户ID
		//
		// # 返回:
		//   - err: error, 错误
		UserDelete(ctx context.Context, getUUID string) (err error)
		// UserEdit
		//
		// # 编辑控制台用户
		//
		// 编辑控制台用户, 用于编辑控制台用户. 传入用户ID, 用户名, 邮箱, 手机号, 公司名称, 公司编码, 公司地址, 法人代表, 是否为管理员.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - getData: *v1.ConsoleUserEditReq, 编辑控制台用户请求
		//
		// # 返回:
		//   - err: error, 错误
		UserEdit(ctx context.Context, getData *v1.ConsoleUserEditReq) (err error)
		UserGetList(ctx context.Context) (userDetail []*duser.UserDetailed, err error)
	}
)

var (
	localConsole IConsole
)

func Console() IConsole {
	if localConsole == nil {
		panic("implement not found for interface IConsole, forgot register?")
	}
	return localConsole
}

func RegisterConsole(i IConsole) {
	localConsole = i
}
