package v1

import (
	"SmartPower/internal/model/dto/duser"
	"github.com/gogf/gf/v2/frame/g"
)

// UserCurrentReq
//
// # 获取当前用户信息
//
// 获取当前用户信息, 用于获取当前登录用户的信息.
type UserCurrentReq struct {
	g.Meta `path:"/api/v1/user/current" method:"Get" summary:"获取当前用户信息" tags:"用户控制器"`
}

// UserCurrentRes
//
// # 获取当前用户信息返回
//
// 获取当前用户信息返回.
type UserCurrentRes struct {
	g.Meta             `mime:"application/json" dc:"获取当前用户信息返回"`
	duser.UserDetailed `summary:"用户信息返回"`
}
