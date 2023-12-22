// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MedalAppConfigDao is the data access object for table medal_app_config.
type MedalAppConfigDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns MedalAppConfigColumns // columns contains all the column names of Table for convenient usage.
}

// MedalAppConfigColumns defines and stores column names for table medal_app_config.
type MedalAppConfigColumns struct {
	Id              string // 主键id
	MedalId         string // 勋章id
	AppName         string // app名称
	MedalNameConfig string // 勋章名称配置json
	MedalDescConfig string // 勋章描述配置json
}

// medalAppConfigColumns holds the columns for table medal_app_config.
var medalAppConfigColumns = MedalAppConfigColumns{
	Id:              "id",
	MedalId:         "medal_id",
	AppName:         "app_name",
	MedalNameConfig: "medal_name_config",
	MedalDescConfig: "medal_desc_config",
}

// NewMedalAppConfigDao creates and returns a new DAO object for table data access.
func NewMedalAppConfigDao() *MedalAppConfigDao {
	return &MedalAppConfigDao{
		group:   "default",
		table:   "medal_app_config",
		columns: medalAppConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MedalAppConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MedalAppConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MedalAppConfigDao) Columns() MedalAppConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MedalAppConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MedalAppConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MedalAppConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
