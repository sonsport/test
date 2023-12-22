// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QuestUserDetailDao is the data access object for table quest_user_detail.
type QuestUserDetailDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns QuestUserDetailColumns // columns contains all the column names of Table for convenient usage.
}

// QuestUserDetailColumns defines and stores column names for table quest_user_detail.
type QuestUserDetailColumns struct {
	Id             string // id
	QuestId        string // 任务id
	QuestUserId    string // 参与用户id
	ExpirationTime string // 过期时间
	QuestAchieved  string // 任务达成值 1 次数任务代表1次 计时为1分钟
	IsComplete     string // 是否完成
	CompleteTime   string // 完成时间
	QuestKey       string // 任务唯一标识 key唯一则表示套娃任务 区分等级
	QuestRank      string // 任务等级 套娃任务等级递增
	CreateTime     string // createdAt
	UpdateTime     string // updatedAt
}

// questUserDetailColumns holds the columns for table quest_user_detail.
var questUserDetailColumns = QuestUserDetailColumns{
	Id:             "id",
	QuestId:        "quest_id",
	QuestUserId:    "quest_user_id",
	ExpirationTime: "expiration_time",
	QuestAchieved:  "quest_achieved",
	IsComplete:     "is_complete",
	CompleteTime:   "complete_time",
	QuestKey:       "quest_key",
	QuestRank:      "quest_rank",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewQuestUserDetailDao creates and returns a new DAO object for table data access.
func NewQuestUserDetailDao() *QuestUserDetailDao {
	return &QuestUserDetailDao{
		group:   "default",
		table:   "quest_user_detail",
		columns: questUserDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *QuestUserDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *QuestUserDetailDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *QuestUserDetailDao) Columns() QuestUserDetailColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *QuestUserDetailDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *QuestUserDetailDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *QuestUserDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
