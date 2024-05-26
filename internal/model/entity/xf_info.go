// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// XfInfo is the golang structure for table xf_info.
type XfInfo struct {
	SystemId  string      `json:"systemId"  orm:"system_id"  ` // 系统主键
	Keyword   string      `json:"keyword"   orm:"keyword"    ` // 键
	Value     string      `json:"value"     orm:"value"      ` // 值
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
}
