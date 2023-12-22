// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SafeInfoDao is the data access object for table safe_info.
type SafeInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SafeInfoColumns // columns contains all the column names of Table for convenient usage.
}

// SafeInfoColumns defines and stores column names for table safe_info.
type SafeInfoColumns struct {
	Id         string //
	UserId     string // 用户Id
	Content    string // 内容
	SafeType   string // 类型，1头像，2昵称，3签名，后续扩展
	State      string // 状态，0待审核，1审核通过，2审核不通过
	CreateTime string //
	UpdateTime string //
}

// safeInfoColumns holds the columns for table safe_info.
var safeInfoColumns = SafeInfoColumns{
	Id:         "id",
	UserId:     "user_id",
	Content:    "content",
	SafeType:   "safe_type",
	State:      "state",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSafeInfoDao creates and returns a new DAO object for table data access.
func NewSafeInfoDao() *SafeInfoDao {
	return &SafeInfoDao{
		group:   "default",
		table:   "safe_info",
		columns: safeInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SafeInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SafeInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SafeInfoDao) Columns() SafeInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SafeInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SafeInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SafeInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
