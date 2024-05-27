package xerror

/*
ECode

# 错误码接口

错误码接口, 用于返回错误码信息.

# 方法:
  - Code: int, 错误码
  - Message: string, 错误信息
  - Data: interface{}, 数据
*/
type ECode interface {
	// Code - 返回一个错误码的具体信息
	Code() int
	// Message - 返回错误码所对应的中文解释大致错误
	Message() string
	// Output - 返回一个英文规范错误说明
	Output() string
}

// ==================================================
// 常见错误代码定义。
// 保留了内部错误代码：代码 100 <= HERE <= 200。
// ==================================================

var (
	Success             = LocalCode{output: "Success", code: 100, message: "操作成功"}
	ServerInternalError = LocalCode{output: "ServerInternalError", code: 101, message: "服务器内部错误"}
	UnknownError        = LocalCode{output: "UnknownError", code: 102, message: "未知错误"}
	PageNotFound        = LocalCode{output: "PageNotFound", code: 110, message: "页面不存在"}
	StatusForbidden     = LocalCode{output: "StatusForbidden", code: 111, message: "禁止访问"}
	UserAlreadyExists   = LocalCode{output: "UserAlreadyExists", code: 112, message: "用户已存在"}
	UserNotExist        = LocalCode{output: "UserNotExist", code: 113, message: "用户不存在"}
)

func BaseLocalCode(output string, code int, message string) ECode {
	return LocalCode{output: output, code: code, message: message}
}
