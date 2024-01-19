package consts

const (
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	GTokenAdminPrefix        = "Admin:"             // Admin的gtoken前缀
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"

	GTokenUserPrefix   = "User:" // User的gtoken前缀
	CtxUserId          = "CtxUserId"
	CtxUserName        = "CtxUserName"
	CtxUserAvatar      = "CtxUserAvatar"
	CtxUserSign        = "CtxUserSign"
	CtxUserStatus      = "CtxUserStatus"
	CtxUserStatusBlock = 2

	// 错误提示
	ErrLoginFailMsg = "账号或密码错误！"
)
