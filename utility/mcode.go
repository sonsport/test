package utility

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

var (
	// Succeed 底层错误码 0 - 999
	Succeed                = gcode.New(0, "success", "")
	PermissionInsufficient = gcode.New(403, "Permission exception", "Permission exception")

	RequestRepeat = gcode.New(409, "request to repeat", "request to repeat")
	Internal      = gcode.New(500, "Internal Error", "An error occurred internally")

	// ParamError 业务层 1000 - 4999
	ParamError   = gcode.New(1000, "Parameter error", "Parameter error")
	ReqError     = gcode.New(1001, "Request error", "Request error")
	UserNotFound = gcode.New(1002, "User does not exist", "An error occurred internally")
	PwdError     = gcode.New(1003, "wrong password", "An error occurred internally")
	UNAUTHORIZED = gcode.New(1004, "", "") //未经授权的访问
	// UseBan 用户被封禁
	UseBan = gcode.New(1005, "UseBan", "UseBan")
	// Device 设备被封禁
	Device                = gcode.New(1006, "DeviceBan", "DeviceBan")
	NotAdminErr           = gcode.New(1007, "Not admin", "current user is not admin")
	ReSetBLErr            = gcode.New(1008, "Reset black list", "current user is in black list")
	NotInBlErr            = gcode.New(1009, "Not in black list", "current user is not in black list")
	OssTokenErr           = gcode.New(1012, "Get upload fail", "Get upload fail")
	NoAccessJoinRoom      = gcode.New(1014, "No access join live room", "No access join live room")
	CreateCodapayOrderErr = gcode.New(1015, "Create order fail", "Create order fail")
	CreateXenditOrderErr  = gcode.New(1016, "Create order fail", "Create order fail")
	NotSupportPayment     = gcode.New(1017, "Not Support Payment", "Not Support Payment")
	CreateOrderErr        = gcode.New(1018, "Create order fail", "Create order fail")
	InvalidNotification   = gcode.New(1019, "Invalid notification", "Invalid notification")

	AnchorApplyRepeat     = gcode.New(1020, "您已经是主播了，无需重复申请", "您已经是主播了，无需重复申请")
	AnchorNotExist        = gcode.New(1021, "您要查询的主播不存在", "您要查询的主播不存在")
	GuildNotExist         = gcode.New(1022, "您搜索的公会不存在", "您搜索的公会不存在")
	GuildCreateNotAllowed = gcode.New(1023, "当前账号已经创建过公会了，请重新注册账号创建公会", "当前账号已经创建过公会了，请重新注册账号创建公会")
	GuildNameAlready      = gcode.New(1024, "公会名称已经存在了，换一个名称试试吧", "公会名称已经存在了，换一个名称试试吧")
	CreateRoomErr         = gcode.New(1025, "Create room fail", "Create room fail")

	InsufficientBalance  = gcode.New(1026, "Insufficient balance", "Insufficient balance")
	UnusualGiftRecipient = gcode.New(1027, "unusual gift recipient", "unusual gift recipient")
	NotFound             = gcode.New(1028, "Not Found", "Resource does not exist")

	BeginLiveFailed       = gcode.New(1030, "Live room begin failed", "Live room begin failed")
	RoomAlreadyClosed     = gcode.New(1031, "Live room already closed", "Live room already closed")
	NoAuthorizedToOperate = gcode.New(1032, "No authorized to operate", "No authorized to operate")

	// NoAuthorizedJoinRoom 踢出无法进入直播间
	NoAuthorizedJoinRoom = gcode.New(1033, "No authorized to join room", "No authorized to join room")
	// ProductNotFound 商品不存在或已下架
	ProductNotFound = gcode.New(1034, "The product does not exist or has been removed", "The product does not exist or has been removed")
	// UserDressUpExpired 商品过期了
	UserDressUpExpired = gcode.New(1035, "The product expired", "The product expired")

	// ErrAuthSignCheckError 游戏错误码
	ErrAuthSignCheckError  = gcode.New(1040, "auth签名校验失败", "auth签名校验失败")
	ErrAuthBusiIDInfoError = gcode.New(1042, "auth busi_id info 校验失败", "auth busi_id info 校验失败")
	ErrCreateOrderError    = gcode.New(1043, "订单创建失败", "订单创建失败")
	//NotAcceptBattle 对方不接受pk
	NotAcceptBattle  = gcode.New(1044, "The other party does not accept pk", "The other party does not accept pk")
	RoomAdminNoQuota = gcode.New(1045, "Room manager slots are exhausted", "房间管理员名额用完")
	InviteErr        = gcode.New(1046, "Invite Fail", "Invite Fail")
	NotLiveErr       = gcode.New(1047, "Your broadcast is disabled, pls try again later", "Your broadcast is disabled, pls try again later")
	// CouponsAreInvalid 优惠券过期
	CouponsAreInvalid = gcode.New(1048, "Coupons are invalid", "Coupons are invalid")
	//LikeActivityIsNotSupported 点赞不支持
	LikeActivityIsNotSupported = gcode.New(1049, "Like Activity Is Not Supported", "Like Activity Is Not Supported")
	IsNotFirst                 = gcode.New(1050, "IsNotFirst", "IsNotFirst")
	// CreateDiamondPoolErr 生成钻石抽奖奖池失败
	CreateDiamondPoolErr = gcode.New(1051, "Create diamond pool err", "Create diamond pool err")
	// CreateDiamondDrawCountErr 钻石抽奖次数错误
	CreateDiamondDrawCountErr = gcode.New(1052, "Create diamond draw count err", "Create diamond draw count err")
	// FrequentOperation 操作频繁
	FrequentOperation = gcode.New(1053, "Frequent operation", "Frequent operation")
	//LackOfProps 道具不足
	LackOfProps = gcode.New(1054, "Lack Of Props", "Lack Of Props")
	//RedEnvelopeExpired 红包已经过期
	RedEnvelopeExpired = gcode.New(1055, "RedEnvelopeExpired", "")
	//RedEnvelopeAlreadyGrabbed 红包已经抢完
	RedEnvelopeAlreadyGrabbed = gcode.New(1056, "RedEnvelopeAlreadyGrabbed", "")
	//RedEnvelopeBillingFailed 未存在结算红包数据
	RedEnvelopeBillingFailed = gcode.New(1057, "RedEnvelopeBillingFailed", "")
	//RoomRedEnvelopesExist 房间红包存在
	RoomRedEnvelopesExist = gcode.New(1058, "RoomRedEnvelopesExist", "")
	//LackOfGiftProps 道具不足
	LackOfGiftProps = gcode.New(1059, "LackOfGiftProps", "")
	//PhoneNotMatch 手机号和预留的手机不一致
	PhoneNotMatch = gcode.New(1060, "PhoneNotMatch", "")
	//SmsTimeTooShort 验证码发送时间间隔太短
	SmsTimeTooShort = gcode.New(1061, "SmsTimeTooShort", "")
	//SmsCodeMismatch 验证码不匹配
	SmsCodeMismatch = gcode.New(1062, "SmsCodeMismatch", "")
	//BalanceIsLock 余额锁定
	BalanceIsLock = gcode.New(1063, "BalanceIsLock", "")
	//AgentRoleError 代理商身份异常
	AgentRoleError = gcode.New(1064, "AgentRoleError", "")
	//CaptchaAbnormal 验证码异常
	CaptchaAbnormal = gcode.New(1065, "验证码异常", "")
	//AccountAlreadyExists 账号已经存在
	AccountAlreadyExists = gcode.New(1067, "账号已经存在", "")
	//ApplyAlreadyExists 申请已经存在
	ApplyAlreadyExists = gcode.New(1068, "ApplyAlreadyExists", "")
	//RoomIsFullErr 房间已经满了
	RoomIsFullErr = gcode.New(1069, "RoomIsFullErr", "RoomIsFullErr")
	RoomIsNotJoin = gcode.New(1070, "RoomIsNotJoin", "")

	//AnchorNotThisGuild 主播不是该公会旗下
	AnchorNotThisGuild = gcode.New(1101, "Agency ini tidak ada host", "Agency ini tidak ada host")
	AnchorNumNotEnough = gcode.New(1102, " jumlah host tidak cukup", " jumlah host tidak cukup")
	GuildAlreadyApply  = gcode.New(1103, "agency sudah daftar", "agency sudah daftar")
	//InsufficientPermissionInvite 邀新权限不足
	InsufficientPermissionInvite = gcode.New(1104, "邀新权限不足", "邀新权限不足")
	//AccountAssociationPermissionInvite 账号关联停止邀新
	AccountAssociationPermissionInvite = gcode.New(1105, "账号关联停止邀新", "账号关联停止邀新")
	// UserIsNotNewUser 不是新用户
	UserIsNotNewUser = gcode.New(1201, "UserIsNotNewUser", "")
	//FeedIsCheck 动态审核
	FeedIsCheck = gcode.New(1200, "审核模式动态发布提示", "审核模式动态发布提示")

	CreateTopicSuccess            = gcode.New(1300, "Create Topic Success", "Create Topic Success")               //话题创建成功 等待审核
	TopicHasCreated               = gcode.New(1301, "The topic has been created.", "The topic has been created.") //话题已经创建，不能重复创建
	BattleNotReuse                = gcode.New(1302, "不能重复使用", "")                                                 //不能重复使用
	HostRejectCode                = gcode.New(1303, "", "")                                                       //主播已拒绝连麦申请
	InsufficientKickOutPermission = gcode.New(1304, "", "")                                                       //踢出权限不足

	// SqlResultFail 数据库错误码 5000 - 10000
	SqlResultFail   = gcode.New(5000, "Internal Error", "Internal Error")
	RedisResultFail = gcode.New(5001, "Internal Error", "Internal Error")
)
