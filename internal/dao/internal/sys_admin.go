// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAdminDao is the data access object for table sys_admin.
type SysAdminDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SysAdminColumns // columns contains all the column names of Table for convenient usage.
}

// SysAdminColumns defines and stores column names for table sys_admin.
type SysAdminColumns struct {
	Id         string //
	UserName   string // 用户名
	NickName   string // 昵称
	Password   string // 密码
	CreateTime string //
	UpdateTime string //
	Status     string //
}

// sysAdminColumns holds the columns for table sys_admin.
var sysAdminColumns = SysAdminColumns{
	Id:         "id",
	UserName:   "user_name",
	NickName:   "nick_name",
	Password:   "password",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Status:     "status",
}

// NewSysAdminDao creates and returns a new DAO object for table data access.
func NewSysAdminDao() *SysAdminDao {
	return &SysAdminDao{
		group:   "default",
		table:   "sys_admin",
		columns: sysAdminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAdminDao) Columns() SysAdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
