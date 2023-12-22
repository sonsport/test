// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityChannelConfigDao is the data access object for table activity_channel_config.
type ActivityChannelConfigDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns ActivityChannelConfigColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityChannelConfigColumns defines and stores column names for table activity_channel_config.
type ActivityChannelConfigColumns struct {
	Id          string //
	AcId        string // 活动Id activity_config 自增id
	ShowChannel string // 渠道：ALL全部渠道、指定渠道
	CreateTime  string //
	UpdateTime  string //
}

// activityChannelConfigColumns holds the columns for table activity_channel_config.
var activityChannelConfigColumns = ActivityChannelConfigColumns{
	Id:          "id",
	AcId:        "ac_id",
	ShowChannel: "show_channel",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewActivityChannelConfigDao creates and returns a new DAO object for table data access.
func NewActivityChannelConfigDao() *ActivityChannelConfigDao {
	return &ActivityChannelConfigDao{
		group:   "default",
		table:   "activity_channel_config",
		columns: activityChannelConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityChannelConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityChannelConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityChannelConfigDao) Columns() ActivityChannelConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityChannelConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityChannelConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityChannelConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
