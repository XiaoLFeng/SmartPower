// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// XfRole is the golang structure for table xf_role.
type XfRole struct {
	Ruuid       string      `json:"ruuid"       orm:"ruuid"        ` // 角色表主键
	Name        string      `json:"name"        orm:"name"         ` // 角色名
	DisplayName string      `json:"displayName" orm:"display_name" ` // 展示名称
	Description string      `json:"description" orm:"description"  ` // 描述
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // 修改时间
}
