package v1

import (
	"SmartPower/internal/model/dto/duser"
	"github.com/gogf/gf/v2/frame/g"
)

// ConsoleUserAddReq
//
// # 管理员添加用户
//
// 添加用户接口, 用于添加用户. 用户在这里可以进行添加用户，添加用户的时候需要用户提供用户名、邮箱、手机号、公司名称、公司编码、公司地址、法人代表，添加成功后会返回添加成功的信息.
//
// # 参数:
//   - Username: string, 用户名
//   - Email: string, 邮箱
//   - Phone: string, 手机号
//   - Company: string, 公司名称
//   - CompanyCods: string, 公司编码
//   - CompanyAddress: string, 公司地址
//   - Representative: string, 法人代表
//   - IsAdmin: bool, 是否为管理员
type ConsoleUserAddReq struct {
	g.Meta         `path:"/api/v1/console/user" method:"Post" summary:"添加用户" tags:"管理控制器"`
	Username       string `json:"username" v:"required|length:6,30#请输入用户名|账号长度为 6-30 位"`
	Email          string `json:"email" v:"required|email#请输入邮箱|请输入正确的邮箱格式"`
	Phone          string `json:"phone" v:"required|phone#请输入手机号|请输入正确的手机号"`
	Company        string `json:"company"`
	CompanyCods    string `json:"company_cods"`
	CompanyAddress string `json:"company_address"`
	Representative string `json:"representative"`
	IsAdmin        bool   `json:"has_admin" v:"required#请输入是否为管理员" default:"false"`
}

// ConsoleUserAddRes
//
// # 管理员添加用户返回
//
// 添加用户返回.
type ConsoleUserAddRes struct {
	g.Meta `mime:"application/json"`
	duser.DUserForConsoleCreate
}
