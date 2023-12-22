// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityPlanDao is the data access object for table activity_plan.
type ActivityPlanDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ActivityPlanColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityPlanColumns defines and stores column names for table activity_plan.
type ActivityPlanColumns struct {
	Id           string //
	ActivityType string // 活动类别,榜单活动
	Title        string // 活动标题
	Icon         string // 图片
	Explain      string // 活动说明
	StartTime    string // 活动开始时间
	EndTime      string // 活动结束时间
	TimeLevel    string // 活动数据类型,（自然周、自然天、时）也可以自定义时间类型（最小颗粒度：小时）、周期内（all）：0
	RewardCycle  string // 奖励发放周期配置：小时、天、周、默认0不奖励、统一指定时间奖励
	Object       string // 奖励对象：主播/公会长/用户 或者多个
	Source       string // 奖励数据源：rank_diamonds、rank_coin、rank_gift(根据礼物id判断)……,以此类推；根据key标记调用对应的封装接口 不配置默认为不调用数据源（特殊活动/非模板活动）
	AppChannel   string // 能参加活动的渠道
	Sort         string // 排序
	State        string // 活动状态是否有效，1有效，2无效
	CreateTime   string //
	UpdateTime   string //
}

// activityPlanColumns holds the columns for table activity_plan.
var activityPlanColumns = ActivityPlanColumns{
	Id:           "id",
	ActivityType: "activity_type",
	Title:        "title",
	Icon:         "icon",
	Explain:      "explain",
	StartTime:    "start_time",
	EndTime:      "end_time",
	TimeLevel:    "time_level",
	RewardCycle:  "reward_cycle",
	Object:       "object",
	Source:       "source",
	AppChannel:   "app_channel",
	Sort:         "sort",
	State:        "state",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewActivityPlanDao creates and returns a new DAO object for table data access.
func NewActivityPlanDao() *ActivityPlanDao {
	return &ActivityPlanDao{
		group:   "default",
		table:   "activity_plan",
		columns: activityPlanColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityPlanDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityPlanDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityPlanDao) Columns() ActivityPlanColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityPlanDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityPlanDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityPlanDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
