// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorRankRuleDao is the data access object for table anchor_rank_rule.
type AnchorRankRuleDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns AnchorRankRuleColumns // columns contains all the column names of Table for convenient usage.
}

// AnchorRankRuleColumns defines and stores column names for table anchor_rank_rule.
type AnchorRankRuleColumns struct {
	Id         string // 主键ID
	Level      string // 等级从0开始, 0为等级1
	Min        string // 最小值
	Max        string // 最大值
	CreateTime string //
	UpdateTime string //
}

// anchorRankRuleColumns holds the columns for table anchor_rank_rule.
var anchorRankRuleColumns = AnchorRankRuleColumns{
	Id:         "id",
	Level:      "level",
	Min:        "min",
	Max:        "max",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewAnchorRankRuleDao creates and returns a new DAO object for table data access.
func NewAnchorRankRuleDao() *AnchorRankRuleDao {
	return &AnchorRankRuleDao{
		group:   "default",
		table:   "anchor_rank_rule",
		columns: anchorRankRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnchorRankRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AnchorRankRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AnchorRankRuleDao) Columns() AnchorRankRuleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AnchorRankRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnchorRankRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AnchorRankRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
