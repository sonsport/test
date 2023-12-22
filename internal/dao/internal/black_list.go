// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlackListDao is the data access object for table black_list.
type BlackListDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns BlackListColumns // columns contains all the column names of Table for convenient usage.
}

// BlackListColumns defines and stores column names for table black_list.
type BlackListColumns struct {
	Id         string // ID
	UserId     string // 用户ID
	OperatorId string // 操作者ID
	BlockedId  string // 被拉黑的用户ID
	EndTime    string // 拉黑结束时间,为0就是永久拉黑
	CreateTime string //
	UpdateTime string //
}

// blackListColumns holds the columns for table black_list.
var blackListColumns = BlackListColumns{
	Id:         "id",
	UserId:     "user_id",
	OperatorId: "operator_id",
	BlockedId:  "blocked_Id",
	EndTime:    "end_time",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewBlackListDao creates and returns a new DAO object for table data access.
func NewBlackListDao() *BlackListDao {
	return &BlackListDao{
		group:   "default",
		table:   "black_list",
		columns: blackListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlackListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlackListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlackListDao) Columns() BlackListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlackListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlackListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlackListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
