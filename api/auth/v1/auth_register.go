package v1

import "github.com/gogf/gf/v2/frame/g"

/*
AuthRegisterReq

# 用户注册

用户注册请求, 注册成功后返回用户信息.

# 参数:
  - user: string, 账号
  - password: string, 密码
*/
type AuthRegisterReq struct {
	g.Meta   `path:"/api/v1/auth/login" method:"Post" summary:"用户注册" tags:"授权控制器"`
	User     string `json:"user" v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Password string `json:"password" v:"required|length:6,30#请输入密码|密码长度为:min到:max位"`
}

type AuthRegisterRes struct {
}
