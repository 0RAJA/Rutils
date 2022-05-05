package errcode

import "net/http"

//以 1 开头表示公共错误码
var (
	Success                     = NewError(0, "成功", http.StatusOK)
	ServerErr                   = NewError(1000, "服务内部错误", http.StatusInternalServerError)
	InvalidParamsErr            = NewError(1001, "入参错误", http.StatusBadRequest)
	NotFoundErr                 = NewError(1002, "无结果", http.StatusNotFound)
	UnauthorizedAuthNotExistErr = NewError(1003, "鉴权失败, 无法解析", http.StatusUnauthorized)
	UnauthorizedTokenErr        = NewError(1004, "鉴权失败，Token 错误", http.StatusUnauthorized)
	UnauthorizedTokenTimeoutErr = NewError(1005, "鉴权失败，Token 超时", http.StatusUnauthorized)
	TooManyRequestsErr          = NewError(1006, "请求过多", http.StatusTooManyRequests)
	TimeOutErr                  = NewError(1007, "请求超时", http.StatusRequestTimeout)
)
