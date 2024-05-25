// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// XfCompanies is the golang structure for table xf_companies.
type XfCompanies struct {
	Cods           string      `json:"cods"           orm:"cods"           ` // 企业唯一信用社会代码
	Name           string      `json:"name"           orm:"name"           ` // 企业名字
	Address        string      `json:"address"        orm:"address"        ` // 企业所在地址
	Representative string      `json:"representative" orm:"representative" ` // 法定代表人
	Uuid           string      `json:"uuid"           orm:"uuid"           ` // 指定用户
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"     ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"     ` // 更新时间
}
