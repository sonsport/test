// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TopicInfoDao is the data access object for table topic_info.
type TopicInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TopicInfoColumns // columns contains all the column names of Table for convenient usage.
}

// TopicInfoColumns defines and stores column names for table topic_info.
type TopicInfoColumns struct {
	Id             string // 自增Id
	UserId         string // 用户Id
	Title          string // 话题标题
	CoverImg       string // 封面图
	Score          string // 值越大 热度越高
	TopicType      string // 0普通话题 1官方 2审核模式话题
	RecommendScore string // 推荐 默认是0  排序倒序 值越大越靠前
	State          string // 状态：0待审核 1审核通过 2审核不通过
	CreateTime     string //
	UpdateTime     string //
}

// topicInfoColumns holds the columns for table topic_info.
var topicInfoColumns = TopicInfoColumns{
	Id:             "id",
	UserId:         "user_id",
	Title:          "title",
	CoverImg:       "cover_img",
	Score:          "score",
	TopicType:      "topic_type",
	RecommendScore: "recommend_score",
	State:          "state",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewTopicInfoDao creates and returns a new DAO object for table data access.
func NewTopicInfoDao() *TopicInfoDao {
	return &TopicInfoDao{
		group:   "default",
		table:   "topic_info",
		columns: topicInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TopicInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TopicInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TopicInfoDao) Columns() TopicInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TopicInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TopicInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TopicInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
