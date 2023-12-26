package consts

const (
	ProjectName          = "fuya_ark_api"
	ProjectUsage         = "fuya_ark项目服务端"
	ProjectBrief         = "fuya_ark项目服务端"
	CaptchaDefaultName   = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey           = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	GTokenFrontendPrefix = "User:"              //gtoken登录 前台用户 前缀区分
	//for user
	CtxUserId              = "CtxUserId"
	CtxUserName            = "CtxUserName"
	CtxAppName             = "CtxAppName"             // app名称
	CtxAppChannel          = "CtxAppChannel"          // 渠道
	CtxClientLanguage      = "CtxClientLanguage"      // 用户语言
	CtxDeviceId            = "CtxDeviceId"            // 设备号
	CtxAppVersion          = "CtxAppVersion"          // app版本
	CtxClientMode          = "CtxClientMode"          // 手机型号
	CtxAppVersionCode      = "CtxAppVersionCode"      // app数字版本
	CtxClientSystemVersion = "CtxClientSystemVersion" // 手机系统版本
	CtxClientIP            = "CtxClientIP"            // 请求Ip
	CtxClientOs            = "CtxClientOs"            // 1为安卓 2为ios

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

const (
	TimeLayoutDate     = "2006-01-02"
	TimeLayoutTime     = "2006-01-02 15:04:05"
	TimeLayoutTimeHour = "2006-01-02 15"
	TimeLayoutMonth    = "200601"
	TimeLayoutDay      = "20060102"
	TimeSecondDate     = "20060102150405"
)
