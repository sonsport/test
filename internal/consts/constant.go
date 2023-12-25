package consts

// ===================支付/花费相关START===================
// 支付渠道状态 定义常量
const (
	// PaymentStateOff 上下架状态
	PaymentStateOff = 0 //下架
	// PaymentStateOn 上下架状态
	PaymentStateOn = 1 //上架
)

const (
	LocalHost = "127.0.0.1"
)

const (
	CurrencyCodeDefault = 0   //默认
	CurrencyCodeID      = 360 //印尼
	CurrencyCodeMY      = 458 //马来西亚
	SystemOsDefault     = 0   //0默认全部系统  1安卓 2ios
	SystemOsAndroid     = 1   //0默认全部系统  1安卓 2ios
	SystemOsIos         = 2   //0默认全部系统  1安卓 2ios
	SystemOsServer      = 3   //0默认全部系统  1安卓 2ios 3运营号
)

// CommodityStateOn 商品档位状态
const (
	CommodityStateOn       = 1 //上架
	DiscountTypeNone       = 0 //无折扣折扣
	DiscountTypeFirst      = 1 //首冲折扣
	OneDiscountTypeAgent   = 2 //一级代理充值
	TwoDiscountTypeAgent   = 3 //二级代理充值
	ThreeDiscountTypeAgent = 4 //三级代理充值
	FourDiscountTypeAgent  = 5 //四级代理充值

)

const (
	PayTypeOffLine  = 0 //线下支付
	PayTypeGoogle   = 1 //Google支付
	PayTypeApple    = 2 //Apple支付
	PayTypeXendit   = 3 //xenditpay支付
	PayTypeCodapay  = 4 //codapay支付
	PayTypePayerMax = 5 //PayerMax支付
	PayTypeHaiPay   = 6 //HaiPay支付
)

// 订单状态
const (
	OrderStatePending = 0 //待支付
	OrderStatePaid    = 1 //已支付
	OrderStateRefund  = 2 //退款
	OrderStateTimeOut = 3 //超时

)

// 支付状态/ 交易支付状态公用 0: 未知 1: 支付成功 2:支付中  3: 支付失败
const (
	PayStatusUnknown  = int8(0)
	PayStatusSuccess  = int8(1)
	PayStatusWaiting  = int8(2)
	PayStatusFailed   = int8(3)
	PayStatusRefunded = int8(4) // 退款
)

// 订单通知处理状态
const (
	PayNoticePending = int8(0) //待处理
	PayNoticeSuccess = int8(1) //1成功  2失败
	PayNoticeFail    = int8(2) //2失败
	PayNoticeIgnore  = int8(3) //3重复通知等需要忽略掉
)

var DepletionTypes = []int{
	DepletionTypeRecharge, DepletionTypeSendGift, DepletionTypeSendLuckyGift, DepletionTypeLuckyGiftWinning,
	DepletionTypGameCharge, DepletionTypGameRecharge, DepletionShopCharge, InviteRewards, DepletionTypSettlementRecharge, RechargeRewards, DepletionTypeSystemCharge,
	DepletionTypeNewUser, DepletionTypeHotAward, DepletionTypeCoupon, DepletionTypeSystemRecharge, TurntableDraw, ComplexPrizeDraw, GameCommissionIncome,
	RedEnvelopeCharge, RedEnvelopeReCharge, RoomUnlockCharge, WishGiftRecharge, RoomUnlockRecharge, ActivityRecharge, NobleCharge, RaffleCharge, AgentRecharge,
	QuestRecharge, DepletionNew3DaysReward}

// 交易类型
const (
	DepletionTypeRecharge          = 1  //支付充值
	DepletionTypeSendGift          = 2  //送礼物 普通
	DepletionTypeSendLuckyGift     = 3  //送礼物 幸运
	DepletionTypeLuckyGiftWinning  = 4  //幸运礼物中奖加钻
	DepletionTypGameCharge         = 5  //游戏相关-游戏减钻
	DepletionTypGameRecharge       = 6  //游戏相关-游戏加钻
	DepletionShopCharge            = 7  //商城购买装扮
	InviteRewards                  = 8  //邀新奖励
	DepletionTypSettlementRecharge = 9  //结算兑换钻石
	RechargeRewards                = 10 //充值奖励
	DepletionTypeNewUser           = 11 //新用户奖励
	DepletionTypeHotAward          = 12 //热门主播每分钟开播奖励
	DepletionTypeCoupon            = 13 //钻石优惠券奖励
	DepletionTypeSystemRecharge    = 14 //后台充值
	TurntableDraw                  = 15 //转盘抽奖中奖送钻
	ComplexPrizeDraw               = 16 //活动抽奖中奖送钻
	GameCommissionIncome           = 17 //游戏抽成提现
	RedEnvelopeCharge              = 18 //红包扣除
	RedEnvelopeReCharge            = 19 //红包加钻
	RoomUnlockCharge               = 20 //房间解锁费用
	WishGiftRecharge               = 21 //心愿礼物奖励
	RoomUnlockRecharge             = 22 //房间解锁收益
	ActivityRecharge               = 23 //活动赠送
	AgentRecharge                  = 24 //代理充值
	NobleCharge                    = 26 //开通贵族
	RaffleCharge                   = 27 //钻石抽奖
	DepletionTypeSystemCharge      = 30 //后台扣除
	QuestRecharge                  = 31 //任务赠送
	DepletionNew3DaysReward        = 32 //新主播3周内3天开播达标6小时奖励3000钻石
)

// 收、支类型
const (
	BalanceTypeExpend = 1 //支出
	BalanceTypeIncome = 2 //收入
)

//===================支付、花费相关END===================

type LoginType int64

// OauthToken loginType
const (
	LoginAccount  LoginType = 0 //账号登陆
	LoginGoogle   LoginType = 1 //谷歌登陆
	LoginFacebook LoginType = 2 //fb登陆
	LoginApple    LoginType = 3 //apple登陆

)

// IncomeType 收益类型
type IncomeType int64

// IncomeRecord DepletionType
const (
	NUll       IncomeType = -1 //空
	Gift       IncomeType = 2  //礼物 普通
	LuckyGift  IncomeType = 3  //礼物 幸运
	Settlement IncomeType = 4  //系统结算
	Recharge   IncomeType = 5  //后台充金币
	Charge     IncomeType = 6  //后台扣金币
	Quest      IncomeType = 7  //任务充金币
)

// USA 地区编码常量
const USA = 840
const USACn = "美国"
const IndonesiaCn = "印度尼西亚"
const MalaysiaCn = "马来西亚"
const DefaultAvatar = "/common/fuya_def_head.png"

const (
	UserNormal   = 0 //普通用户
	UserAnchor   = 1 //认证主播
	UserMaster   = 2 //认证公会长
	UserOfficial = 3 //官方号 //官方号发的动态、话题都属于官方
	UserServer   = 4 //运营号
	UserRobot    = 6 //机器人
	UserManager  = 8 //系统管理员

)

// 礼物相关
const (
	GiftTypeNormal   = 0 //普通礼物
	GiftTypeLucky    = 1 //幸运礼物
	GiftTypeActivity = 2 //活动礼物
	ExclusiveGift    = 3 //专属礼物
)

// 幸运礼物奖池类型相关
const (
	LuckyGiftTypePool90      = 0 //90%返奖率奖池
	LuckyGiftTypePool95      = 1 //95%返奖率奖池
	LuckyGiftTypePool95_5000 = 2 //95%返奖率奖池 有5000倍
)

// 关注相关
const (
	FollowNone = 0 //无关注
	Follower   = 1 //已关注
	FollowEach = 2 //相互关注
)

// 房间开播信息
const (
	RoomStatusAll         = -1 //所有状态
	RoomStatusPre         = 0  //预开播
	RoomStatusLiving      = 1  //开播中
	RoomStatusEnd         = 2  //直播正常结束
	RoomStatusTickOut     = 3  //由巡管或运营管理员将直播间关闭
	RoomStatusPollOut     = 4  //由服务器轮询出异常退出房间
	RoomStatusPreOut      = 5  //预开播时检测到有未关播的房间，将其退出房间
	RoomStatusOtherOut    = 6  //其他异常退出房间
	RoomStatusPreEnd      = 7  //正常关闭预开播
	RoomStatusLeave       = 8  //暂时离开
	UserExitRoomByManager = 9  //管理后台退出

)

// 私密直播间状态 0 正常直播间，1 小黑屋，2 钻石房间 3 密码房间
const (
	PrivateStatusNormal     = 0 //正常直播间
	PrivateStatusSmallBlack = 1 //小黑屋
	PrivateStatusDiamond    = 2 //钻石房间
	PrivateStatusPassword   = 3 //密码房间
)

// appConfig
const (
	AppConfigDefaultSystem  = 0 //默认所有系统
	AppConfigDefaultChannel = 0 //默认渠道
	AppConfigDefaultVersion = 0 //默认版本
)

const (
	UserExitRoomNil        = 0 //当前用户没有退出房间
	UserExitRoomBySelf     = 1 //正常退出房间
	UserExitRoomByRoomAdm  = 2 //房间管理员踢出房间
	UserExitRoomByAuthor   = 3 //主播踢出房间
	UserExitRoomBySuperAdm = 4 //超级管理员踢出房间
	UserExitRoomByOther    = 5 //异常原因退出房间
)

// 房间pk枚举
const (
	//RoomLiveBattleTypeRandom 房间pk 随机
	RoomLiveBattleTypeRandom = iota
	//RoomLiveBattleTypeSpecify 房间pk 指定
	RoomLiveBattleTypeSpecify
	//RoomLiveBattleTypeAgentMatch 房间pk 工会比赛
	RoomLiveBattleTypeAgentMatch
	//RoomLiveBattleTypeMatch pk天梯赛
	RoomLiveBattleTypeMatch
)

// 房间PK时的状态  0 匹配中 1 pk中 2 邀请人pk取消 3 对方拒绝 4 惩罚中 5 pk正常结束 6 pk异常结束
const (
	PkStatusMatch      = iota //匹配中
	PkStatusConnect           //pk中
	PkStatusCancel            //邀请人pk取消
	PkStatusReject            //对方拒绝
	PkStatusPunishment        //惩罚中
	PkStatusEnd               //pk正常结束
	PkStatusError             //pk异常结束
)

// 直播间PK时的状态  0 未pk 1 随机PK匹配中 2 邀请PK中 3 正在PK
const (
	LiveBattleStatusNull = iota //未pk
	LivePkStatusRandom          //随机PK匹配中
	LivePkStatusInvite          //邀请PK中
	LivePkStatusIng             //正在PK
)

const (
	BattleTypeRand       = 0 //随机pk
	BattleTypeAppoint    = 1 //指定pk
	BattleTypeBackground = 2 //后台拉起
	BattleTypeMatch      = 3 //pk天梯赛

)

// 用户表相关
const (
	UserLiveStateNoLive = 0 //未开播
	UserLiveStateLiving = 1 //已开播
)

//  黑名单相关

const (
	BlackEndTimeForever = 0 //永久拉黑
)

// 列表类型
const (
	HotList    = 0
	NewList    = 1
	FollowList = 2
)

const (
	GuildStatusOk    = 1 // 生效
	GuildStatusBreak = 2 // 解散

	GuildApplyStatusWait   = 0 //等待审核
	GuildApplyStatusOk     = 1 //审核成功
	GuildApplyStatusReject = 2 //审核被拒

	MyGuildStatusNone   = 0 // 无工会
	MyGuildStatusWait   = 1 // 待审核
	MyGuildStatusMaster = 2 // 审核通过，工会长
	MyGuildStatusAnchor = 3 // 公会主播

	GuildAnchorRoleMaster = 1 //公会长
	GuildAnchorRoleMember = 2 //成员

	GuildAnchorStatusOk     = 1 //成员正常
	GuildAnchorStatusFailed = 2 //成员失效

	// AnchorStar0 主播星级
	AnchorStar0 = 0
	AnchorStar1 = 1
	AnchorStar2 = 2
	AnchorStar3 = 3
	AnchorStar4 = 4
	AnchorStar5 = 5

	//主播星级对应排序分数

	AnchorStarScore0 = 0
	AnchorStarScore1 = 100000
	AnchorStarScore2 = 200000
	AnchorStarScore3 = 300000
	AnchorStarScore4 = 400000
	AnchorStarScore5 = 500000

	//主播对应标签

	AnchorLabelTypeNew    = 0 //最新主播
	AnchorLabelTypeTop    = 1 //热一主播
	AnchorLabelTypePretty = 2 //颜值主播
	AnchorLabelTypeDance  = 3 //热舞主播
	AnchorLabelTypeCharm  = 4 //魅力主播
	AnchorLabelTypeKitty  = 5 //可爱主播
)

const (
	MomentTypeLike          = 0 //点赞
	MomentTypeReportComment = 1 //回复动态
	MomentTypeReportFeed    = 2 //回复评论
)

// 房间管理员相关
const (
	RoomAdminInvalid = 0 //无效
	RoomAdminValid   = 1 //有效

	RoomAdminMaxCount = 5 //房间管理员最多
)

// 房间操作
const (
	RoomShutUpMin      = 1 //直播间禁言N分钟
	RoomShutUpForever  = 2 //直播间永久禁言
	RoomKickOutHour    = 3 //直播间踢出N小时
	RoomKickOutForever = 4 //直播间永久踢出
)

type TimeLevel int

const (
	TimeLevelHours TimeLevel = 1 //统计时间维度小时
	TimeLevelDay   TimeLevel = 2 //统计时间维度日
)

const OneDaySeconds = 86400                 //一天的秒
const FiveMinSeconds = 60 * 5               //五分钟的秒
const HalfHoursSeconds = 60 * 30            //半小时的秒
const OneHoursSeconds = 60 * 60             //一小时的秒
const TwoHoursSeconds = OneHoursSeconds * 2 //两小时的秒

const (
	BoardUserByRoom = 31 //直播间-用户贡献榜
)

// StatisticsGameKeys 统计游戏
var StatisticsGameKeys = []string{"all", "10001", "10002", "10005", "10006"}

const (
	RoomScoreTypeByInit        = 0 //初始化
	RoomScoreTypeByGift        = 1 //收到的礼物
	RoomScoreTypeByLive        = 2 //开播时长
	RoomScoreTypeByViewer      = 3 //进来的观众
	RoomScoreTypeByHotOne      = 4 //热一礼物
	RoomScoreTypeByHotRotation = 5 //热门轮播榜单
	RoomScoreTypeByTemporary   = 6 //临时时间到期

)
const (
	GuildTotalDataTypeLastWeek = 1
	GuildTotalDataTypeThisWeek = 0
)

const (
	UserBoardLastWeek = 1
	UserBoardThisWeek = 0
)

// 客户端语言
const (
	LanguageID = "id" //印尼语
	LanguageEN = "en" //英语
	LanguageMS = "ms" //马来语
)

// 直播统计相关
const (
	MinCharmEffectLiveDay = 7200 //针对全平台，性感主播有效直播天 最低有效直播累计要满3600秒
	MinEffectLiveDay      = 3600 //针对全平台，有效直播天 最低有效直播累计要满3600秒
	MinEffectLiveSeconds  = 1800 //针对全平台，一个直播最低要播满1800秒才会统计进有效直播秒数
)

const (
	ShareByRecharge  = 0 //邀请充值奖励
	ShareByWatchLive = 1 //邀请观看直播奖励

	InvalidRewardForWatchLive = 2 // 邀请看播奖励无效
)

const (
	ShareNone        = 0 //初始化
	ShareWaiting     = 1 //待加币
	ShareLiveSuccess = 2 //开播30分钟奖励成功发放
)

const (
	TimeDayHour  = 86400           // 24 * 60* 60 一天
	TimeDayWeek  = TimeDayHour * 7 // 24 * 60* 60*7 一天
	TimeOneMonth = 2592000         // 30 * 24 * 60 * 60 一个月
)

const (
	CommodityGiftTypeShop   = 1 //赠送类 装扮
	CommodityGiftTypeExp    = 2 //赠送类 经验
	CommodityAgentPropsGift = 3 //代理赠送类型 礼物
	CommodityAgentPropsVip  = 4 //代理赠送类型 vip
	CommodityAgentPropsShop = 5 //代理赠送类型 装扮
)

const (
	TurntableTypeShop    = 1                       //赠送类 1 装扮
	TurntableTypeDiamond = 2                       //赠送类 2 钻石
	TurntableTypeCoupons = 3                       //赠送类 3 优惠券
	TurntableTypeGift    = 5                       //赠送类 5 礼物
	DiamondPicUrl3       = "diamonds/diamond3.png" // 钻石图片路由
	DiamondPicUrl100     = "diamonds/diamond100.png"
	DiamondPicUrl20000   = "diamonds/diamond20000.png"
	CouponsPicUrl        = "common/game/1675909435751.png" // 钻石图片路由
)

const (
	// PropsLogFromTypeGift      = 1  // 1  赠送扣除
	// PropsLogFromTypeGive      = 2  // 2  赠送领取
	// PropsLogFromTypeActive    = 4  // 4  活跃领取
	// PropsLogFromTypeLike      = 5  // 5  点赞动态
	// PropsLogFromTypeComment   = 6  // 6  评论动态
	// PropsLogFromTypeFocus     = 7  // 7  关注
	// PropsLogFromTypeRank      = 8  // 8  排行榜
	// PropsLogFromTypeLook1     = 9  // 9  看播1分钟
	// PropsLogFromTypeLook5     = 10 // 10 看播5分钟
	// PropsLogFromTypeLook30    = 11 // 11 看播30分钟
	// PropsLogFromTypeGiftProps = 12 // 12 赠送道具
	// PropsLogFromTypeSendGift  = 14 // 14 送礼物
	// PropsLogFromTypeBuyDress  = 15 // 15 购买装扮

	PropsLogFromTypeUse      = 3  // 3  使用扣除
	PropsLogFromTypeRecharge = 13 // 13 充值
)

// 任务类型
const (
	QuestTypeNesting         = iota // 1  套娃任务
	QuestTypeSignIn                 // 1  签到
	QuestTypeFeed                   // 2  发布动态
	QuestTypeTopUp                  // 3  累积充值钻石
	QuestTypeDynamic                // 4  每日活跃值
	QuestTypeWatchLive              // 5  观看直播
	QuestTypeGift                   // 6  送礼钻石
	QuestTypeFocus                  // 7  关注
	QuestTypeShare                  // 8  分享
	QuestTypeBeFocused              // 9  被关注
	QuestTypeBeShared               // 10 被分享
	QuestTypeIncome                 // 11 金币收入
	QuestTypeGiftUser               // 12 直播送礼人数
	QuestTypeDayWaterDiamond        // 13 主播日流水钻石
	QuestTypeOneLiveTime            // 14 单场开播时间
	QuestTypeOneBeFocused           // 15 单场被关注
	QuestTypeWeekDynamic            // 17 本周活跃

)

// 活动类型
const (
	ActivityTypeSendGiftMsg = 1 // 赠送活动礼物消息通知
	ActivityTypeCloseLive   = 2 //关播发送消息
)

// 封禁相关
const (
	BanSendPrivateMessage = 1 //禁止发私信
	BanSendRoomMessage    = 2 //禁发直播间消息
	BanLive               = 3 //禁开播
	BanLogin              = 4 //禁登录
	BanDevice             = 5 //禁设备
)

const (
	GameIsNotHow    = 0 // 游戏不是热门
	GameIsHot       = 1 // 热门游戏
	GameIsNotMini   = 0 // 全屏游戏
	GameIsMini      = 1 // 半屏游戏
	GameTypeJoyUs   = 0 // 自研游戏类型
	GameTypeJoyPlay = 1 // JoyPlay游戏
	GameTypeCocos   = 2 // Cocos游戏
	GameAddCoin     = 1 // 游戏加币
	GameSubCoin     = 2 // 游戏减少币
)

const (
	RechargePointsType   = 1 // 1充值
	CumulativePointsType = 2 // 2累计积分
	HourFirstPointsType  = 4 // 4小时第一加成
	RankPointsType       = 5 // 5排名加成
)

const (
	RoomAnonymousAvatar = "https://static.himilive.vip/common/banner/sun_glasses.png" //匿名头像
	MysteryManAvatar    = "https://static.himilive.vip/common/xunzhang/shenmiren.png" //神秘人头像
)

// 内容安全相关
const (
	SafeInfoPortrait = 1 //头像
	SafeInfoNickName = 2 //昵称
	SafeInfoIntro    = 3 //签名
)

const (
	WaitSafeInfo   = 0 //待审核
	AllowSafeInfo  = 1 //通过
	RefuseSafeInfo = 2 //不通过
)

// 日活周活月活
const (
	ActiveDay   = 1
	ActiveWeek  = 2
	ActiveMonth = 3
)

// 警告相关
const (
	Warn       = 1 //单纯警告，无惩罚
	WarnAndBan = 2 //警告且封禁
	CloseRoom  = 3 //关闭直播间
)

const (
	ConditionsTypeUserRankWeek   = iota + 1 //用户周榜
	ConditionsTypeUserRankMoon              //用户月榜
	ConditionsTypeAnchorRankWeek            //主播周榜
	ConditionsTypeAnchorRankMoon            //主播月榜
	ConditionsType7DayLogin                 //潜力王 7天登录
	ConditionsTypeSettlement                //勤劳之星 上周结算
	ConditionsTypeRecharge                  //钻石之星 充值
)

const ( //特权类型
	PrivilegeTypeMount         = iota + 1 //进场坐骑
	PrivilegeTypeHeadFrame                //头像框
	PrivilegeTypeBadge                    //徽章
	PrivilegeTypeBubbleFrame              //气泡框
	PrivilegeTypeGoldenLight              //金光
	PrivilegeTypeExpBonus                 //经验值加成
	PrivilegeTypeListHide                 //榜单隐身
	PrivilegeTypeRoomHide                 //进房隐身
	PrivilegeTypeUserHide                 //神秘人
	PrivilegeTypeDotKick                  //防踢防禁言
	PrivilegeTypeExclusiveGift            //专属礼物
	PrivilegeTypeDiamondBonus             //充值钻石加成
)

const UserIdCtxKey = "UserId"

const (
	ApplyPopupReal = 1 // Apple评分真弹窗
	ApplyPopupFake = 2 // Apple评分假弹窗

	PopupActionBlockTheReport   = 1 // 拉黑举报
	PopupActionLiveStreamingLag = 2 // 直播卡顿
	PopupActionSwitchFrequently = 3 // 频繁切换直播间
	PopupActionSendGift         = 4 // 送礼
	PopupActionCloseLive        = 5 // 主播关播

	HourSeconds = 60 * 60
	DaySeconds  = 24 * 60 * 60
)

const ( //修改相关
	UpdateStar   = 1 //调整星级
	UpdateGuild  = 2 //换公会
	UpdateLabel  = 3 //换标签
	UpdateRemark = 4 //换备注印象
)

const ( //通知类型 notice_type
	SystemResetStar = 1 //系统自动调整星级
	ManResetStar    = 2 //人工调整星级
)

const ( //通知消息动作类型 notice_type
	SystemStar_Keep = 0 //系统星级保级
	SystemStar_Up   = 1 //系统星级升级
	SystemStar_Down = 2 //系统星级降级
	ManStar_Up      = 1 //人工星级降级
	ManStar_Down    = 2 //人工星级降级
)

const ( //运营号通知状态
	ServerNoticeState_Unhandle = 0 //未处理
)

// 运营号通知相关
const (
	NoticeTypeSystem         = 1  //系统星级调整
	NoticeTypePerson         = 2  //人工星级调整
	NoticeTypeLabel          = 3  //人工标签调整
	NoticeTypeBan            = 4  //人工封禁
	NoticeTypeWarn           = 5  //人工警告
	NoticeTypeLiveAllow      = 6  //主播试播通过
	NoticeTypeSystemUnSettle = 7  //系统通知周结算
	NoticeTypeUnLive         = 8  //系统通知未开播
	NoticeTypeShare          = 9  //人工通知邀新
	NoticeTypeSystemShare    = 10 //系统通知邀新

	MessageTypeSysMaintain       = 0 //系统星级保持
	MessageTypeSysRise           = 1 //系统星级升级
	MessageTypeSysDecline        = 2 //系统星级降级
	MessageTypeHumRise           = 1 //系统星级升级
	MessageTypeHumDecline        = 2 //系统星级降级
	MessageTypeLabelUpdate       = 1 //调整标签
	MessageTypeBanPrivateMessage = 1 //人工封禁发私信
	MessageTypeBanRoomChat       = 2 //人工封禁发公聊
	MessageTypeBanLive           = 3 //人工禁开播
	MessageTypeBanLogin          = 4 //人工禁登录
	MessageTypeBanDeVice         = 5 //人工禁设备
	MessageTypeWarn1             = 1 //开车中直播
	MessageTypeWarn2             = 2 //未成年直播
	MessageTypeWarn3             = 3 //同时播2个app
	MessageTypeWarn4             = 4 //不是同一个主播
	MessageTypeWarn5             = 5 //睡播
	MessageTypeWarn6             = 6 //黑播
	MessageTypeWarn7             = 7 //吵嚷/声音不清
	MessageTypeWarn8             = 8 //空播太久
	MessageTypeLiveAllow         = 1 //试播通过
	MessageTypeSystemUnSettle    = 1 //系统通知主播已开播未结算
	MessageTypeSystemUnBank      = 2 //系统通知可结算未填写银行卡信息
	MessageTypeSystemUnLive3     = 1 //系统通知3天未开播主播
	MessageTypeSystemUnLive7     = 7 //系统通知7天未开播主播
	MessageTypePersonShareFail   = 1 //人工通知邀新失败
	MessageTypeSystemShareMax    = 1 //系统通知邀新达到上限

	NoticeUnRead = 0 //未读
	NoticeRead   = 1 //已知晓
	NoticeRemark = 2 //已处理
)

const ( //话题状态
	TopicState_Pending = 0 //未审核
	TopicState_Pass    = 1 //审核通过
	TopicState_NoPass  = 2 //不通过
	TopicState_Deleted = 3 //已删除
	TopicType_Ordinary = 0 //0普通话题 1官方 2审核模式话题
	TopicType_Official = 1 //官方
	TopicType_Review   = 2 //审核模式话题
)

const ( //话题分类
	TopicClassify_My     = 0 //我的话题
	TopicClassify_Hot    = 1 //热门话题
	TopicClassify_Follow = 2 //关注话题
)

const ( //动态分类
	FeedOfficial = 0 //官方动态
	FeedPlaza    = 1 //广场动态
	FeedFriends  = 2 //好友动态
	FeedTopic    = 3 //话题动态
)

const ( //奖励相关
	RewardForNewAnchorLive = 1 //3周内新主播3天有效直播，时长大于6小时
)

const ( //新官方消息类型
	TitleTypeShare  = 1 //邀请好友 Undang teman
	TitleTypeSystem = 2 //系统消息 Info sistem
	TitleTypeSettle = 3 //结算消息 Info Penyelesaian
	TitleTypeActive = 4 //活动信息 Info Event
	TitleTypeBattle = 5 //天梯赛信息 PK Legends

	SubTitleTypeDiamondReward = 1 //钻石奖励 Hadiah diamond
	SubTitleTypeDiamondIncome = 2 //钻石收益 Bonus diamond

	LastTitleTypeInviter = 1 //邀请人 Pengundang
	LastTitleTypeInvited = 2 //被邀请人 Teman diundang

)

var GameWarnToUser = []string{
	"fanwd@fuya.live", "chenyan@fuya.live", "zhupeisheng@fuya.live", "caizhaoliang@fuya.live",
	"hufeijie@fuya.live", "mofan@fuya.live", "zenghuiyang@fuya.live", "liuhang@fuya.live",
}

var SettleToUser = []string{
	"fanwd@fuya.live", "chenyan@fuya.live", "caizhaoliang@fuya.live", "hufeijie@fuya.live", "zenghuiyang@fuya.live", "2112484385@qq.com",
}

const SettleTempCommonMsg = `
	<div style="background-color:#ECECEC; padding: 35px;"> 
	<table cellpadding="0" align="center" style="width: 600px; margin: 0px auto; text-align: left; position: relative; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 5px; border-bottom-left-radius: 5px; font-size: 14px; font-family:微软雅黑, 黑体; line-height: 1.5; box-shadow: rgb(153, 153, 153) 0px 0px 5px; border-collapse: collapse; background-position: initial initial; background-repeat: initial initial;background:#fff;"> 
		 <tbody> 
		 <tr> 
			 <th valign="middle" style="height: 15px; line-height: 15px; padding: 15px 35px; border-bottom-width: 1px; border-bottom-style: solid; border-bottom-color: #505053; background-color: #181821; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;"> 
				 <font face="微软雅黑" size="2" style="color: rgb(255, 255, 255); ">FUYA数据统计服务</font> 
			 </th> 
		 </tr> 
		 <tr> 
			 <td> 
				 <div style="padding:25px 35px 40px; background-color:#fff;"> 
					 <h2 style="margin: 5px 0px; "> 
						 <font color="#FD882E" style="line-height: 20px; "> 
							 <font style="line-height: 22px; " size="3"> 
								 HimiLive本期收支数据-结算中心：</font> 
						 </font> 
					 </h2> 
					 <p style="color: #242830; font-size: 12px;"><b>薪资周期：%v</b><br> 
						 主播结算人数：%v<br> 
						 主播结算总额(IDR)：%v<br> 
						 公会结算人数：%v<br> 
						 公会结算总额(IDR)：%v<br> 
						 每周排行奖励(IDR)：%v<br>
						 结算总额(IDR)：%v<br> 
						 汇率：2200<br> 
						 约合人民币(CNY)：%v<br></p> 
						 <p style="color: #242830; font-size: 12px;"><b>充值周期：%v</b><br> 
						 充值总额(IDR)：%v<br> 
						 汇率：2200<br> 
						 约合人民币(CNY)：%v<br></p>
						 <p style="color: #242830; font-size: 13px;"></p>结算占比：<font color="#ED3851"><b>%v</b></font></p>
						 <p style="color: #242830; font-size: 12px;"> 
						 实际结算流水占比：<font color="#FD882E"><b>%v</b></font><br>
						 </p> <br>
						 <div style="width:700px;margin:0 auto;"> 
						 <div style="padding:10px 10px 0;border-top:1px solid #ccc;color:#747474;margin-bottom:20px;line-height:1.3em;font-size:12px;"> 
							 <p>此邮件为系统邮件，请勿回复<br> 
								 请保管好您的邮箱，避免账号被他人盗用 
							 </p> 
							 <p>FUYA Technology Co., Ltd.</p> 
						 </div> 
					 </div> 
				 </div> 
			 </td> 
		 </tr> 
		 </tbody> 
	 </table> 
 </div>`

const SettleDingMsg = `
<div style="background-color:#ECECEC; padding: 35px;">

<table cellpadding="0" align="center" style="width: 600px; margin: 0px auto; text-align: left; position: relative; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 5px; border-bottom-left-radius: 5px; font-size: 14px; font-family:微软雅黑, 黑体; line-height: 1.5; box-shadow: rgb(153, 153, 153) 0px 0px 5px; border-collapse: collapse; background-position: initial initial; background-repeat: initial initial;background:#fff;">

<tbody>

<tr>

<th valign="middle" style="height: 15px; line-height: 15px; padding: 15px 35px; border-bottom-width: 1px; border-bottom-style: solid; border-bottom-color: #505053; background-color: #181821; border-top-left-radius: 5px; border-top-right-radius: 5px; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px;"><font face="微软雅黑" size="2" style="color: rgb(255, 255, 255); ">FUYA数据统计服务</font></th>

</tr>

<tr>

<td>

<div style="padding:25px 35px 40px; background-color:#fff;">

## <font color="#FD882E" style="line-height: 20px; "><font style="line-height: 22px; " size="3">HimiLive本期收支数据-结算中心：</font></font>

**薪资周期：%v**  
主播结算人数：%v  
主播结算总额(IDR)：%v  
公会结算人数：%v  
公会结算总额(IDR)：%v  
每周排行奖励(IDR)：%v  
结算总额(IDR)：%v  
汇率：2200  
约合人民币(CNY)：%v  

**充值周期：%v**  
充值总额(IDR)：%v  
汇率：2200  
约合人民币(CNY)：%v  

结算占比：<font color="#ED3851">**%v**</font>

实际结算流水占比：<font color="#FD882E">**%v**</font>  

</td>

</tr>

</tbody>

</table>

</div>`

const GameTempMailMsg = `
		<blockquote><p><strong>异常详情</strong><br/>
			游戏名: %v<br/>
			游戏中奖：<font color=red>%v</font><br/>
			游戏流水：<font color=warning>%v</font><br/>
			盈亏：<font color=blue>%v</font><br/>
			利润率：<font color=warning>%.2f&#37;</font><br/>
			请关注此异常。<br/>
			后台地址，请点击：<a href='https://himi-manager.himilive.vip/#/dashboard'>himi-后台管理 </a></p>
		</blockquote>
		<p>&nbsp;</p>`

const ApiErr = `<blockquote><p><strong>异常详情</strong><br/>
			subject: %v<br/>
			host:<font color=red>%v</font><br/>
			api：<font color=blue>%v</font><br/>
			异常详情：<font color=warning>%v</font><br/>
			请关注此异常。<br/>
			后台地址，请点击：<a href='https://himi-manager.himilive.vip/#/dashboard'>himi-后台管理 </a></p>
		</blockquote>
		<p>&nbsp;</p>`

const ServerDownErr = `> **健康检查失败** <br/>
> 服务器名: %v <br/>
> 服务器ip:<font color="green">%v</font>  
> 检查时间：<font color="red">%v</font>  
> 请关注此异常。  
> 后台地址，请点击：[himi-后台管理](https://himi-manager.himilive.vip/#/dashboard)`

const GameDinglMsg = `
> **异常详情**  
> 游戏名: %v  
> 游戏中奖：<font color="red">%v</font>  
> 游戏流水：<font color="warning">%v</font>  
> 盈亏：<font color="blue">%v</font>  
> 利润率：<font color="warning">%v</font>  
> 请关注此异常。  
> 后台地址，请点击：[himi-后台管理](https://himi-manager.himilive.vip/#/dashboard)`

const GiftTempMailMsg = `
		<blockquote><p><strong>异常详情</strong><br/>
			幸运礼物id: %v<br/>
			礼物中奖：<font color=red>%v</font><br/>
			礼物流水：<font color=warning>%v</font><br/>
			盈亏：<font color=blue>%v</font><br/>
			利润率：<font color=warning>%v&#37;</font><br/>
			请关注此异常。<br/>
			后台地址，请点击：<a href='https://himi-manager.himilive.vip/#/dashboard'>himi-后台管理 </a></p>
		</blockquote>
		<p>&nbsp;</p>`

const GiftDingMsg = `
> **异常详情**  
> 幸运礼物id: %v  
> 礼物中奖：<font color="red">%v</font>  
> 礼物流水：<font color="warning">%v</font>  
> 盈亏：<font color="blue">%v</font>  
> 利润率：<font color="warning">%v</font>
> 请关注此异常。  
> 后台地址，请点击：[himi-后台管理](https://himi-manager.himilive.vip/#/dashboard)`

const AgentOfflineDinglMsg = `
> **BCA线下充值上传凭证**  
> 时间: %v  
> 用户id: %v  
> 订单id: %v  
> 等级: %v  
>  ![凭证图片](%v)  
> 请及时审核。  
> 后台地址，请点击：[himi-后台管理](https://himi-manager.himilive.vip/#/dashboard)`

const AgentBalanceExamineMailMsg = `
		<blockquote><p><strong>代理余额校验异常</strong><br/>
			%v点 代理钻石：<font color=red>%v</font><br/>
			%v点 代理钻石：<font color=red>%v</font><br/>
			差额：<font color=blue>%v</font><br/>
			请关注此异常。<br/>
			后台地址，请点击：<a href='https://himi-manager.himilive.vip/#/dashboard'>himi-后台管理 </a></p>
		</blockquote>
		<p>&nbsp;</p>`

const AgentBalanceExamineDingMsg = `
> **代理余额校验异常**  
> %v点 代理钻石：<font color=red>%v</font>  
> %v点 代理钻石：<font color=red>%v</font>  
> 差额：<font color=blue>%v</font>  
> 请关注此异常。  
> 后台地址，请点击：[himi-后台管理 ](https://himi-manager.himilive.vip/#/dashboard)`

const AnchorLiveTimeLongDingMsg = `
> **提醒：直播时长过长**  
> **Peringatan: Durasi siaran langsung terlalu lama** <br/>
> kali ini host <font color=red>%v</font>   live sudah kelebih 8 jam,silakan perhatian status live host<br/>
> 后台地址，请点击：[himi-后台管理 ](https://himi-manager.himilive.vip/#/dashboard)`
