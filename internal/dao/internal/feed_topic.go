// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedTopicDao is the data access object for table feed_topic.
type FeedTopicDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns FeedTopicColumns // columns contains all the column names of Table for convenient usage.
}

// FeedTopicColumns defines and stores column names for table feed_topic.
type FeedTopicColumns struct {
	Id         string // 自增Id
	TopicId    string // 话题Id
	FeedId     string // 动态Id
	CreateTime string //
	UpdateTime string //
}

// feedTopicColumns holds the columns for table feed_topic.
var feedTopicColumns = FeedTopicColumns{
	Id:         "id",
	TopicId:    "topic_id",
	FeedId:     "feed_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewFeedTopicDao creates and returns a new DAO object for table data access.
func NewFeedTopicDao() *FeedTopicDao {
	return &FeedTopicDao{
		group:   "default",
		table:   "feed_topic",
		columns: feedTopicColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeedTopicDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeedTopicDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeedTopicDao) Columns() FeedTopicColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeedTopicDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeedTopicDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeedTopicDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
