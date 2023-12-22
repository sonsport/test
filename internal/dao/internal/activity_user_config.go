// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityUserConfigDao is the data access object for table activity_user_config.
type ActivityUserConfigDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ActivityUserConfigColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityUserConfigColumns defines and stores column names for table activity_user_config.
type ActivityUserConfigColumns struct {
	Id         string //
	AcId       string // 活动Id activity_config 自增id
	ShowUser   string // 展示对象：ALL 全部、UNPAID未充值用户 、PAID已经充值用户 、HOST主播公会 后续扩展
	CreateTime string //
	UpdateTime string //
}

// activityUserConfigColumns holds the columns for table activity_user_config.
var activityUserConfigColumns = ActivityUserConfigColumns{
	Id:         "id",
	AcId:       "ac_id",
	ShowUser:   "show_user",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewActivityUserConfigDao creates and returns a new DAO object for table data access.
func NewActivityUserConfigDao() *ActivityUserConfigDao {
	return &ActivityUserConfigDao{
		group:   "default",
		table:   "activity_user_config",
		columns: activityUserConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityUserConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityUserConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityUserConfigDao) Columns() ActivityUserConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityUserConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityUserConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityUserConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
