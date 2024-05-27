package xerror

type IError interface {
	getErrorInfo() error
	getECode() *ECode
}

func (e *XError) getErrorInfo() error {
	return e.error
}

func (e *XError) getECode() *ECode {
	return &e.code
}
