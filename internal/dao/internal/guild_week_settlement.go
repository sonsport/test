// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GuildWeekSettlementDao is the data access object for table guild_week_settlement.
type GuildWeekSettlementDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns GuildWeekSettlementColumns // columns contains all the column names of Table for convenient usage.
}

// GuildWeekSettlementColumns defines and stores column names for table guild_week_settlement.
type GuildWeekSettlementColumns struct {
	Id                     string //
	SettlementWeek         string // 周 example:202230
	GuildId                string // 公会ID
	UserId                 string // 用户ID
	LiveTime               string // 直播时长（s）
	PayCommission          string // 支付手续费
	TotalIncome            string // 工会所有主播 工会总收益金币收益 单位分
	EffectLiveDay          string // 工会所有主播 总有效出勤天数
	GiftIncome             string // 工会所有主播 礼物收入 印尼盾 分
	GuildGiftIncome        string // 工会所有主播 公会礼物提成收入 印尼盾 分
	WeekRankRewardForGuild string // 工会所有主播 周榜公会获得提成收入 印尼盾 分
	HostTotalIncomeCut     string // 工会所有主播 公会从主播收入中抽成  印尼盾 分
	HostTotalIncome        string // 工会所有主播 主播总收入 印尼盾 分
	ActivityRewardForGuild string // 工会所有主播 活动奖励给公会
	LiveRewardForGuild     string // 工会所有主播 直播奖励给公会 印尼盾 分
	GuildTotalIncome       string // 公会总收入 印尼盾 分
	SettlementBankName     string // 发放账户名
	SettlementBankAccounts string // 发放银行账号
	SettlementBankPerson   string // 发放银行卡持有人
	IsSettle               string // 是否已经结算过 0未结算  1结算过
	UpStandardHostCount    string // 工会所有主播 达标结算主播数
	NotStandardHostCount   string // 工会所有主播 未达标结算主播数
	GuildCutRate           string // 工会抽成
	HostTotalIncomeBefore  string // 工会所有主播 未扣除主播总收入 印尼盾 分
	IsPayToGuild           string // 工资是否发放给公会 0非 1是
	CreateTime             string // 创建时间
	UpdateTime             string // 更新时间
}

// guildWeekSettlementColumns holds the columns for table guild_week_settlement.
var guildWeekSettlementColumns = GuildWeekSettlementColumns{
	Id:                     "id",
	SettlementWeek:         "settlement_week",
	GuildId:                "guild_id",
	UserId:                 "user_id",
	LiveTime:               "live_time",
	PayCommission:          "pay_commission",
	TotalIncome:            "total_income",
	EffectLiveDay:          "effect_live_day",
	GiftIncome:             "gift_income",
	GuildGiftIncome:        "guild_gift_income",
	WeekRankRewardForGuild: "week_rank_reward_for_guild",
	HostTotalIncomeCut:     "host_total_income_cut",
	HostTotalIncome:        "host_total_income",
	ActivityRewardForGuild: "activity_reward_for_guild",
	LiveRewardForGuild:     "live_reward_for_guild",
	GuildTotalIncome:       "guild_total_income",
	SettlementBankName:     "settlement_bank_name",
	SettlementBankAccounts: "settlement_bank_accounts",
	SettlementBankPerson:   "settlement_bank_person",
	IsSettle:               "is_settle",
	UpStandardHostCount:    "up_standard_host_count",
	NotStandardHostCount:   "not_standard_host_count",
	GuildCutRate:           "guild_cut_rate",
	HostTotalIncomeBefore:  "host_total_income_before",
	IsPayToGuild:           "is_pay_to_guild",
	CreateTime:             "create_time",
	UpdateTime:             "update_time",
}

// NewGuildWeekSettlementDao creates and returns a new DAO object for table data access.
func NewGuildWeekSettlementDao() *GuildWeekSettlementDao {
	return &GuildWeekSettlementDao{
		group:   "default",
		table:   "guild_week_settlement",
		columns: guildWeekSettlementColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GuildWeekSettlementDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GuildWeekSettlementDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GuildWeekSettlementDao) Columns() GuildWeekSettlementColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GuildWeekSettlementDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GuildWeekSettlementDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GuildWeekSettlementDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
