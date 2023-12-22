// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedCommentMarkDao is the data access object for table feed_comment_mark.
type FeedCommentMarkDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns FeedCommentMarkColumns // columns contains all the column names of Table for convenient usage.
}

// FeedCommentMarkColumns defines and stores column names for table feed_comment_mark.
type FeedCommentMarkColumns struct {
	Id            string // 自增Id
	UserId        string // 用户Id
	FeedCommentId string // 评论Id
	Mark          string // 标记类型：1喜欢、 后续扩展
	CreateTime    string //
	UpdateTime    string //
}

// feedCommentMarkColumns holds the columns for table feed_comment_mark.
var feedCommentMarkColumns = FeedCommentMarkColumns{
	Id:            "id",
	UserId:        "user_id",
	FeedCommentId: "feed_comment_id",
	Mark:          "mark",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewFeedCommentMarkDao creates and returns a new DAO object for table data access.
func NewFeedCommentMarkDao() *FeedCommentMarkDao {
	return &FeedCommentMarkDao{
		group:   "default",
		table:   "feed_comment_mark",
		columns: feedCommentMarkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeedCommentMarkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeedCommentMarkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeedCommentMarkDao) Columns() FeedCommentMarkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeedCommentMarkDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeedCommentMarkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeedCommentMarkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
