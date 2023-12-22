// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TopicFollowDao is the data access object for table topic_follow.
type TopicFollowDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TopicFollowColumns // columns contains all the column names of Table for convenient usage.
}

// TopicFollowColumns defines and stores column names for table topic_follow.
type TopicFollowColumns struct {
	Id         string // 自增Id
	TopicId    string // 话题Id
	UserId     string // 用户Id
	CreateTime string //
	UpdateTime string //
}

// topicFollowColumns holds the columns for table topic_follow.
var topicFollowColumns = TopicFollowColumns{
	Id:         "id",
	TopicId:    "topic_id",
	UserId:     "user_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewTopicFollowDao creates and returns a new DAO object for table data access.
func NewTopicFollowDao() *TopicFollowDao {
	return &TopicFollowDao{
		group:   "default",
		table:   "topic_follow",
		columns: topicFollowColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TopicFollowDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TopicFollowDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TopicFollowDao) Columns() TopicFollowColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TopicFollowDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TopicFollowDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TopicFollowDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
