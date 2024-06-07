package v1

import "github.com/gogf/gf/v2/frame/g"

// ConsoleUserEditReq
//
// # 编辑用户
//
// 编辑用户接口, 用于编辑用户. 用户在这里可以进行编辑用户，编辑用户的时候需要用户提供用户UUID，编辑成功后会返回编辑成功的信息.
//
// # 参数:
//   - UUID: string, 用户UUID
//   - Username: string, 用户名
//   - Email: string, 邮箱
//   - Phone: string, 手机号
//   - Company: string, 公司名称
//   - CompanyCods: string, 公司编码
//   - CompanyAddress: string, 公司地址
//   - Representative: string, 法人代表
//   - IsAdmin: bool, 是否为管理员
type ConsoleUserEditReq struct {
	g.Meta         `path:"/api/v1/console/user" method:"Put" summary:"编辑用户" tags:"管理控制器"`
	UUID           string `json:"uuid" v:"required|regex:^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$#请输入用户UUID|请输入正确的用户UUID"`
	Username       string `json:"username" v:"required|length:6,30#请输入用户名|账号长度为 6-30 位"`
	Email          string `json:"email" v:"required|email#请输入邮箱|请输入正确的邮箱格式"`
	Phone          string `json:"phone" v:"required|phone#请输入手机号|请输入正确的手机号"`
	Company        string `json:"company"`
	CompanyCods    string `json:"company_cods"`
	CompanyAddress string `json:"company_address"`
	Representative string `json:"representative"`
	IsAdmin        bool   `json:"admin" v:"required#请输入是否为管理员"`
}

// ConsoleUserEditRes
//
// # 编辑用户返回
//
// 编辑用户返回.
type ConsoleUserEditRes struct {
	g.Meta `mime:"application/json"`
}
