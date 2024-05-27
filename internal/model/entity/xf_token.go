// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// XfToken is the golang structure for table xf_token.
type XfToken struct {
	Tuuid     string      `json:"tuuid"     orm:"tuuid"      ` // 令牌主键
	Uuid      string      `json:"uuid"      orm:"uuid"       ` // 用户主键
	Token     string      `json:"token"     orm:"token"      ` // 授权令牌
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	ExpiredAt *gtime.Time `json:"expiredAt" orm:"expired_at" ` // 过期时间
}
