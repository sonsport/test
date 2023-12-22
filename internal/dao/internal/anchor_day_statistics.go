// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorDayStatisticsDao is the data access object for table anchor_day_statistics.
type AnchorDayStatisticsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns AnchorDayStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// AnchorDayStatisticsColumns defines and stores column names for table anchor_day_statistics.
type AnchorDayStatisticsColumns struct {
	StatisticsId            string // 主键id
	StatisticsTime          string // 统计时间
	UserId                  string // 用户id
	GuildId                 string // 工会id
	AllIncome               string // 总收益
	NormalGiftIncome        string // 普通礼物收益
	LuckyGiftIncome         string // 幸运礼物收益
	LiveTime                string // 直播时长（s）
	EffectLiveTime          string // 有效直播时长（s） 大于30分钟的直播才算有效直播累计
	IsEffectDay             string // 是否有效天 0非 1是
	PkCount                 string // 总pk次数
	EffectivePkCount        string // 有效pk次数
	WatchMemberCount        string // 观看人数
	NewFansCount            string // 新增粉丝数
	SendGiftPerson          string // 新增送礼人数
	IsFirstLive             string // 是否首播
	Role                    string // 角色
	RewardDayForGuild       string // 当天对公会的奖励（印尼盾 单位分），比如主播当天首播满1小时给公会的奖励
	RewardDayForGuildRemark string // 活动奖励说明
	SumHotAward             string // 热门钻石奖励
	HotEffectLive           string // 热门开播
	CreateTime              string // 创建时间
	DayWaterDiamond         string // 直播天流水
}

// anchorDayStatisticsColumns holds the columns for table anchor_day_statistics.
var anchorDayStatisticsColumns = AnchorDayStatisticsColumns{
	StatisticsId:            "statistics_id",
	StatisticsTime:          "statistics_time",
	UserId:                  "user_id",
	GuildId:                 "guild_id",
	AllIncome:               "all_income",
	NormalGiftIncome:        "normal_gift_income",
	LuckyGiftIncome:         "lucky_gift_income",
	LiveTime:                "live_time",
	EffectLiveTime:          "effect_live_time",
	IsEffectDay:             "is_effect_day",
	PkCount:                 "pk_count",
	EffectivePkCount:        "effective_pk_count",
	WatchMemberCount:        "watch_member_count",
	NewFansCount:            "new_fans_count",
	SendGiftPerson:          "send_gift_person",
	IsFirstLive:             "is_first_live",
	Role:                    "role",
	RewardDayForGuild:       "reward_day_for_guild",
	RewardDayForGuildRemark: "reward_day_for_guild_remark",
	SumHotAward:             "sum_hot_award",
	HotEffectLive:           "hot_effect_live",
	CreateTime:              "create_time",
	DayWaterDiamond:         "day_water_diamond",
}

// NewAnchorDayStatisticsDao creates and returns a new DAO object for table data access.
func NewAnchorDayStatisticsDao() *AnchorDayStatisticsDao {
	return &AnchorDayStatisticsDao{
		group:   "default",
		table:   "anchor_day_statistics",
		columns: anchorDayStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnchorDayStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AnchorDayStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AnchorDayStatisticsDao) Columns() AnchorDayStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AnchorDayStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnchorDayStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AnchorDayStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
