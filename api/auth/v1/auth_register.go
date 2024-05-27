package v1

import "github.com/gogf/gf/v2/frame/g"

// AuthRegisterReq
//
// # 用户注册
//
// 用户注册请求, 注册成功后返回用户信息.
//
// # 参数:
//   - duser: string, 账号
//   - password: string, 密码
type AuthRegisterReq struct {
	g.Meta         `path:"/api/v1/auth/login" method:"Post" summary:"用户注册" tags:"授权控制器"`
	Username       string `json:"username" v:"required|length:6,30#请输入用户名|账号长度为:min到:max位"`
	Email          string `json:"email" v:"required|email#请输入邮箱|请输入正确的邮箱格式"`
	Phone          string `json:"phone" v:"required|phone#请输入手机号|请输入正确的手机号"`
	Password       string `json:"password" v:"required|length:6,30#请输入密码|密码长度为:min到:max位"`
	Company        string `json:"company" v:"required#请输入公司名称"`
	CompanyCods    string `json:"company_cods" v:"required#请输入公司编码"`
	CompanyAddress string `json:"company_address" v:"required#请输入公司地址"`
	Representative string `json:"representative" v:"required#请输入法人代表"`
}

// AuthRegisterRes
//
// # 用户注册返回
//
// 用户注册返回, 注册成功后返回用户信息.
type AuthRegisterRes struct {
	g.Meta `mime:"application/json" dc:"用户注册返回"`
}
