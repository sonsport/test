// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoblePrivilegeDao is the data access object for table noble_privilege.
type NoblePrivilegeDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns NoblePrivilegeColumns // columns contains all the column names of Table for convenient usage.
}

// NoblePrivilegeColumns defines and stores column names for table noble_privilege.
type NoblePrivilegeColumns struct {
	PrivilegeId     string //
	PrivilegeName   string // 特权名称
	PrivilegeIcon   string // 特权图片
	PrivilegeStates string // 特权状态 1 正常
	IsPrivilege     string // 是否特权 1特权 2商品
	PrivilegeType   string // 特权类型
	PrivilegeDesc   string // 描述 多语言配置 {en:xxx}
	NobleType       string // 贵族类型 1 vip
	Sort            string // 排序
	CreateTime      string //
	UpdateTime      string //
}

// noblePrivilegeColumns holds the columns for table noble_privilege.
var noblePrivilegeColumns = NoblePrivilegeColumns{
	PrivilegeId:     "privilege_id",
	PrivilegeName:   "privilege_name",
	PrivilegeIcon:   "privilege_icon",
	PrivilegeStates: "privilege_states",
	IsPrivilege:     "is_privilege",
	PrivilegeType:   "privilege_type",
	PrivilegeDesc:   "privilege_desc",
	NobleType:       "noble_type",
	Sort:            "sort",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// NewNoblePrivilegeDao creates and returns a new DAO object for table data access.
func NewNoblePrivilegeDao() *NoblePrivilegeDao {
	return &NoblePrivilegeDao{
		group:   "default",
		table:   "noble_privilege",
		columns: noblePrivilegeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoblePrivilegeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoblePrivilegeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoblePrivilegeDao) Columns() NoblePrivilegeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoblePrivilegeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoblePrivilegeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoblePrivilegeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
