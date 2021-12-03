package errcode

import (
	"fmt"
	"net/http"
)

//编写常用的一些错误处理公共方法，标准化我们的错误输出

type Error struct {
	code    int      `json:"code,omitempty"`
	msg     string   `json:"msg,omitempty"`
	details []string `json:"details,omitempty"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{
		code:    code,
		msg:     msg,
		details: []string{},
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) MsgF(args []interface{}) string {
	return fmt.Sprintf(e.msg, args)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newErr := *e
	for _, d := range details {
		newErr.details = append(newErr.details, d)
	}
	return &newErr
}

// StatusCode 相对特殊的是 StatusCode 方法，
// 它主要用于针对一些特定错误码进行状态码的转换，
// 因为不同的内部错误码在 HTTP 状态码中都代表着不同的意义，
// 我们需要将其区分开来，便于客户端以及监控/报警等系统的识别和监听
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}

func SwitchErrorCode(err interface{}) *Error {
	switch err.(type) {
	case *Error:
		return err.(*Error)
	default:
		return ServerError
	}
}
