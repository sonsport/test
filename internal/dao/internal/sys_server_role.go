// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysServerRoleDao is the data access object for table sys_server_role.
type SysServerRoleDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysServerRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysServerRoleColumns defines and stores column names for table sys_server_role.
type SysServerRoleColumns struct {
	Id         string // ID
	AdminId    string // 关联admin主键
	Role       string // 角色0普通，1管理
	Group      string // 管理组，1雅加达，2三宝垄
	CreateTime string // 创建日期
	UpdateTime string // 更新时间
}

// sysServerRoleColumns holds the columns for table sys_server_role.
var sysServerRoleColumns = SysServerRoleColumns{
	Id:         "id",
	AdminId:    "admin_id",
	Role:       "role",
	Group:      "group",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewSysServerRoleDao creates and returns a new DAO object for table data access.
func NewSysServerRoleDao() *SysServerRoleDao {
	return &SysServerRoleDao{
		group:   "default",
		table:   "sys_server_role",
		columns: sysServerRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysServerRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysServerRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysServerRoleDao) Columns() SysServerRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysServerRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysServerRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysServerRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
