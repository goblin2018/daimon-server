package e

import "errors"

// 将 error 转化为 Error
func CastError(err error) Error {
	if err == nil {
		return OK
	}
	er := new(customError)
	if errors.As(err, &er) {
		return er
	}
	return Unknown.Add(err.Error())
}

func IsError(err error, err1 Error) bool {
	cE := CastError(err)
	return cE.Code() == err1.Code()
}

type Error interface {
	Error() string
	Code() int32
	// 从 error 创建
	From(error) Error
	// 添加 error msg
	Add(string) Error
	// 修改 error msg
	Set(string) Error
}
type customError struct {
	code int32
	msg  string
}

func (e customError) Error() string {
	return e.msg
}
func (e customError) Code() int32 {
	return e.code
}
func (e customError) From(err error) Error {
	if err == nil {
		return nil
	}
	return New(e.code, err.Error())
}
func (e customError) Set(msg string) Error {
	return New(e.code, msg)
}
func (e customError) Add(msg string) Error {
	return New(e.code, e.msg+" "+msg)
}
func New(code int32, msg string) Error {
	return &customError{code, msg}
}
