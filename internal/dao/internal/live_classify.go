// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiveClassifyDao is the data access object for table live_classify.
type LiveClassifyDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns LiveClassifyColumns // columns contains all the column names of Table for convenient usage.
}

// LiveClassifyColumns defines and stores column names for table live_classify.
type LiveClassifyColumns struct {
	Id                  string //
	ParentId            string // 上级id 0 为顶级
	ClassifyType        string // 1 推荐 2 最新 3 热门 4 pk 5 密码房 6 钻石房 7 标签主播 8 无标签主播 9 新秀 10 夜播
	ClassifyDefaultName string // 默认分类名
	ClassifyRemark      string // 分类备注
	ClassifyStatus      string // 1 正常
	ClassifyRank        string // 排序
	CreateTime          string //
	UpdateTime          string //
}

// liveClassifyColumns holds the columns for table live_classify.
var liveClassifyColumns = LiveClassifyColumns{
	Id:                  "id",
	ParentId:            "parent_id",
	ClassifyType:        "classify_type",
	ClassifyDefaultName: "classify_default_name",
	ClassifyRemark:      "classify_remark",
	ClassifyStatus:      "classify_status",
	ClassifyRank:        "classify_rank",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
}

// NewLiveClassifyDao creates and returns a new DAO object for table data access.
func NewLiveClassifyDao() *LiveClassifyDao {
	return &LiveClassifyDao{
		group:   "default",
		table:   "live_classify",
		columns: liveClassifyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LiveClassifyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LiveClassifyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LiveClassifyDao) Columns() LiveClassifyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LiveClassifyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LiveClassifyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LiveClassifyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
