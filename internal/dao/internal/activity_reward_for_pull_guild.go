// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityRewardForPullGuildDao is the data access object for table activity_reward_for_pull_guild.
type ActivityRewardForPullGuildDao struct {
	table   string                            // table is the underlying table name of the DAO.
	group   string                            // group is the database configuration group name of current DAO.
	columns ActivityRewardForPullGuildColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityRewardForPullGuildColumns defines and stores column names for table activity_reward_for_pull_guild.
type ActivityRewardForPullGuildColumns struct {
	Id                string //
	UserId            string // 用户Id 主播id
	GuildId           string // 公会Id 关联guild_info自增
	GuildUserId       string // 用户所属公会长Id （这里是B公会长ID）
	PullerGuildUserId string // A公会拉的B公会 这里存储A公会长的ID
	LiveTime          string // 直播时长 单位秒
	IsReward          string // 是否已经奖励 0未发放 1已发放
	Reward            string // 奖励、存储印尼盾分 奖励给B公会长的
	RewardDate        string // 发放日期
	CreatedTime       string //
	UpdatedTime       string //
}

// activityRewardForPullGuildColumns holds the columns for table activity_reward_for_pull_guild.
var activityRewardForPullGuildColumns = ActivityRewardForPullGuildColumns{
	Id:                "id",
	UserId:            "user_id",
	GuildId:           "guild_id",
	GuildUserId:       "guild_user_id",
	PullerGuildUserId: "puller_guild_user_id",
	LiveTime:          "live_time",
	IsReward:          "is_reward",
	Reward:            "reward",
	RewardDate:        "reward_date",
	CreatedTime:       "created_time",
	UpdatedTime:       "updated_time",
}

// NewActivityRewardForPullGuildDao creates and returns a new DAO object for table data access.
func NewActivityRewardForPullGuildDao() *ActivityRewardForPullGuildDao {
	return &ActivityRewardForPullGuildDao{
		group:   "default",
		table:   "activity_reward_for_pull_guild",
		columns: activityRewardForPullGuildColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityRewardForPullGuildDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityRewardForPullGuildDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityRewardForPullGuildDao) Columns() ActivityRewardForPullGuildColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityRewardForPullGuildDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityRewardForPullGuildDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityRewardForPullGuildDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
