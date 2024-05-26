package xerror

// LocalCode
//
// # 本地错误码
//
// 本地错误码, 用于返回错误码信息.
//
// # 属性:
//   - output: string, 英文输出
//   - code: int, 状态码
//   - message: string, 中文描述
type LocalCode struct {
	output  string
	code    int
	message string
}

// Code
//
// # 错误码
//
// 错误码, 用于返回错误码信息.
func (lc LocalCode) Code() int {
	return lc.code
}

// Message
//
// # 错误信息
//
// 错误信息, 用于返回错误信息.
func (lc LocalCode) Message() string {
	return lc.message
}

// Output
//
// # 英文输出
//
// 英文输出, 用于返回英文输出.
func (lc LocalCode) Output() string {
	return lc.output
}
