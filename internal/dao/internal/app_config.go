// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppConfigDao is the data access object for table app_config.
type AppConfigDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AppConfigColumns // columns contains all the column names of Table for convenient usage.
}

// AppConfigColumns defines and stores column names for table app_config.
type AppConfigColumns struct {
	Id         string // 自增Id
	System     string // 所属系统；0所有,1安卓 2ios
	Version    string // 版本
	Channel    string // 渠道
	Key        string // key
	Value      string // 值
	Remark     string // 备注
	CreateTime string //
	UpdateTime string //
}

// appConfigColumns holds the columns for table app_config.
var appConfigColumns = AppConfigColumns{
	Id:         "id",
	System:     "system",
	Version:    "version",
	Channel:    "channel",
	Key:        "key",
	Value:      "value",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewAppConfigDao creates and returns a new DAO object for table data access.
func NewAppConfigDao() *AppConfigDao {
	return &AppConfigDao{
		group:   "default",
		table:   "app_config",
		columns: appConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AppConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AppConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AppConfigDao) Columns() AppConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AppConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AppConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AppConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
