// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorWeekSettlement is the golang structure of table anchor_week_settlement for DAO operations like Where/Data.
type AnchorWeekSettlement struct {
	g.Meta                     `orm:"table:anchor_week_settlement, do:true"`
	Id                         interface{} // 自增Id
	Week                       interface{} // 周，一年中的第几周 例如:202230 2022年的第30周
	GuildId                    interface{} // 公会ID
	CutRate                    interface{} // 公会抽成 设置范围0-100
	UserId                     interface{} // 用户ID
	TotalIncome                interface{} // 总收益 单位分
	LiveTime                   interface{} // 直播时长（s）
	EffectLiveDay              interface{} // 有效出勤天数
	StarLevel                  interface{} // 星级
	IsFirstLive                interface{} // 本周是否首播 0非首播；1首播
	SettleState                interface{} // 结算状态：  1可以结算  达到结算要求的标星主播或有公会的主播  0：兑换为钻石（不可以结算）；
	SurplusIncome              interface{} // 不结算时兑换完钻石后，剩余金币，结算时这里存储0
	ExchangeDiamonds           interface{} // 结算时兑换钻石数，不结算的话这里存储0
	IncomeGear                 interface{} // 金币档位 例如：0-2500000
	GearCondition              interface{} // 档位条件；例如：周1至周3
	RewardFormulaForHost       interface{} // 例如：coin*0.08*50%
	HostRate                   interface{} // 对应中reward_formula_for_host的最后奖励比例，比如：50% *100的整数
	RewardMoneyForHost         interface{} // 主播金币档位奖励给主播 印尼盾 分
	GiftIncome                 interface{} // 礼物收入=reward_formula_for_host 印尼盾 分
	GuildRate                  interface{} // 公会礼物提成比列  *100的整数
	GuildGiftIncome            interface{} // 公会礼物提成收入 印尼盾 分
	WeekRank                   interface{} // 周排行榜 标星主播按照金币又大到小排名
	WeekRankRewardForHost      interface{} // 周榜主播奖励（印尼盾）同时需要按照例如：周1至周3工作几天时长多久来发放，如果排行榜靠前，时长或出勤天未达到同样不发放奖励  印尼盾 分
	WeekRankRewardForGuildRate interface{} // 周榜公会获得提成比率 *100的整数
	WeekRankRewardForGuild     interface{} // 周榜公会获得提成收入 印尼盾 分
	HostTotalIncomeCutBefore   interface{} // 公会抽成前主播总收入 印尼盾 分
	HostTotalIncomeCut         interface{} // 公会从主播收入中抽成  印尼盾 分
	PayCommission              interface{} // 支付手续费、（is_pay_to_guild此数据关联）给公会该值为0，单独打给主播需要扣除手续 印尼盾分
	HostTotalIncome            interface{} // 主播总收入 印尼盾 分
	ActivityRewardForGuild     interface{} // 活动奖励给公会 印尼盾 分 比如分四周奖励    该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会
	LiveRewardForGuild         interface{} // 直播奖励给公会 印尼盾 分 比如首播1小时  该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会
	GuildTotalIncome           interface{} // 公会总收入 印尼盾 分
	IsPayToGuild               interface{} // 工资是否发放给公会 0非 1是
	SettlementBankName         interface{} // 发放账户名
	SettlementBankAccounts     interface{} // 发放银行账号
	SettlementBankPerson       interface{} // 发放银行卡持有人
	IsSettle                   interface{} // 是否已经结算过 0未结算  1结算过
	ClassifyWeekRank           interface{} // 类型周排行
	ClassifyWeekRankType       interface{} // 类型周排行类型 1 新档位 2 中档位 3 高档位
	WaterRank                  interface{} // 周流水排名
	WaterDiamond               interface{} // 周流水钻石
	CreateTime                 interface{} // 创建时间
	UpdateTime                 interface{} // 更新时间
}
