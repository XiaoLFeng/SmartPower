// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "SmartPower/api/auth/v1"
	"SmartPower/internal/model/dto/duser"
	"context"
)

type (
	IAuth interface {
		// AuthLogin
		//
		// # 用户登录逻辑
		//
		// 用户登录逻辑, 登录成功后返回用户信息.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - duser: string, 用户名/邮箱/手机号
		//   - password: string, 密码
		//
		// # 返回:
		//   - err: error, 错误
		AuthLogin(ctx context.Context, user string, password string) (dUser *duser.UserCurrent, err error)
		// UserRegister
		//
		// # 用户注册
		//
		// 用户注册接口, 用于用户注册. 用户在这里可以进行注册，注册的时候需要用户提供基本注册信息以及企业的认证信息，注册成功后会返回注册成功的信息.
		// 用户的注册需要用户提供必要注册相关的信息，注册成功后会返回注册成功的信息. 若注册失败将会抛出指定的错误信息.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - req: *v1.AuthRegisterReq, 请求
		//
		// # 返回:
		//   - err: error, 错误
		UserRegister(ctx context.Context, req *v1.AuthRegisterReq) (err error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
