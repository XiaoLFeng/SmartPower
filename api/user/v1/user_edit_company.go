package v1

import "github.com/gogf/gf/v2/frame/g"

// UserEditCompanyReq
//
// # 编辑企业信息
//
// 编辑企业信息. 用户在这里可以进行编辑企业信息，编辑企业信息的时候需要用户提供企业编码, 企业名称, 企业地址, 企业法人，编辑成功后会返回编辑成功的信息.
//
// # 参数:
//   - cods: string, 企业编码
//   - name: string, 企业名称
//   - address: string, 企业地址
//   - representative: string, 企业法人
type UserEditCompanyReq struct {
	g.Meta         `path:"/api/v1/user/company" method:"Put" summary:"编辑企业信息" tags:"用户控制器"`
	Cods           string `json:"cods" v:"required#请输入统一社会信用代码" summary:"企业编码" example:"cods"`
	Name           string `json:"name" v:"required#请输入企业名称" summary:"企业名称" example:"name"`
	Address        string `json:"address" v:"required#请输入企业地址" summary:"企业地址" example:"address"`
	Representative string `json:"representative" v:"required#请输入企业法人" summary:"企业法人" example:"representative"`
}

// UserEditCompanyRes
//
// # 编辑企业信息返回
//
// 编辑企业信息返回.
type UserEditCompanyRes struct {
	g.Meta `mime:"application/json" summary:"编辑企业信息返回"`
}
