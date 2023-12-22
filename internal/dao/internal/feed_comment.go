// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedCommentDao is the data access object for table feed_comment.
type FeedCommentDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns FeedCommentColumns // columns contains all the column names of Table for convenient usage.
}

// FeedCommentColumns defines and stores column names for table feed_comment.
type FeedCommentColumns struct {
	Id             string // 自增Id
	UserId         string // 评论用户Id
	FeedId         string // 动态Id
	ParentId       string // 父级评论ID，  0标识对一级动态回复、 >0标识对评论回复只存储一级回复的Id（自增id）
	ToUserId       string // A回复了B，这里存储B的用户ID ，一级回复id这里存储0
	CommentId      string // 该条评论针对那个评论的回复，一级回复存储0，针对评论的的回复存储回复评论的ID，该数据只是存储
	ReplyContext   string // 评论/回复内容
	ReplyLikeCount string // 评论点赞数
	State          string // 0未审核、1审核通过、2审核不通过、3系统删除、4用户自己删除
	CreateTime     string //
	UpdateTime     string //
}

// feedCommentColumns holds the columns for table feed_comment.
var feedCommentColumns = FeedCommentColumns{
	Id:             "id",
	UserId:         "user_id",
	FeedId:         "feed_id",
	ParentId:       "parent_id",
	ToUserId:       "to_user_id",
	CommentId:      "comment_id",
	ReplyContext:   "reply_context",
	ReplyLikeCount: "reply_like_count",
	State:          "state",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewFeedCommentDao creates and returns a new DAO object for table data access.
func NewFeedCommentDao() *FeedCommentDao {
	return &FeedCommentDao{
		group:   "default",
		table:   "feed_comment",
		columns: feedCommentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeedCommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeedCommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeedCommentDao) Columns() FeedCommentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeedCommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeedCommentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeedCommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
