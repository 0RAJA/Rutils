package errcode

//业务错误码

//login

var (
	CodeSignUPErr           = NewError(20000000, "注册失败")
	CodeUserExistErr        = NewError(20000001, "用户名已存在")
	CodeUserIsNotExist      = NewError(20000002, "用户不存在")
	CodeLoginErr            = NewError(20000003, "用户名或密码有误")
	CodeAuthIsNotExist      = NewError(20000004, "header中auth为空")
	CodeAuthFormatErr       = NewError(20000005, "header中auth格式有误")
	CodeRefreshTokenErr     = NewError(20000006, "refreshToken错误")
	CodeRefreshTokenFail    = NewError(20000007, "刷新Token失败")
	CodeMultiDeviceLoginErr = NewError(20000008, "多设备登陆")
	CodeResultNoROWSErr     = NewError(20000009, "查询不到任何数据")
	CodeNoCommunityID       = NewError(20000010, "查询不到community_id对应的信息")
	CodeAuthorIDNotFind     = NewError(20000011, "查询不到作者信息")
	CodePostTimeOut         = NewError(20000012, "帖子过期不可投票")
	CodeAlreadyVoted        = NewError(20000013, "帖子已经投票")
	CodeNoPostID            = NewError(20000014, "查询不到post_id对应的信息")
	CodeUploadFileErr       = NewError(20000015, "上传文件错误")
)
