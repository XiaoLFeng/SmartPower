package v1

import "github.com/gogf/gf/v2/frame/g"

// AuthChangePasswordReq
//
// # 修改密码
//
// 修改密码接口, 用于用户修改密码. 用户在这里可以进行修改密码，修改密码的时候需要用户提供当前密码以及新密码，修改成功后会返回修改成功的信息.
//
// # 参数:
//   - now_password: string, 当前密码
//   - new_password: string, 新密码
type AuthChangePasswordReq struct {
	g.Meta      `path:"/api/v1/auth/password/change" method:"Put" summary:"修改密码" tags:"授权控制器"`
	NowPassword string `json:"now_password" v:"required#当前密码不能为空" summary:"当前密码"`
	NewPassword string `json:"new_password" v:"required#新密码不能为空" summary:"新密码"`
}

// AuthChangePasswordRes
//
// # 修改密码
//
// 修改密码接口, 用于用户修改密码. 用户在这里可以进行修改密码，修改密码的时候需要用户提供当前密码以及新密码，修改成功后会返回修改成功的信息.
type AuthChangePasswordRes struct {
	g.Meta `mime:"application/json"`
}
