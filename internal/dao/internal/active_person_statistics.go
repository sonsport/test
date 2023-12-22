// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivePersonStatisticsDao is the data access object for table active_person_statistics.
type ActivePersonStatisticsDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns ActivePersonStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// ActivePersonStatisticsColumns defines and stores column names for table active_person_statistics.
type ActivePersonStatisticsColumns struct {
	Id               string // 主键ID
	BeginTime        string // 开始时间
	EndTime          string // 结束时间
	AppChannel       string // 渠道
	AppVersion       string // app版本
	ActiveType       string // 类型1日活，2周活，3月活
	ActiveUsersCount string // 活跃用户
}

// activePersonStatisticsColumns holds the columns for table active_person_statistics.
var activePersonStatisticsColumns = ActivePersonStatisticsColumns{
	Id:               "id",
	BeginTime:        "begin_time",
	EndTime:          "end_time",
	AppChannel:       "app_channel",
	AppVersion:       "app_version",
	ActiveType:       "active_type",
	ActiveUsersCount: "active_users_count",
}

// NewActivePersonStatisticsDao creates and returns a new DAO object for table data access.
func NewActivePersonStatisticsDao() *ActivePersonStatisticsDao {
	return &ActivePersonStatisticsDao{
		group:   "default",
		table:   "active_person_statistics",
		columns: activePersonStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivePersonStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivePersonStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivePersonStatisticsDao) Columns() ActivePersonStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivePersonStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivePersonStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivePersonStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
