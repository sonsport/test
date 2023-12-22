// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MedalUserDao is the data access object for table medal_user.
type MedalUserDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns MedalUserColumns // columns contains all the column names of Table for convenient usage.
}

// MedalUserColumns defines and stores column names for table medal_user.
type MedalUserColumns struct {
	Id          string //
	UserId      string // 用户id
	MedalId     string // 勋章id
	AccessCount string // 获得次数
	CreateTime  string //
	UpdateTime  string //
}

// medalUserColumns holds the columns for table medal_user.
var medalUserColumns = MedalUserColumns{
	Id:          "id",
	UserId:      "user_id",
	MedalId:     "medal_id",
	AccessCount: "access_count",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewMedalUserDao creates and returns a new DAO object for table data access.
func NewMedalUserDao() *MedalUserDao {
	return &MedalUserDao{
		group:   "default",
		table:   "medal_user",
		columns: medalUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MedalUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MedalUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MedalUserDao) Columns() MedalUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MedalUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MedalUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MedalUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
