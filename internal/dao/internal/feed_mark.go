// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedMarkDao is the data access object for table feed_mark.
type FeedMarkDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns FeedMarkColumns // columns contains all the column names of Table for convenient usage.
}

// FeedMarkColumns defines and stores column names for table feed_mark.
type FeedMarkColumns struct {
	Id         string // 自增Id
	UserId     string // 用户Id
	FeedId     string // 动态Id
	Mark       string // 标记类型：1喜欢、2不感兴趣 后续扩展
	CreateTime string //
	UpdateTime string //
}

// feedMarkColumns holds the columns for table feed_mark.
var feedMarkColumns = FeedMarkColumns{
	Id:         "id",
	UserId:     "user_id",
	FeedId:     "feed_id",
	Mark:       "mark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewFeedMarkDao creates and returns a new DAO object for table data access.
func NewFeedMarkDao() *FeedMarkDao {
	return &FeedMarkDao{
		group:   "default",
		table:   "feed_mark",
		columns: feedMarkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeedMarkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeedMarkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeedMarkDao) Columns() FeedMarkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeedMarkDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeedMarkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeedMarkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
