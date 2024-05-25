// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// XfUser is the golang structure for table xf_user.
type XfUser struct {
	Uuid      string      `json:"uuid"      orm:"uuid"       ` // 用户 id
	Username  string      `json:"username"  orm:"username"   ` // 用户名
	Email     string      `json:"email"     orm:"email"      ` // 邮箱
	Phone     string      `json:"phone"     orm:"phone"      ` // 手机号
	Password  string      `json:"password"  orm:"password"   ` // 用户密码
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 修改时间
}
