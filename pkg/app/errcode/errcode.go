package errcode

import (
	"fmt"
)

//编写常用的一些错误处理公共方法，标准化我们的错误输出

type Err interface {
	error
	HCode() int
}

type Error struct {
	Code     int      `json:"code,omitempty"`
	Msg      string   `json:"msg,omitempty"`
	Details  []string `json:"details,omitempty"`
	httpCode int
}

var codes = map[int]string{}

func NewError(code int, msg string, httpCode int) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{
		Code:     code,
		Msg:      msg,
		Details:  []string{},
		httpCode: httpCode,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息: %s, 详细信息:%s", e.SCode(), e.SMsg(), e.SDetails())
}

func (e *Error) SCode() int {
	return e.Code
}

func (e *Error) SMsg() string {
	return e.Msg
}

func (e *Error) HCode() int {
	return e.httpCode
}

func (e *Error) SDetails() []string {
	return e.Details
}

func (e *Error) WithDetails(details ...string) *Error {
	newErr := *e
	for _, d := range details {
		newErr.Details = append(newErr.Details, d)
	}
	return &newErr
}
