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
		// UserEditCompany
		//
		// # 编辑企业信息
		//
		// 编辑企业信息. 用户在这里可以进行编辑企业信息，编辑企业信息的时候需要用户提供企业编码, 企业名称, 企业地址, 企业法人，编辑成功后会返回编辑成功的信息.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - cods: string, 企业编码
		//   - name: string, 企业名称
		//   - address: string, 企业地址
		//   - principal: string, 企业法人
		//
		// # 返回:
		//   - err: error, 错误信息
		UserEditCompany(ctx context.Context, cods string, name string, address string, principal string) (err error)
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
		//   - err: error, 错误信息
		GetUserCurrent(ctx context.Context) (getUser *duser.UserDetailed, err error)
		// UserEditProfile
		//
		// # 编辑用户信息
		//
		// 编辑用户信息. 用户在这里可以进行编辑用户信息，编辑用户信息的时候需要用户提供邮箱, 手机号，编辑成功后会返回编辑成功的信息.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - phone: string, 手机号
		//   - email: string, 邮箱
		//
		// # 返回:
		//   - err: error, 错误信息
		UserEditProfile(ctx context.Context, phone string, email string) (err error)
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
