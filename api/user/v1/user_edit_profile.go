package v1

import "github.com/gogf/gf/v2/frame/g"

// UserEditProfileReq
//
// # 编辑用户信息
//
// 编辑用户信息. 用户在这里可以进行编辑用户信息，编辑用户信息的时候需要用户提供邮箱, 手机号，编辑成功后会返回编辑成功的信息.
//
// # 参数:
//   - Email: string, 邮箱
//   - Phone: string, 手机号
type UserEditProfileReq struct {
	g.Meta `path:"/api/v1/user/profile" method:"Put" summary:"编辑用户信息" tags:"用户控制器"`
	Email  string `json:"email" v:"required|email#请输入邮箱|请输入正确的邮箱" summary:"邮箱" example:"gm@x-lf.cn"`
	Phone  string `json:"phone" v:"required|phone#请输入手机号|请输入正确的手机号" summary:"手机号" example:"18888888888"`
}

// UserEditProfileRes
//
// # 编辑用户信息返回
//
// 编辑用户信息返回.
type UserEditProfileRes struct {
	g.Meta `mime:"application/json" summary:"编辑用户信息返回"`
}
