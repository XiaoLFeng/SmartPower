package xerror

import "strings"

// XError
//
// # 报错结构体
//
// 用于处理自定义报错的输出内容信息
type XError struct {
	error error  // 具体的报错信息
	text  string // 创建错误时的自定义错误文本，当其代码不为零时可能为空。
	code  ECode  // 如有必要，请提供错误代码。
}

// Error - 自定义报错处理
func (e *XError) Error() string {
	if e == nil {
		return ""
	}
	errStr := e.text
	if errStr == "" && e.code != nil {
		errStr = e.code.Message()
	}
	if e.error != nil {
		if errStr != "" {
			errStr += ": "
		}
		errStr += e.error.Error()
	}
	return errStr
}

// GetECode - 获取详细错误信息
func GetECode(err error) *ECode {
	if err != nil {
		if e, ok := err.(IError); ok {
			return e.getECode()
		} else {
			return nil
		}
	} else {
		return nil
	}
}

// NewError - 创建新的错误
func NewError(code ECode, text ...string) error {
	return &XError{
		text: strings.Join(text, ","),
		code: code,
	}
}

// NewErrorHasError - 创建新的错误并且携带 error 信息
func NewErrorHasError(code ECode, errorData error, text ...string) error {
	return &XError{
		error: errorData,
		text:  strings.Join(text, ","),
		code:  code,
	}
}
