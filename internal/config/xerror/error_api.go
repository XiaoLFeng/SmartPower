package xerror

type IError interface {
	getErrorInfo() error
	getECode() *ECode
}
