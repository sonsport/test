// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorWeekSettlementDao is the data access object for table anchor_week_settlement.
type AnchorWeekSettlementDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns AnchorWeekSettlementColumns // columns contains all the column names of Table for convenient usage.
}

// AnchorWeekSettlementColumns defines and stores column names for table anchor_week_settlement.
type AnchorWeekSettlementColumns struct {
	Id                         string // 自增Id
	Week                       string // 周，一年中的第几周 例如:202230 2022年的第30周
	GuildId                    string // 公会ID
	CutRate                    string // 公会抽成 设置范围0-100
	UserId                     string // 用户ID
	TotalIncome                string // 总收益 单位分
	LiveTime                   string // 直播时长（s）
	EffectLiveDay              string // 有效出勤天数
	StarLevel                  string // 星级
	IsFirstLive                string // 本周是否首播 0非首播；1首播
	SettleState                string // 结算状态：  1可以结算  达到结算要求的标星主播或有公会的主播  0：兑换为钻石（不可以结算）；
	SurplusIncome              string // 不结算时兑换完钻石后，剩余金币，结算时这里存储0
	ExchangeDiamonds           string // 结算时兑换钻石数，不结算的话这里存储0
	IncomeGear                 string // 金币档位 例如：0-2500000
	GearCondition              string // 档位条件；例如：周1至周3
	RewardFormulaForHost       string // 例如：coin*0.08*50%
	HostRate                   string // 对应中reward_formula_for_host的最后奖励比例，比如：50% *100的整数
	RewardMoneyForHost         string // 主播金币档位奖励给主播 印尼盾 分
	GiftIncome                 string // 礼物收入=reward_formula_for_host 印尼盾 分
	GuildRate                  string // 公会礼物提成比列  *100的整数
	GuildGiftIncome            string // 公会礼物提成收入 印尼盾 分
	WeekRank                   string // 周排行榜 标星主播按照金币又大到小排名
	WeekRankRewardForHost      string // 周榜主播奖励（印尼盾）同时需要按照例如：周1至周3工作几天时长多久来发放，如果排行榜靠前，时长或出勤天未达到同样不发放奖励  印尼盾 分
	WeekRankRewardForGuildRate string // 周榜公会获得提成比率 *100的整数
	WeekRankRewardForGuild     string // 周榜公会获得提成收入 印尼盾 分
	HostTotalIncomeCutBefore   string // 公会抽成前主播总收入 印尼盾 分
	HostTotalIncomeCut         string // 公会从主播收入中抽成  印尼盾 分
	PayCommission              string // 支付手续费、（is_pay_to_guild此数据关联）给公会该值为0，单独打给主播需要扣除手续 印尼盾分
	HostTotalIncome            string // 主播总收入 印尼盾 分
	ActivityRewardForGuild     string // 活动奖励给公会 印尼盾 分 比如分四周奖励    该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会
	LiveRewardForGuild         string // 直播奖励给公会 印尼盾 分 比如首播1小时  该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会
	GuildTotalIncome           string // 公会总收入 印尼盾 分
	IsPayToGuild               string // 工资是否发放给公会 0非 1是
	SettlementBankName         string // 发放账户名
	SettlementBankAccounts     string // 发放银行账号
	SettlementBankPerson       string // 发放银行卡持有人
	IsSettle                   string // 是否已经结算过 0未结算  1结算过
	ClassifyWeekRank           string // 类型周排行
	ClassifyWeekRankType       string // 类型周排行类型 1 新档位 2 中档位 3 高档位
	WaterRank                  string // 周流水排名
	WaterDiamond               string // 周流水钻石
	CreateTime                 string // 创建时间
	UpdateTime                 string // 更新时间
}

// anchorWeekSettlementColumns holds the columns for table anchor_week_settlement.
var anchorWeekSettlementColumns = AnchorWeekSettlementColumns{
	Id:                         "id",
	Week:                       "week",
	GuildId:                    "guild_id",
	CutRate:                    "cut_rate",
	UserId:                     "user_id",
	TotalIncome:                "total_income",
	LiveTime:                   "live_time",
	EffectLiveDay:              "effect_live_day",
	StarLevel:                  "star_level",
	IsFirstLive:                "is_first_live",
	SettleState:                "settle_state",
	SurplusIncome:              "surplus_income",
	ExchangeDiamonds:           "exchange_diamonds",
	IncomeGear:                 "income_gear",
	GearCondition:              "gear_condition",
	RewardFormulaForHost:       "reward_formula_for_host",
	HostRate:                   "host_rate",
	RewardMoneyForHost:         "reward_money_for_host",
	GiftIncome:                 "gift_income",
	GuildRate:                  "guild_rate",
	GuildGiftIncome:            "guild_gift_income",
	WeekRank:                   "week_rank",
	WeekRankRewardForHost:      "week_rank_reward_for_host",
	WeekRankRewardForGuildRate: "week_rank_reward_for_guild_rate",
	WeekRankRewardForGuild:     "week_rank_reward_for_guild",
	HostTotalIncomeCutBefore:   "host_total_income_cut_before",
	HostTotalIncomeCut:         "host_total_income_cut",
	PayCommission:              "pay_commission",
	HostTotalIncome:            "host_total_income",
	ActivityRewardForGuild:     "activity_reward_for_guild",
	LiveRewardForGuild:         "live_reward_for_guild",
	GuildTotalIncome:           "guild_total_income",
	IsPayToGuild:               "is_pay_to_guild",
	SettlementBankName:         "settlement_bank_name",
	SettlementBankAccounts:     "settlement_bank_accounts",
	SettlementBankPerson:       "settlement_bank_person",
	IsSettle:                   "is_settle",
	ClassifyWeekRank:           "classify_week_rank",
	ClassifyWeekRankType:       "classify_week_rank_type",
	WaterRank:                  "water_rank",
	WaterDiamond:               "water_diamond",
	CreateTime:                 "create_time",
	UpdateTime:                 "update_time",
}

// NewAnchorWeekSettlementDao creates and returns a new DAO object for table data access.
func NewAnchorWeekSettlementDao() *AnchorWeekSettlementDao {
	return &AnchorWeekSettlementDao{
		group:   "default",
		table:   "anchor_week_settlement",
		columns: anchorWeekSettlementColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnchorWeekSettlementDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AnchorWeekSettlementDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AnchorWeekSettlementDao) Columns() AnchorWeekSettlementColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AnchorWeekSettlementDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnchorWeekSettlementDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AnchorWeekSettlementDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
