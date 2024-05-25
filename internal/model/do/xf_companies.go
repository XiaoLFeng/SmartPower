// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// XfCompanies is the golang structure of table xf_companies for DAO operations like Where/Data.
type XfCompanies struct {
	g.Meta         `orm:"table:xf_companies, do:true"`
	Cods           interface{} // 企业唯一信用社会代码
	Name           interface{} // 企业名字
	Address        interface{} // 企业所在地址
	Representative interface{} // 法定代表人
	Uuid           interface{} // 指定用户
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
