// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorLabelApplyDao is the data access object for table anchor_label_apply.
type AnchorLabelApplyDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns AnchorLabelApplyColumns // columns contains all the column names of Table for convenient usage.
}

// AnchorLabelApplyColumns defines and stores column names for table anchor_label_apply.
type AnchorLabelApplyColumns struct {
	Id         string //
	UserId     string //
	Label      string // 标签
	States     string // 状态 1 待审核 2 已通过 3 已驳回
	CreateTime string //
	UpdateTime string //
}

// anchorLabelApplyColumns holds the columns for table anchor_label_apply.
var anchorLabelApplyColumns = AnchorLabelApplyColumns{
	Id:         "id",
	UserId:     "user_id",
	Label:      "label",
	States:     "states",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewAnchorLabelApplyDao creates and returns a new DAO object for table data access.
func NewAnchorLabelApplyDao() *AnchorLabelApplyDao {
	return &AnchorLabelApplyDao{
		group:   "default",
		table:   "anchor_label_apply",
		columns: anchorLabelApplyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnchorLabelApplyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AnchorLabelApplyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AnchorLabelApplyDao) Columns() AnchorLabelApplyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AnchorLabelApplyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnchorLabelApplyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AnchorLabelApplyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
