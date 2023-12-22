// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserCheckFirstDao is the data access object for table user_check_first.
type UserCheckFirstDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns UserCheckFirstColumns // columns contains all the column names of Table for convenient usage.
}

// UserCheckFirstColumns defines and stores column names for table user_check_first.
type UserCheckFirstColumns struct {
	Id         string //
	UserId     string // 用户Id
	CheckType  string // 检查类型
	CreateTime string // 创建时间
	UpdateTime string // 最后更新时间
}

// userCheckFirstColumns holds the columns for table user_check_first.
var userCheckFirstColumns = UserCheckFirstColumns{
	Id:         "id",
	UserId:     "user_id",
	CheckType:  "check_type",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewUserCheckFirstDao creates and returns a new DAO object for table data access.
func NewUserCheckFirstDao() *UserCheckFirstDao {
	return &UserCheckFirstDao{
		group:   "default",
		table:   "user_check_first",
		columns: userCheckFirstColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserCheckFirstDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserCheckFirstDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserCheckFirstDao) Columns() UserCheckFirstColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserCheckFirstDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserCheckFirstDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserCheckFirstDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
