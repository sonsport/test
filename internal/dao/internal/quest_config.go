// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QuestConfigDao is the data access object for table quest_config.
type QuestConfigDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns QuestConfigColumns // columns contains all the column names of Table for convenient usage.
}

// QuestConfigColumns defines and stores column names for table quest_config.
type QuestConfigColumns struct {
	Id               string // id
	QuestNameJson    string // 任务名称
	QuestIcon        string // 任务图标
	QuestStartTime   string // 任务开始时间
	QuestEndTime     string // 任务结束时间
	QuestKey         string // 任务唯一标识 key唯一则表示套娃任务 区分等级
	QuestParentId    string // 任务上级id -1 为一级任务
	QuestRank        string // 任务等级 套娃任务等级递增
	QuestCycle       string // 任务周期 1 每日 2 每周
	QuestTheme       string // 任务主题 1 日常任务 2 主播任务 3 用户钻石任务 4 用户活跃宝箱
	QuestUserScope   string // 任务用户范围 1 所有用户 2 星级主播 3 公会长
	QuestType        string // 任务类型 common.QuestTypeSignIn
	QuestAchieved    string // 任务达成值 1 次数任务代表1次 计时为1分钟
	IsReward         string // 任务是否有奖励 1 有
	QuestIsMust      string // 任务是否必须 1 是
	QuestSort        string // 任务排序
	QuestStates      string // 任务状态
	CreateTime       string // createdAt
	UpdateTime       string // updatedAt
	AutomaticRelease string // 自动发放 1 是
}

// questConfigColumns holds the columns for table quest_config.
var questConfigColumns = QuestConfigColumns{
	Id:               "id",
	QuestNameJson:    "quest_name_json",
	QuestIcon:        "quest_icon",
	QuestStartTime:   "quest_start_time",
	QuestEndTime:     "quest_end_time",
	QuestKey:         "quest_key",
	QuestParentId:    "quest_parent_id",
	QuestRank:        "quest_rank",
	QuestCycle:       "quest_cycle",
	QuestTheme:       "quest_theme",
	QuestUserScope:   "quest_user_scope",
	QuestType:        "quest_type",
	QuestAchieved:    "quest_achieved",
	IsReward:         "is_reward",
	QuestIsMust:      "quest_is_must",
	QuestSort:        "quest_sort",
	QuestStates:      "quest_states",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
	AutomaticRelease: "automatic_release",
}

// NewQuestConfigDao creates and returns a new DAO object for table data access.
func NewQuestConfigDao() *QuestConfigDao {
	return &QuestConfigDao{
		group:   "default",
		table:   "quest_config",
		columns: questConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *QuestConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *QuestConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *QuestConfigDao) Columns() QuestConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *QuestConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *QuestConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *QuestConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
