// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GuildWeekSettlement is the golang structure for table guild_week_settlement.
type GuildWeekSettlement struct {
	Id                     int    `json:"id"                     description:""`
	SettlementWeek         int    `json:"settlementWeek"         description:"周 example:202230"`
	GuildId                int64  `json:"guildId"                description:"公会ID"`
	UserId                 int64  `json:"userId"                 description:"用户ID"`
	LiveTime               int    `json:"liveTime"               description:"直播时长（s）"`
	PayCommission          int64  `json:"payCommission"          description:"支付手续费"`
	TotalIncome            int64  `json:"totalIncome"            description:"工会所有主播 工会总收益金币收益 单位分"`
	EffectLiveDay          int    `json:"effectLiveDay"          description:"工会所有主播 总有效出勤天数"`
	GiftIncome             int64  `json:"giftIncome"             description:"工会所有主播 礼物收入 印尼盾 分"`
	GuildGiftIncome        int64  `json:"guildGiftIncome"        description:"工会所有主播 公会礼物提成收入 印尼盾 分"`
	WeekRankRewardForGuild int64  `json:"weekRankRewardForGuild" description:"工会所有主播 周榜公会获得提成收入 印尼盾 分"`
	HostTotalIncomeCut     int64  `json:"hostTotalIncomeCut"     description:"工会所有主播 公会从主播收入中抽成  印尼盾 分"`
	HostTotalIncome        int64  `json:"hostTotalIncome"        description:"工会所有主播 主播总收入 印尼盾 分"`
	ActivityRewardForGuild int64  `json:"activityRewardForGuild" description:"工会所有主播 活动奖励给公会"`
	LiveRewardForGuild     int64  `json:"liveRewardForGuild"     description:"工会所有主播 直播奖励给公会 印尼盾 分"`
	GuildTotalIncome       int64  `json:"guildTotalIncome"       description:"公会总收入 印尼盾 分"`
	SettlementBankName     string `json:"settlementBankName"     description:"发放账户名"`
	SettlementBankAccounts string `json:"settlementBankAccounts" description:"发放银行账号"`
	SettlementBankPerson   string `json:"settlementBankPerson"   description:"发放银行卡持有人"`
	IsSettle               int    `json:"isSettle"               description:"是否已经结算过 0未结算  1结算过"`
	UpStandardHostCount    int64  `json:"upStandardHostCount"    description:"工会所有主播 达标结算主播数"`
	NotStandardHostCount   int64  `json:"notStandardHostCount"   description:"工会所有主播 未达标结算主播数"`
	GuildCutRate           int64  `json:"guildCutRate"           description:"工会抽成"`
	HostTotalIncomeBefore  int64  `json:"hostTotalIncomeBefore"  description:"工会所有主播 未扣除主播总收入 印尼盾 分"`
	IsPayToGuild           int    `json:"isPayToGuild"           description:"工资是否发放给公会 0非 1是"`
	CreateTime             int64  `json:"createTime"             description:"创建时间"`
	UpdateTime             int64  `json:"updateTime"             description:"更新时间"`
}
