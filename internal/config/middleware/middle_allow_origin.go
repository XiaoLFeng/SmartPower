package middleware

import "github.com/gogf/gf/v2/net/ghttp"

// MiddleAllowOrigin
//
// # 允许跨域
//
// 允许跨域, 用于允许跨域请求.
//
// # 参数:
//   - r: *ghttp.Request, 请求
func MiddleAllowOrigin(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
