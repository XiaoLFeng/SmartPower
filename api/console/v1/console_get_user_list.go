package v1

import (
	"SmartPower/internal/model/dto/duser"
	"github.com/gogf/gf/v2/frame/g"
)

// ConsoleUserGetListReq
//
// # 获取用户列表
//
// 获取用户列表接口, 用于获取用户列表. 用户在这里可以进行获取用户列表，获取用户列表的时候需要用户提供用户列表的信息，获取成功后会返回获取成功的信息.
type ConsoleUserGetListReq struct {
	g.Meta `path:"/api/v1/console/user" method:"Get" summary:"获取用户列表" tags:"管理控制器"`
}

// ConsoleUserGetListRes
//
// # 获取用户列表返回
//
// 获取用户列表返回.
type ConsoleUserGetListRes struct {
	g.Meta `mime:"application/json"`
	List   []*duser.UserDetailed `json:"list"`
}
