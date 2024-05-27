package v1

import (
	"SmartPower/internal/model/dto/duser"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthLoginReq
//
// # 用户登录
//
// 用户登录请求, 登录成功后返回用户信息. 传入用户名/邮箱/手机号和密码.
//
// # 参数:
//   - duser: string, 用户名/邮箱/手机号
//   - password: string, 密码
type AuthLoginReq struct {
	g.Meta   `path:"/api/v1/auth/login" method:"Get" summary:"用户登录" tags:"授权控制器"`
	User     string `json:"user" v:"required#请输入用户名/邮箱/手机号" summary:"用户名/邮箱/手机号" example:"admin"`
	Password string `json:"password" v:"required#请输入密码" summary:"密码" example:"admin"`
}

// AuthLoginRes
//
// # 用户登录返回
//
// 用户登录返回, 登录成功后返回用户信息.
type AuthLoginRes struct {
	g.Meta `mime:"application/json" dc:"用户登录返回"`
	duser.UserCurrent
}
