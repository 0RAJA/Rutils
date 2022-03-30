package errcode

import (
	"fmt"
	"net/http"
)

//编写常用的一些错误处理公共方法，标准化我们的错误输出

type Error struct {
	Code    int      `json:"code,omitempty"`
	Msg     string   `json:"msg,omitempty"`
	Details []string `json:"details,omitempty"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{
		Code:    code,
		Msg:     msg,
		Details: []string{},
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息: %s", e.SCode(), e.SMsg())
}

func (e *Error) SCode() int {
	return e.Code
}

func (e *Error) SMsg() string {
	return e.Msg
}

func (e *Error) MsgF(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args)
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

// StatusCode 相对特殊的是 StatusCode 方法，
// 它主要用于针对一些特定错误码进行状态码的转换，
// 因为不同的内部错误码在 HTTP 状态码中都代表着不同的意义，
// 我们需要将其区分开来，便于客户端以及监控/报警等系统的识别和监听
func (e *Error) StatusCode() int {
	switch e.SCode() {
	case Success.SCode():
		return http.StatusOK
	case ServerErr.SCode():
		return http.StatusInternalServerError
	case InvalidParamsErr.SCode():
		return http.StatusBadRequest
	case InsufficientPermissionsErr.SCode():
		fallthrough
	case UnauthorizedAuthNotExistErr.SCode():
		fallthrough
	case UnauthorizedTokenErr.SCode():
		fallthrough
	case UnauthorizedTokenGenerateErr.SCode():
		fallthrough
	case UnauthorizedNotLoginErr.SCode():
		fallthrough
	case UnauthorizedTokenTimeoutErr.SCode():
		return http.StatusUnauthorized
	case TimeOutErr.SCode():
		fallthrough
	case TooManyRequestsErr.SCode():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}

func SwitchErrorCode(err interface{}) *Error {
	switch err.(type) {
	case *Error:
		return err.(*Error)
	default:
		return ServerErr
	}
}
