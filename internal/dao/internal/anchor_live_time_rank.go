// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorLiveTimeRankDao is the data access object for table anchor_live_time_rank.
type AnchorLiveTimeRankDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns AnchorLiveTimeRankColumns // columns contains all the column names of Table for convenient usage.
}

// AnchorLiveTimeRankColumns defines and stores column names for table anchor_live_time_rank.
type AnchorLiveTimeRankColumns struct {
	Id                string //
	UserId            string // 用户ID
	RankDateTime      string // 排行日期
	EffectLiveSeconds string // 有效直播时长
	CreateTime        string //
	UpdateTime        string //
}

// anchorLiveTimeRankColumns holds the columns for table anchor_live_time_rank.
var anchorLiveTimeRankColumns = AnchorLiveTimeRankColumns{
	Id:                "id",
	UserId:            "user_id",
	RankDateTime:      "rank_date_time",
	EffectLiveSeconds: "effect_live_seconds",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
}

// NewAnchorLiveTimeRankDao creates and returns a new DAO object for table data access.
func NewAnchorLiveTimeRankDao() *AnchorLiveTimeRankDao {
	return &AnchorLiveTimeRankDao{
		group:   "default",
		table:   "anchor_live_time_rank",
		columns: anchorLiveTimeRankColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnchorLiveTimeRankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AnchorLiveTimeRankDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AnchorLiveTimeRankDao) Columns() AnchorLiveTimeRankColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AnchorLiveTimeRankDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnchorLiveTimeRankDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AnchorLiveTimeRankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
