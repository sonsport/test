// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GuildWeekSettlement is the golang structure of table guild_week_settlement for DAO operations like Where/Data.
type GuildWeekSettlement struct {
	g.Meta                 `orm:"table:guild_week_settlement, do:true"`
	Id                     interface{} //
	SettlementWeek         interface{} // 周 example:202230
	GuildId                interface{} // 公会ID
	UserId                 interface{} // 用户ID
	LiveTime               interface{} // 直播时长（s）
	PayCommission          interface{} // 支付手续费
	TotalIncome            interface{} // 工会所有主播 工会总收益金币收益 单位分
	EffectLiveDay          interface{} // 工会所有主播 总有效出勤天数
	GiftIncome             interface{} // 工会所有主播 礼物收入 印尼盾 分
	GuildGiftIncome        interface{} // 工会所有主播 公会礼物提成收入 印尼盾 分
	WeekRankRewardForGuild interface{} // 工会所有主播 周榜公会获得提成收入 印尼盾 分
	HostTotalIncomeCut     interface{} // 工会所有主播 公会从主播收入中抽成  印尼盾 分
	HostTotalIncome        interface{} // 工会所有主播 主播总收入 印尼盾 分
	ActivityRewardForGuild interface{} // 工会所有主播 活动奖励给公会
	LiveRewardForGuild     interface{} // 工会所有主播 直播奖励给公会 印尼盾 分
	GuildTotalIncome       interface{} // 公会总收入 印尼盾 分
	SettlementBankName     interface{} // 发放账户名
	SettlementBankAccounts interface{} // 发放银行账号
	SettlementBankPerson   interface{} // 发放银行卡持有人
	IsSettle               interface{} // 是否已经结算过 0未结算  1结算过
	UpStandardHostCount    interface{} // 工会所有主播 达标结算主播数
	NotStandardHostCount   interface{} // 工会所有主播 未达标结算主播数
	GuildCutRate           interface{} // 工会抽成
	HostTotalIncomeBefore  interface{} // 工会所有主播 未扣除主播总收入 印尼盾 分
	IsPayToGuild           interface{} // 工资是否发放给公会 0非 1是
	CreateTime             interface{} // 创建时间
	UpdateTime             interface{} // 更新时间
}
