// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityConfigDao is the data access object for table activity_config.
type ActivityConfigDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns ActivityConfigColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityConfigColumns defines and stores column names for table activity_config.
type ActivityConfigColumns struct {
	Id          string //
	Title       string // 活动标题
	ResUrl      string // 资源图 活动横图
	LinkAddress string // 活动地址
	StartTime   string // 活动开始时间
	EndTime     string // 活动结束时间
	State       string // 状态 0下架 1上架
	Sort        string // 排序，顺序排列
	Remark      string // 备注
	CreateTime  string //
	UpdateTime  string //
}

// activityConfigColumns holds the columns for table activity_config.
var activityConfigColumns = ActivityConfigColumns{
	Id:          "id",
	Title:       "title",
	ResUrl:      "res_url",
	LinkAddress: "link_address",
	StartTime:   "start_time",
	EndTime:     "end_time",
	State:       "state",
	Sort:        "sort",
	Remark:      "remark",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewActivityConfigDao creates and returns a new DAO object for table data access.
func NewActivityConfigDao() *ActivityConfigDao {
	return &ActivityConfigDao{
		group:   "default",
		table:   "activity_config",
		columns: activityConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityConfigDao) Columns() ActivityConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
