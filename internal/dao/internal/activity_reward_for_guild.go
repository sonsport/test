// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityRewardForGuildDao is the data access object for table activity_reward_for_guild.
type ActivityRewardForGuildDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns ActivityRewardForGuildColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityRewardForGuildColumns defines and stores column names for table activity_reward_for_guild.
type ActivityRewardForGuildColumns struct {
	Id             string //
	UserId         string // 用户Id 主播id
	GuildId        string // 公会Id 关联guild_info自增
	ActivityType   string // 活动类型  0默认 1首播活动 2四周活动 后续扩展
	EffectLiveTime string // 有效直播时长（s） 首播、四周活动存储热门直播时长
	EffectDay      string // 有效直播天
	IsReward       string // 是否已经奖励 0未发放 1已发放
	Reward         string // 奖励、存储印尼盾分，后续其他使用可以存储别的 根据不同的活动类型
	RewardDate     string // 发放日期
	CreatedTime    string //
	UpdatedTime    string //
}

// activityRewardForGuildColumns holds the columns for table activity_reward_for_guild.
var activityRewardForGuildColumns = ActivityRewardForGuildColumns{
	Id:             "id",
	UserId:         "user_id",
	GuildId:        "guild_id",
	ActivityType:   "activity_type",
	EffectLiveTime: "effect_live_time",
	EffectDay:      "effect_day",
	IsReward:       "is_reward",
	Reward:         "reward",
	RewardDate:     "reward_date",
	CreatedTime:    "created_time",
	UpdatedTime:    "updated_time",
}

// NewActivityRewardForGuildDao creates and returns a new DAO object for table data access.
func NewActivityRewardForGuildDao() *ActivityRewardForGuildDao {
	return &ActivityRewardForGuildDao{
		group:   "default",
		table:   "activity_reward_for_guild",
		columns: activityRewardForGuildColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityRewardForGuildDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityRewardForGuildDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityRewardForGuildDao) Columns() ActivityRewardForGuildColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityRewardForGuildDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityRewardForGuildDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityRewardForGuildDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
