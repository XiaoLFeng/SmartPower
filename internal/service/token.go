// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"SmartPower/internal/model/entity"
	"context"
)

type (
	IToken interface {
		// CheckTokenExist
		//
		// # 检查令牌是否存在
		//
		// 检查令牌是否存在, 传入令牌进行检查, 若令牌存在则返回错误信息. 对输入的令牌进行检查，只允许搜索挡枪用户存在的令牌. 不允许跨用户查询其他用户
		// 的令牌相关的信息. 若令牌不存在则返回对应的报错信息，若令牌存在且为正确的内容则不进行任何报错的抛出.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - token: string, 令牌
		//
		// # 返回:
		//   - err: error, 错误
		CheckTokenExist(ctx context.Context, token string) (err error)
		// GetUserByToken
		//
		// # 根据令牌获取用户信息
		//
		// 根据令牌获取用户信息, 传入令牌进行检查, 若令牌存在则返回用户信息. 对输入的令牌进行检查，只允许搜索挡枪用户存在的令牌. 不允许跨用户查询其他用户
		// 的令牌相关的信息. 若令牌不存在则返回对应的报错信息，若令牌存在且为正确的内容则返回用户信息.
		//
		// # 参数:
		//   - ctx: context.Context, 上下文
		//   - token: string, 令牌
		//
		// # 返回:
		//   - user: *entity.XfUser, 用户信息
		//   - err: error, 错误s
		GetUserByToken(ctx context.Context, token string) (user *entity.XfUser, err error)
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
