// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AnchorWeekSettlement is the golang structure for table anchor_week_settlement.
type AnchorWeekSettlement struct {
	Id                         uint64 `json:"id"                         description:"自增Id"`
	Week                       int    `json:"week"                       description:"周，一年中的第几周 例如:202230 2022年的第30周"`
	GuildId                    int64  `json:"guildId"                    description:"公会ID"`
	CutRate                    uint   `json:"cutRate"                    description:"公会抽成 设置范围0-100"`
	UserId                     int64  `json:"userId"                     description:"用户ID"`
	TotalIncome                int64  `json:"totalIncome"                description:"总收益 单位分"`
	LiveTime                   int    `json:"liveTime"                   description:"直播时长（s）"`
	EffectLiveDay              int    `json:"effectLiveDay"              description:"有效出勤天数"`
	StarLevel                  int    `json:"starLevel"                  description:"星级"`
	IsFirstLive                int    `json:"isFirstLive"                description:"本周是否首播 0非首播；1首播"`
	SettleState                int    `json:"settleState"                description:"结算状态：  1可以结算  达到结算要求的标星主播或有公会的主播  0：兑换为钻石（不可以结算）；"`
	SurplusIncome              uint64 `json:"surplusIncome"              description:"不结算时兑换完钻石后，剩余金币，结算时这里存储0"`
	ExchangeDiamonds           uint64 `json:"exchangeDiamonds"           description:"结算时兑换钻石数，不结算的话这里存储0"`
	IncomeGear                 string `json:"incomeGear"                 description:"金币档位 例如：0-2500000"`
	GearCondition              string `json:"gearCondition"              description:"档位条件；例如：周1至周3"`
	RewardFormulaForHost       string `json:"rewardFormulaForHost"       description:"例如：coin*0.08*50%"`
	HostRate                   int    `json:"hostRate"                   description:"对应中reward_formula_for_host的最后奖励比例，比如：50% *100的整数"`
	RewardMoneyForHost         int64  `json:"rewardMoneyForHost"         description:"主播金币档位奖励给主播 印尼盾 分"`
	GiftIncome                 uint64 `json:"giftIncome"                 description:"礼物收入=reward_formula_for_host 印尼盾 分"`
	GuildRate                  int    `json:"guildRate"                  description:"公会礼物提成比列  *100的整数"`
	GuildGiftIncome            uint64 `json:"guildGiftIncome"            description:"公会礼物提成收入 印尼盾 分"`
	WeekRank                   int    `json:"weekRank"                   description:"周排行榜 标星主播按照金币又大到小排名"`
	WeekRankRewardForHost      int64  `json:"weekRankRewardForHost"      description:"周榜主播奖励（印尼盾）同时需要按照例如：周1至周3工作几天时长多久来发放，如果排行榜靠前，时长或出勤天未达到同样不发放奖励  印尼盾 分"`
	WeekRankRewardForGuildRate int    `json:"weekRankRewardForGuildRate" description:"周榜公会获得提成比率 *100的整数"`
	WeekRankRewardForGuild     int64  `json:"weekRankRewardForGuild"     description:"周榜公会获得提成收入 印尼盾 分"`
	HostTotalIncomeCutBefore   uint64 `json:"hostTotalIncomeCutBefore"   description:"公会抽成前主播总收入 印尼盾 分"`
	HostTotalIncomeCut         int64  `json:"hostTotalIncomeCut"         description:"公会从主播收入中抽成  印尼盾 分"`
	PayCommission              uint64 `json:"payCommission"              description:"支付手续费、（is_pay_to_guild此数据关联）给公会该值为0，单独打给主播需要扣除手续 印尼盾分"`
	HostTotalIncome            uint64 `json:"hostTotalIncome"            description:"主播总收入 印尼盾 分"`
	ActivityRewardForGuild     int64  `json:"activityRewardForGuild"     description:"活动奖励给公会 印尼盾 分 比如分四周奖励    该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会"`
	LiveRewardForGuild         int64  `json:"liveRewardForGuild"         description:"直播奖励给公会 印尼盾 分 比如首播1小时  该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会"`
	GuildTotalIncome           int64  `json:"guildTotalIncome"           description:"公会总收入 印尼盾 分"`
	IsPayToGuild               int    `json:"isPayToGuild"               description:"工资是否发放给公会 0非 1是"`
	SettlementBankName         string `json:"settlementBankName"         description:"发放账户名"`
	SettlementBankAccounts     string `json:"settlementBankAccounts"     description:"发放银行账号"`
	SettlementBankPerson       string `json:"settlementBankPerson"       description:"发放银行卡持有人"`
	IsSettle                   int    `json:"isSettle"                   description:"是否已经结算过 0未结算  1结算过"`
	ClassifyWeekRank           int    `json:"classifyWeekRank"           description:"类型周排行"`
	ClassifyWeekRankType       int    `json:"classifyWeekRankType"       description:"类型周排行类型 1 新档位 2 中档位 3 高档位"`
	WaterRank                  int    `json:"waterRank"                  description:"周流水排名"`
	WaterDiamond               int    `json:"waterDiamond"               description:"周流水钻石"`
	CreateTime                 int64  `json:"createTime"                 description:"创建时间"`
	UpdateTime                 int64  `json:"updateTime"                 description:"更新时间"`
}
