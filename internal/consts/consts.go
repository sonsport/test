package consts

const (
	ProjectName          = "fuya_ark项目服务端"
	ProjectUsage         = "fuya_ark项目服务端"
	ProjectBrief         = "fuya_ark项目服务端"
	Version              = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName   = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey           = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	GTokenFrontendPrefix = "User:"              //gtoken登录 前台用户 前缀区分
	//for user
	CtxUserId        = "CtxUserId"
	CtxUserName      = "CtxUserName"
	CtxUserAvatar    = "CtxUserAvatar"
	CtxUserGender    = "CtxUserGender"
	CtxUserLiveState = "CtxUserLiveState"
	//for 登录相关
	TokenType          = "Bearer"
	CacheModeRedis     = 2
	BackendServerName  = "ark_server"
	FrontendMultiLogin = true
	GTokenExpireIn     = 10 * 24 * 60 * 60
	//统一管理错误提示
	CodeMissingParameterMsg = "请检查是否缺少参数"
	ErrLoginFailMsg         = "登录失败，账号或密码错误"
	ErrSecretAnswerMsg      = "密保问题不正确"
	ResourcePermissionFail  = "没有权限操作"
)
