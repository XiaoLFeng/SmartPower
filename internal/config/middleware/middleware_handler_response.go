package middleware

import (
	"SmartPower/internal/config/xerror"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

/*
BaseResponse

# 基础响应

基础响应结构体, 用于返回基础响应信息.

# 字段:
  - output: string, 英文输出
  - code: int, 状态码
  - message: string, 中文描述
  - error_message: string, 错误信息(可选)
  - data: interface{}, 数据
*/
type BaseResponse struct {
	Output       string      `json:"output" dc:"英文输出"`
	Code         int         `json:"code" dc:"状态码"`
	Message      string      `json:"message" dc:"中文描述"`
	ErrorMessage string      `json:"error_message,omitempty" dc:"错误信息"`
	Data         interface{} `json:"data,omitempty" dc:"数据"`
}

func HandlerResponseMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

	// 自定义缓冲内容，退出默认信息返回
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = *xerror.GetECode(err)
	)
	if r.GetError() != nil {
		if gerror.Code(err) == gcode.CodeNil {
			code = xerror.ServerInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = xerror.PageNotFound
			case http.StatusForbidden:
				code = xerror.StatusForbidden
			default:
				code = xerror.UnknownError
			}
			err = xerror.NewError(code, msg)
			r.SetError(err)
		} else {
			code = xerror.Success
		}
	}

	r.Response.WriteJson(BaseResponse{
		Output:       code.Output(),
		Code:         code.Code(),
		Message:      code.Message(),
		ErrorMessage: msg,
		Data:         res,
	})
}
