// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InviterInfoDao is the data access object for table inviter_info.
type InviterInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns InviterInfoColumns // columns contains all the column names of Table for convenient usage.
}

// InviterInfoColumns defines and stores column names for table inviter_info.
type InviterInfoColumns struct {
	Id              string //
	UserId          string //
	Stat            string //
	InviteUserCount string //
	CreateTime      string //
	UpdateTime      string //
}

// inviterInfoColumns holds the columns for table inviter_info.
var inviterInfoColumns = InviterInfoColumns{
	Id:              "id",
	UserId:          "user_id",
	Stat:            "stat",
	InviteUserCount: "invite_user_count",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// NewInviterInfoDao creates and returns a new DAO object for table data access.
func NewInviterInfoDao() *InviterInfoDao {
	return &InviterInfoDao{
		group:   "default",
		table:   "inviter_info",
		columns: inviterInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InviterInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InviterInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InviterInfoDao) Columns() InviterInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InviterInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InviterInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InviterInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
