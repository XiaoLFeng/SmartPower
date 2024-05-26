// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// XfInfo is the golang structure of table xf_info for DAO operations like Where/Data.
type XfInfo struct {
	g.Meta    `orm:"table:xf_info, do:true"`
	SystemId  interface{} // 系统主键
	Keyword   interface{} // 键
	Value     interface{} // 值
	UpdatedAt *gtime.Time // 更新时间
}
