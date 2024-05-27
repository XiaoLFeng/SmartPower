// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// XfUser is the golang structure of table xf_user for DAO operations like Where/Data.
type XfUser struct {
	g.Meta    `orm:"table:xf_user, do:true"`
	Uuid      interface{} // 用户 id
	Username  interface{} // 用户名
	Email     interface{} // 邮箱
	Phone     interface{} // 手机号
	Password  interface{} // 用户密码
	Role      interface{} // 角色主键
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
