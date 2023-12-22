// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NobleConfigDao is the data access object for table noble_config.
type NobleConfigDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns NobleConfigColumns // columns contains all the column names of Table for convenient usage.
}

// NobleConfigColumns defines and stores column names for table noble_config.
type NobleConfigColumns struct {
	Id            string //
	NobleId       string // 贵族id
	PrivilegeId   string // 特权id
	PrivilegeLink string // 特权值 类型为加成类型时val为倍数、商品时为商品id
	CreateTime    string //
	UpdateTime    string //
}

// nobleConfigColumns holds the columns for table noble_config.
var nobleConfigColumns = NobleConfigColumns{
	Id:            "id",
	NobleId:       "noble_id",
	PrivilegeId:   "privilege_id",
	PrivilegeLink: "privilege_link",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewNobleConfigDao creates and returns a new DAO object for table data access.
func NewNobleConfigDao() *NobleConfigDao {
	return &NobleConfigDao{
		group:   "default",
		table:   "noble_config",
		columns: nobleConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NobleConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NobleConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NobleConfigDao) Columns() NobleConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NobleConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NobleConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NobleConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
