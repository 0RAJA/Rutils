package errcode

//以 1 开头表示公共错误码
var (
	Success                      = NewError(0, "成功")
	ServerErr                    = NewError(1000, "服务内部错误")
	InvalidParamsErr             = NewError(1001, "入参错误")
	NotFoundErr                  = NewError(1002, "无结果")
	UnauthorizedAuthNotExistErr  = NewError(1003, "鉴权失败, 无法解析")
	UnauthorizedTokenErr         = NewError(1004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeoutErr  = NewError(1005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerateErr = NewError(1006, "鉴权失败，Token 生成失败")
	TooManyRequestsErr           = NewError(1007, "请求过多")
	TimeOutErr                   = NewError(1008, "请求超时")
	UnauthorizedNotLoginErr      = NewError(1009, "未登录")
	LoginErr                     = NewError(1010, "登录失败")
	InsufficientPermissionsErr   = NewError(1011, "鉴权失败,权限不足")
)
