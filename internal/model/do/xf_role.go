// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// XfRole is the golang structure of table xf_role for DAO operations like Where/Data.
type XfRole struct {
	g.Meta      `orm:"table:xf_role, do:true"`
	Ruuid       interface{} // 角色表主键
	Name        interface{} // 角色名
	DisplayName interface{} // 展示名称
	Description interface{} // 描述
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
