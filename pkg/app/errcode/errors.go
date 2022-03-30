package errcode

//以 7,8 开头表示个人项目错误码

var (
	ErrCommentNotFind       = NewError(7000, "评论不存在")
	ErrPostNotFind          = NewError(7001, "不存在此文章")
	ErrPostNotEqual         = NewError(7002, "回复的评论不在回复的文章下")
	ErrUsernameNotFind      = NewError(7003, "用户不存在")
	ErrPasswordNotEqual     = NewError(7004, "密码错误")
	ErrPasswordEncodeFailed = NewError(7005, "加密密码失败")
	ErrPasswordRepeat       = NewError(7006, "密码与原密码相同")
	ErrDeleteManagerSelf    = NewError(7007, "不能删除自己")
	ErrStateRepeat          = NewError(7008, "状态设置重复")
	ErrDeletedState         = NewError(7009, "删除状态异常")
	ErrListPostInfosOptions = NewError(7010, "列出帖子信息选项异常")
	ErrAuthorizationFailed  = NewError(7011, "授权失败")
	ErrTokenNotExpired      = NewError(7012, "token未过期")
	ErrLoggedIn             = NewError(7013, "已经登录")
	ErrCommentLengthErr     = NewError(7014, "评论长度有误")
	ErrUsernameLengthErr    = NewError(7015, "用户名长度有误")
)

var (
	ExtErr              = NewError(8001, "file suffix is not supported")
	FileSizeErr         = NewError(8002, "exceeded maximum file limit")
	CreatePathErr       = NewError(8003, "failed to create save directory")
	CompetenceErr       = NewError(8004, "insufficient file permissions")
	RepeatedFileTypeErr = NewError(8005, "DuplicateFileType")
)
