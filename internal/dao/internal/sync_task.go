// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SyncTaskDao is the data access object for table sync_task.
type SyncTaskDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SyncTaskColumns // columns contains all the column names of Table for convenient usage.
}

// SyncTaskColumns defines and stores column names for table sync_task.
type SyncTaskColumns struct {
	Id        string // 任务id
	SrcTable  string // 数据来源表
	DstTable  string // 数据去向表
	Remark    string // 备注
	SyncId    string // 已同步/统计的数据来源表的最大记录id
	UpdatedAt string //
}

// syncTaskColumns holds the columns for table sync_task.
var syncTaskColumns = SyncTaskColumns{
	Id:        "id",
	SrcTable:  "src_table",
	DstTable:  "dst_table",
	Remark:    "remark",
	SyncId:    "sync_id",
	UpdatedAt: "updated_at",
}

// NewSyncTaskDao creates and returns a new DAO object for table data access.
func NewSyncTaskDao() *SyncTaskDao {
	return &SyncTaskDao{
		group:   "default",
		table:   "sync_task",
		columns: syncTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SyncTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SyncTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SyncTaskDao) Columns() SyncTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SyncTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SyncTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SyncTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
