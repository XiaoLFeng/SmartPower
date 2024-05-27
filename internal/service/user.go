// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"SmartPower/internal/model/dto/duser"
	"context"
)

type (
	IUser interface {
		// GetUserCurrent
		//
		// # 获取当前用户信息
		//
		// 获取当前用户信息, 用于获取当前登录用户的信息.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//
		// # 返回:
		//   - getUser: *duser.UserDetailed, 获取当前用户信息返回
		//	 - err: error, 错误信息
		GetUserCurrent(ctx context.Context) (getUser *duser.UserDetailed, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
