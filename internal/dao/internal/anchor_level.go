// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorLevelDao is the data access object for table anchor_level.
type AnchorLevelDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns AnchorLevelColumns // columns contains all the column names of Table for convenient usage.
}

// AnchorLevelColumns defines and stores column names for table anchor_level.
type AnchorLevelColumns struct {
	Id         string // 自增Id
	UserId     string // 用户ID
	Level      string // 等级
	Remark     string // 等级
	State      string // 状态1有效，2无效
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// anchorLevelColumns holds the columns for table anchor_level.
var anchorLevelColumns = AnchorLevelColumns{
	Id:         "id",
	UserId:     "user_id",
	Level:      "level",
	Remark:     "remark",
	State:      "state",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewAnchorLevelDao creates and returns a new DAO object for table data access.
func NewAnchorLevelDao() *AnchorLevelDao {
	return &AnchorLevelDao{
		group:   "default",
		table:   "anchor_level",
		columns: anchorLevelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnchorLevelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AnchorLevelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AnchorLevelDao) Columns() AnchorLevelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AnchorLevelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnchorLevelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AnchorLevelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
