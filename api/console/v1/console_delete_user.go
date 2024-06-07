package v1

import "github.com/gogf/gf/v2/frame/g"

// ConsoleUserDeleteReq
//
// # 删除用户
//
// 删除用户接口, 用于删除用户. 用户在这里可以进行删除用户，删除用户的时候需要用户提供用户UUID，删除成功后会返回删除成功的信息.
//
// # 参数:
//   - UUID: string, 用户UUID
type ConsoleUserDeleteReq struct {
	g.Meta `path:"/api/v1/console/user" method:"Delete" summary:"删除用户" tags:"管理控制器"`
	UUID   string `json:"uuid" v:"required|regex:^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$#请输入用户UUID|请输入正确的用户UUID"`
}

// ConsoleUserDeleteRes
//
// # 删除用户返回
//
// 删除用户返回.
type ConsoleUserDeleteRes struct {
	g.Meta `mime:"application/json"`
}
