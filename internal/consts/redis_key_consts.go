package consts

// 礼物相关key
const (
	/*GiftLuckyGiftDCTPre  = "DCT:Pre:"  //前一个奖池*/
	GiftLuckyGiftDCTPre = "DCTPre" //前一个奖池
	/*GiftLuckyGiftDCTNext = "DCT:Next:" //后一个奖池*/
	GiftLuckyGiftDCTNext = "DCTNext" //后一个奖池

	GiftLuckyGiftFive          = "{LuckyGift}:Five"
	GiftLuckyGiftTen           = "{LuckyGift}:Ten"
	GiftLuckyGiftFifty         = "{LuckyGift}:Fifty"
	GiftLuckyGiftHundred       = "{LuckyGift}:Hundred"
	GiftLuckyGiftFiveHundred   = "{LuckyGift}:FiveHundred"
	GiftLuckyGiftCursor        = "{LuckyGift}:Cursor" //游标
	GiftLuckyGiftCashbackRatio = "{LuckyGift}:Ratio"
	GiftLuckyGiftMakeNum       = "{LuckyGift}:MakeNum"    //幸运礼物生成次数
	GiftLuckyGiftTotalRatio    = "{LuckyGift}:TotalRatio" //幸运礼物生成返奖率总计 这里是整数扩大100倍 比如返奖率是0.78=78

	//新幸运礼物key
	LuckyGiftDSGPre  = "DSGPre"  //前一个奖池
	LuckyGiftDSGNext = "DSGNext" //后一个奖池

	LuckyGiftTen           = "{LuckyGiftV2}:Ten"
	LuckyGiftFifty         = "{LuckyGiftV2}:Fifty"
	LuckyGiftFiveHundred   = "{LuckyGiftV2}:FiveHundred"
	LuckyGiftCursor        = "{LuckyGiftV2}:Cursor" //游标
	LuckyGiftCashbackRatio = "{LuckyGiftV2}:Ratio"
	/*LuckyGiftMakeNum       = "{LuckyGiftV2}:MakeNum"    //幸运礼物生成次数
	LuckyGiftTotalRatio    = "{LuckyGiftV2}:TotalRatio" //幸运礼物生成返奖率总计 这里是整数扩大100倍 比如返奖率是0.78=78*/

	LuckyGiftTenV3           = "{LuckyGiftV3}:Ten"
	LuckyGiftFiftyV3         = "{LuckyGiftV3}:Fifty"
	LuckyGiftFiveHundredV3   = "{LuckyGiftV3}:FiveHundred"
	LuckyGiftFiveThousandV3  = "{LuckyGiftV3}:FiveThousand"
	LuckyGiftCursorV3        = "{LuckyGiftV3}:Cursor"      //游标 用于拥有5000倍的 以同样价格价格作为key
	LuckyGiftTotalCursorV3   = "{LuckyGiftV3}:TotalCursor" //游标 用于10个奖池120000 计算5000倍
	LuckyGiftCashbackRatioV3 = "{LuckyGiftV3}:Ratio"

	//新幸运礼物key end

	RedisGiftSendKey = "Gift:SendKey:" //redis 发送key
	//RedisRoomBattleGiftDiamondsKey 直播pk礼物统计钻石key %v:%v  pkId:userId
	RedisRoomBattleGiftDiamondsKey = "Room:Battle:Gift:Diamonds:"
	//RedisRoomBattleGiftScoreKey 直播pk礼物统计积分key %v:%v  pkId:userId
	RedisRoomBattleGiftScoreKey = "Room:Battle:Gift:Score:"
	//RedisRoomBattleUserRankKey 直播pk房间礼物贡献排行key %v:%v  anchorUserId:pkId
	RedisRoomBattleUserRankKey = "Room:Battle:User:Rank:"
	// GiftHotTopOne 热一礼物置顶
	GiftHotTopOne = "Gift:Hot:Top:One"
	// GiftSenderRankZSetKey 礼物赠送排行key
	GiftSenderRankZSetKey = "Gift:Sender:Rank"
	// GiftReceiverRankZSetKey 礼物接收排行key
	GiftReceiverRankZSetKey = "Gift:Receiver:Rank"
	//RoomWishGiftSendRankingZSetKey 房间心愿礼物发送排行
	RoomWishGiftSendRankingZSetKey  = "Room:Wish:Gift:Send:ZSetKey"
	RoomWishGiftSendRankingCountKey = "Room:Wish:Gift:Send:Count"
	//RoomWishGiftSendRankingReceiveKey 心愿礼物已领取
	RoomWishGiftSendRankingReceiveKey = "Room:Wish:Gift:Send:Receive"
	// RoomInviteLiveMultipleCacheKey 多人房邀请key
	RoomInviteLiveMultipleCacheKey = "Room:Invite:Live:Multiple:CacheKey"
	// RoomInviteAgreeLiveMultipleCacheKey 多人房同意邀请key
	RoomInviteAgreeLiveMultipleCacheKey = "Room:Invite:Agree:Live:Multiple:CacheKey"
	//RedisRoomBattleGiftDiamondsCriticalKey 直播pk礼物统计暴击钻石key %v:%v  pkId:userId
	RedisRoomBattleGiftDiamondsCriticalKey = "Room:Battle:Gift:Critical:Diamonds:"
	// RedisRoomBattleAudioStatesKey 直播pk闭麦key
	RedisRoomBattleAudioStatesKey = "Redis:Room:Battle:Audio:States:Key:"
	// RedisRoomBattleFirstGiftKey 首杀是否已使用key
	RedisRoomBattleFirstGiftKey = "Redis:Room:Battle:First:Gift:Key"
	//GiftActivityDayTimeSendKey 判断活动日时间
	GiftActivityDayTimeSendKey = "Gift:Activity:Day:Time:Send:Key"
	//GiftActivityWeekTimeSendKey 判断活动周时间
	GiftActivityWeekTimeSendKey = "Gift:Activity:Week:Time:Send:Key"
	//GiftActivityWeekGiftTop1SendKey 礼物日榜top1用户
	GiftActivityWeekGiftTop1SendKey = "Gift:Activity:Week:GiftTop1:Send:Key"
	//UserActivityGiftRankSystemMsgString 用户活动礼物系统消息通知发送礼物 key+20230921:礼物Id:userId:名次(5_10/3_5/1_3)
	UserActivityGiftRankSystemMsgString       = "SystemMsg:GiftActivity:Rank:"
	UserActivityGiftRankSystemMsgStringEXPIRE = 2 * 24 * 3600
	//AnchorWeeklyRankNoReachedSystemMsgString 主播周榜未达标消息 key+20230921:userId
	AnchorWeeklyRankNoReachedSystemMsgString       = "SystemMsg:WeeklyRank:NoReached:"
	AnchorWeeklyRankNoReachedSystemMsgStringEXPIRE = 2 * 24 * 3600
	//AnchorWeeklyRankFirstReadSystemMsgString 主播首次达标对应榜单记录 key+周(202343):userId:(对应榜)1为新秀/2为推荐/3为优质
	AnchorWeeklyRankFirstReadSystemMsgString       = "SystemMsg:WeeklyRank:FirstRead:"
	AnchorWeeklyRankFirstReadSystemMsgStringEXPIRE = 8 * 24 * 3600
	//AnchorWeeklyRankFirstSystemMsgString 主播第一名通知 key+周(202343):userId:(对应榜)1为新秀/2为推荐/3为优质
	AnchorWeeklyRankFirstSystemMsgString       = "SystemMsg:WeeklyRank:First:"
	AnchorWeeklyRankFirstSystemMsgStringEXPIRE = 8 * 24 * 3600
	//AnchorWeeklyRankLastCloseLiveString 主播上次发送消息记录当时的榜单及名次 key+周(202343):userId ，  value：(对应榜)1为新秀/2为推荐/3为优质_名次
	AnchorWeeklyRankLastCloseLiveString       = "SystemMsg:WeeklyRank:LastCloseLive:"
	AnchorWeeklyRankLastCloseLiveStringEXPIRE = 8 * 24 * 3600
	//AnchorGiftWaterDiamondSendKey 主播礼物流水延时发送key
	AnchorGiftWaterDiamondSendKey  = "Anchor:Gift:Water:Diamond:SendKey:"
	AnchorQuestWaterDiamondSendKey = "Anchor:Quest:Water:Diamond:SendKey:"
	//UserGiftWaterDiamondKey 用户流水增量缓存key
	UserGiftWaterDiamondKey = "User:Gift:Water:Diamond:"
)
const ( //年终活动排行key
	ChristmasGiftListRankZSetKey       = "Christmas:Gift:List:Rank:ZSetKey"
	ChristmasGiftListRankDetailZSetKey = "Christmas:Gift:List:Rank:Detail:ZSetKey"

	ChristmasUserGiftListRankZSetKey = "Christmas:User:Gift:List:Rank:Detail:ZSetKey"
)

// 直播房间相关key
const (
	RedisBattleEndKey          = "Room:Battle:End:"               //redis 结束pk房间锁 %v pk id
	RoomLiveBattleNotAccept    = "Room:Battle:Not:Accept:"        //当前直播用户不接受pk %v userId
	RoomRealViewer             = "Room:Viewer:Real:"              //真实用户
	RoomRobotViewer            = "Room:Viewer:Robot:"             //单个房间内的机器人
	RoomRobotAllIn             = "Room:Viewer:RobotAll:All"       //所有房间内的机器人
	RoomRobotExisted           = "Room:Viewer:RobotExisted:"      //单个房间内的进入过的机器人
	RoomFansSet                = "Room:Fans:"                     //直播间新增粉丝
	RoomViewerTotalSet         = "Room:Viewer:Total:"             //房间观众累计 （只新增不删除）
	RoomShutUpZSet             = "Room:ShutUp:"                   //直播间永久禁言 Key=Room:ShutUp:主播Id
	RoomKickOutZSet            = "Room:KickOut:"                  //直播间踢出 Key=Room:KickOut:主播Id
	RoomHeartBeat              = "Room:User:HeartBeat:"           //房间心跳缓存
	RoomHotList                = "Room:Author:Hot:All"            //热门榜单
	RoomNewList                = "Room:Author:New:All"            //最新榜单
	RoomHeartBeatTTL           = 300                              //
	UserRobotInfo              = "User:Robot:Info:"               //
	RoomHotLiveTime            = "Room:Author:Hot:LiveTime:"      //主播热门开播时间数据
	RoomHotLiveAward           = "Room:Author:Hot:LiveAward:"     //主播热门开播奖励数据
	RoomUserBan                = "Room:User:LiveBan:"             //开播黑名单
	RoomHotRotate              = "Room:Author:Hot:Rotate"         //热门轮播榜单
	RoomShortScore             = "Room:User:ShortScore:"          //房间临时积分调整结构体
	RoomTopOneUserIdKey        = "Room:Top:One:UserId:Key:"       //主播房间第一名的id
	AnchorLabelSetTime         = "Anchor:Label:Set:Time:"         //主播标签设置时间
	RechargePointsRecordZSet   = "Recharge:Points:Record:ZSet"    //用户充值排名集合
	RoomLiveAnonymityKey       = "Room:Live:Anonymity:Key"        //用户房间匿名key
	RoomLivePlayGameKey        = "Room:Live:Play:Game:"           //用户房间游戏key
	SmallBlackRoomUserIdSetKey = "Small:Black:Room:UserId:SetKey" //加入过用户的小黑屋
	RoomMsgEndMsgSendKey       = "Room:Msg:EndMsg:SendKey"        //房间官方消息
	BattleRewardOnlyKey        = "Room:Battle:Only:"              //pk返奖唯一key
	RoomLeaveLastUnix          = "Room:Leave:LastUnix"            // 用户进入或离开直播间的时间戳
	RoomLeaveCount             = "Room:Leave:Count"               // 用户连续离开直播间且时间小于30s的次数
	RoomAnchorLastPopupUnix    = "Room:Author:Popup:LastUnix"     // 主播上次好评引导弹窗时间戳
	RoomAnchorPopupCount       = "Room:Author:Popup:Count"        // 主播好评引导弹窗次数
	RoomMultipleSoundControl   = "Room:Multiple:Sound:Control:"   // 主播多人房全员声音控制
	RoomMultipleVideoControl   = "Room:Multiple:Video:Control:"   // 主播多人房全员视频控制
	RoomAnchorMixStreamId      = "Room:Anchor:Mix:StreamId"       // 主播预览混流id
	RoomAnchorLiveRoomPicture  = "Room:Anchor:Live:Room:Picture"  // 主播开播封面
)

// 用户相关缓存key
const (
	HeartbeatKey = "User:Heartbeat:"
	//UserTokenKey 用户token对比:%s userId
	UserTokenKey = "User:APP:TOKEN:"
	//AgentUserTokenKey 用户token对比:%s userId
	AgentUserTokenKey = "Agent:User:APP:TOKEN:"
	//RongTokenKey 用户融云token缓存:%s userId
	RongTokenKey = "User:RongCloud:TOKEN:"
	//UserLiveHeartBeatDay 用户当天直播时间心跳次数、用来计算直播时长  :%v:v% yyyy-mm-dd:userId
	UserLiveHeartBeatDay          = "User:Live:Heart:Beat:Day:"
	UserLiveHeartBeatRoomIdRepeat = "User:Live:Heart:Beat:RoomId:Repeat"
	//ViewerHeartBeatUserId 用户房间观看直播心跳
	ViewerHeartBeatUserId = "Viewer:Heart:Beat:UserId:"
	// ViewerDayHeartBeatUserId 用户每日观看直播心跳
	ViewerDayHeartBeatUserId = "Viewer:Day:Heart:Beat:UserId:"
	//UserMonthActiveKey 用户活跃
	UserMonthActiveKey = "User:Active:Month:"
	//UserWeekActiveKey 用户活跃
	UserWeekActiveKey = "User:Active:Week:"
	//UserDayActiveKey 用户活跃
	UserDayActiveKey = "User:Active:Day:"
	//UserRegCacheCaptcha 游戏处理key
	UserRegCacheCaptcha = "User:RegCache:Captcha:"
)

// 订单相关缓存
const (
	OrderTradeLock                       = "Order:Trade:"                      //redis 发送key
	OrderMqConsumedStr                   = "Order:Consumed:"                   //记录消费过的订单消息 避免重复消费
	OrderMqConsumedStrExpire             = 3600 * 2                            //消费消息记录2小时
	OrderRoomDiamondsRankZSetKey         = "Order:Room:Diamonds:Rank"          //房间充值订单钻石排行榜 1主播 2用户
	OrderRoomAnchorContributeRankZSetKey = "Order:Room:Anchor:Contribute:Rank" //房间充值主播贡献榜

)

// 数据库相关缓存
const (
	MysqlUserExperienceRank         = "Mysql:User_Experience:Rank"
	MysqlUserExperience             = "Mysql:User_Experience:UserId:"
	MysqlUserInfo                   = "Mysql:User_Info:UserId:"
	MysqlUserInfoRobotUserId        = "Mysql:User_Info:Robot:All:UserIds"
	MysqlUserRankRuleListAll        = "Mysql:User_Rank_Rule:List:All"
	MysqlRoomInfoByRoomId           = "Mysql:Room_Info:RoomId:"
	MysqlShortUrlByCode             = "Mysql:Short_Url:Code:"
	MysqlInvitedInfoByUid           = "Mysql:Invited_Info:UserId:"
	MysqlInvitedInfoByMid           = "Mysql:Invited_Info:SMID:"
	MysqlInvitedInfoByOwnerUserId   = "Mysql:Invited_Info:OwnerUserId:"
	MysqlBalanceRecordInvitedReward = "Mysql:Balance_Record:InvitedReward:"
	MysqlGuildAnchorByUid           = "Mysql:Guild_Anchor:UserId:"
	MysqlAreaByTitle                = "Mysql:Area:Title::"
	MysqlCommodityGiftCache         = "Mysql:Commodity:Gift:Cache:List"
	MysqlGuildPullInfoByPulledId    = "Mysql:GuildPullInfo:PulledId:"
	MysqlGuildInfoById              = "Mysql:GuildInfo:Id:"
	MysqlGuildMasterInfoById        = "Mysql:GuildInfo:Master:Id:"
	MysqlAnchorGuildByUserId        = "Mysql:AnchorGuild:UserId:"
	MysqlQuestInfoByUserId          = "Mysql:QuestInfo:UserId:"
	MysqlBanByDeviceId              = "Mysql:Ban:DeviceId:"
)

// 4周活动
const (
	ActivityFourWeekZset       = "Activity:FourWeek:" //4周活动 key+userid   score:1,member:effectDay,score:100,member:effectSecond, score:1(1发放过，0未发放过),member:isReward
	ActivityFourWeekZsetEXPIRE = 50 * 24 * 3600       //活动数据缓存50天
	// FourWeekIsReward isReward 是否发放过奖励 1发过  0未发放过
	FourWeekIsReward = "isReward"
	// FourWeekEffectDay effectDay 有效天直播天
	FourWeekEffectDay = "effectDay"
	// FourWeekEffectSecond effectSecond 有效直播秒
	FourWeekEffectSecond = "effectSecond"
	// FourWeekRewardDate 发放时间 score 保存当日的时间戳
	FourWeekRewardDate = "RewardDate"
	// AnchorWeekRankDiamondRewardKey 主播周榜结算钻石奖励key
	AnchorWeekRankDiamondRewardKey = "Anchor:WeekRank:Diamond:RewardKey"
)

// 首播活动
const (
	ActivityFirstLiveRewardString       = "Activity:FirstLive:" //首播活动 key+userId   1或>1标识成功
	ActivityFirstLiveRewardStringEXPIRE = 50 * 24 * 3600        //活动数据缓存50天
)

// 公会拉公会活动
const (
	ActivityGuildPullGuildRewardString       = "Activity:GuildPullGuild:" // 公会拉公会活动 key+userId   1或>1标识成功
	ActivityGuildPullGuildRewardStringEXPIRE = 50 * 24 * 3600             //活动数据缓存50天
)

// 钻石转盘活动
const (
	ActivityDiamondsTurntableDailyWatchTime = "Activity:Diatable:Daily:WatchTime:" // zset用户每日活动观看时长 %v:v% yyyy-mm-dd
	ActivityDiamondsTurntableDailyCount     = "Activity:Diatable:Daily:Count:"     // zset用户每日抽奖次数 %v:v% yyyy-mm-dd
	ActivityDiamondsTurntableDailyRoom      = "Activity:Diatable:Daily:Room:"      // 统计用户有效观看时间 %v:v% yyyy-mm-dd:userId
	ActivityDiamondsTurntablePool           = "Activity:Diatable:Pool"             // list 配置奖池 期数
	ActivityDiamondsTurntableNum            = "Activity:Diatable:Num"              // 奖池 期数
	ActivityDiamondsTurntableUserNX         = "Activity:Diatable:User:nx:"         // 用户奖池锁
	ActivityDiamondsTurntableDailyEXPIRE    = 2 * 24 * 3600                        // 每日数据缓存两天 用于处理报表程序
	ActivityDiamondsTurntableValidTime      = 5 * 60 / 30                          // 钻石转盘活动的有效观看时间
	ActivityDiamondsTurntableLimitTIme      = 5 * 60 / 30                          // 钻石转盘活动的生效观看时间
)

// 活动转盘活动
const (
	ComplexPrizePool    = "Complex:Prize:Pool"   // list 配置抽奖奖池 :v% :activityType
	ComplexPrizePoolNum = "Complex:Diatable:Num" // 奖池 期数 :v% :activityType
)

const (
	// UserIdFirstPayKey 用户首充
	UserIdFirstPayKey = "UserId:First:Pay:"
	// OrderDiscountProductIdCacheKey 折扣订单返回订单缓存
	OrderDiscountProductIdCacheKey = "Order:Discount:ProductId:Cache:Key:"
)

const (
	RedisLockFeedLike = "RedisLock:FeedLike:" //redis 锁
)

const (
	IpCacheKey      = "Ip:Cache:Key:" //缓存ip转换地址
	WarnDingDingKey = "Warn:DingDing:Key"
)

const (
	StatisticHotAward                    = "Statistic:Hot:LiveAward:"                         //热门奖励统计
	StatisticInviteAward                 = "Statistic:User:Invite:"                           //邀请奖励统计
	StatisticDiamondsConsumeByLuck       = "Statistic:Diamonds:Consume:Luck:"                 //幸运礼物钻石消耗统计
	StatisticDiamondsConsumeByNormal     = "Statistic:Diamonds:Consume:Normal:"               //普通礼物钻石消耗统计
	StatisticDiamondsGameExpend          = "Statistic:Diamonds:Game:Expend:"                  //游戏钻石消耗统计
	StatisticDiamondsGameIncome          = "Statistic:Diamonds:Game:Income:"                  //游戏钻石中奖收益统计
	StatisticDiamondsAnchorWeekGameWater = "Statistic:Diamonds:Anchor:Week:Game:Income:Water" //主播周流水游戏钻石
	StatisticCoinsIncomeByNormal         = "Statistic:Coins:Income:Normal:"                   //金币收益统计
	StatisticCoinsIncomeByLuck           = "Statistic:Coins:Income:Luck:"                     //金币收益统计
)

const (
	GameHandierKey          = "Game:Handier:Key:"                        //游戏处理key
	CocosConsumeUniqueId    = "Game:Handier:Key:CocosConsumeUniqueId"    //游戏处理key
	CocosSettlementUniqueId = "Game:Handier:Key:CocosSettlementUniqueId" //游戏处理key
)

const (
	FeedVoteOptionCacheKey = "Feed:Vote:Option:Cache:Key:"
)

const (
	//RedEnvelopeDetailCacheKey 红包钻石详情
	RedEnvelopeDetailCacheKey = "Red:Envelope:Detail:Cache:Key"
	//RedEnvelopeUserDetailCacheKey 房间红包用户钻石详情
	RedEnvelopeUserDetailCacheKey = "Red:Envelope:Detail:User:Cache:Key"
	//RedEnvelopeUserSettlementCacheKey 已结算标记
	RedEnvelopeUserSettlementCacheKey = "Red:Envelope:Detail:Settlement:Cache:Key"
	//RedEnvelopeCountCacheKey 红包已抢数量
	RedEnvelopeCountCacheKey = "Red:Envelope:Count:Cache:Key"
)

const (
	// ActivityConfigUserCache 用户活动配置缓存
	ActivityConfigUserCache         = "Activity:Config:User:Cache"
	ActivityConfigIsPushCache       = "Activity:Config:Is:Push:Cache"
	GetWeekRankListByTypeCacheKey   = "WeekRank:List:ByType:CacheKey"
	GetWeekRankListByTypeV2CacheKey = "WeekRank:List:ByType:V2:CacheKey"
	AnchorLiveWarnCacheKey          = "Anchor:LiveWarn:CacheKey"
)

const (
	//NoblePrivilegeStatesCacheKey 特权设置状态key :%v:%v userId privilegeType
	NoblePrivilegeStatesCacheKey = "Noble:Privilege:States:CacheKey"
	//RoomWarnCacheKey 房间提醒key
	RoomWarnCacheKey = "Room:Warn:CacheKey"
)

const (
	UserLoginPhoneSmsCode  = "User:Login:Phone:Sms:Code:"
	UserPasswordErrorCount = "User:Password:Error:Count:"
)

const (
	StarLevelResetWatchingString       = "StarLevel:Reset:Watching:" //星级调整 观察期
	StarLevelResetWatchingStringEXPIRE = 8 * 24 * 3600               //观察期8天
)

const (
	FeedUserLastWatchOfficialString       = "Feed:User:LastWatch:Official:" //记录用户最后一次查看官方动态时间  +userId
	FeedUserLastWatchOfficialStringEXPIRE = 2 * 24 * 3600
	FeedTopicWatchedUserSet               = "Feed:Topic:WatchedUser:" //记录话题查看人数  话题Id
)

const (
	WhiteListUser    = "White:List:User"    //白名单
	WhiteGameUser    = "White:Game:User"    //游戏白名单
	WhiteIosGameUser = "White:IosGame:User" //ios游戏白名单
)

// 任务活跃值缓存
const (
	DailyActiveValueCacheKey            = "Daily:Active:Value:Cache:Key"
	WeekActiveValueCacheKey             = "Week:Active:Value:Cache:Key"
	UserSignInActiveValueCacheKey       = "User:SignIn:Active:Value:Cache:Key"
	AnchorFirstQuestActiveValueCacheKey = "Anchor:FirstQuest:Active:Value:Cache:Key"
	QuestIsSendWarnCacheKey             = "Quest:IsSend:Warn:Cache:Key"
	QuestRankIsSendWarnCacheKey         = "Quest:Rank:IsSend:Warn:Cache:Key"
	RedisQuestReceiveKey                = "Redis:Quest:Receive:Key"      //redis 领取奖励 %v
	RedisSendRoomActivityKey            = "Redis:Send:Room:Activity:Key" //redis 房间看播延时加载 %v
)

const (
	FilterConfig  = "Filter:Config:"
	LuckyGiftWarn = "Lucky:Gift:Warn" //幸运礼物警告
	AgentBalance  = "Agent:Balance:"
	GiftLanguage  = "Gift:Language:"
	DayRankReward = "Day:Rank:Reward:" //每日排行奖励
)

const (
	// BattleMatchFatiguedCacheKey 天梯赛疲劳
	BattleMatchFatiguedCacheKey = "Battle:Match:Fatigued:CacheKey"
	// BattleMatchPropsUseCacheKey pk赛道具使用缓存
	BattleMatchPropsUseCacheKey = "Battle:Match:Props:Use:CacheKey"
	//BattleMatchSeasonCacheKey pk赛赛季缓存
	BattleMatchSeasonCacheKey = "Battle:Match:Season:CacheKey"
	//BattleMatchFirstRankCacheKey pk赛季最高档位
	BattleMatchFirstRankCacheKey = "Battle:Match:First:Rank:CacheKey"
	// BattleMatchPenaltyEndTimeCacheKey 天梯赛惩罚时间
	BattleMatchPenaltyEndTimeCacheKey = "Battle:Match:Penalty:EndTime:CacheKey"
	//BattleMatchUserIdCacheKey pk用户白名单存储key
	BattleMatchUserIdCacheKey = "Battle:Match:UserId:CacheKey"
)
