// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DiamondTurntableRecordDao is the data access object for table diamond_turntable_record.
type DiamondTurntableRecordDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns DiamondTurntableRecordColumns // columns contains all the column names of Table for convenient usage.
}

// DiamondTurntableRecordColumns defines and stores column names for table diamond_turntable_record.
type DiamondTurntableRecordColumns struct {
	Id             string //
	UserId         string // 用户id
	Cid            string // 配置表id
	Gid            string // 商品表id
	Num            string // 数量
	TurntableCount string // 奖池期数
	CreateTime     string //
	UpdateTime     string //
}

// diamondTurntableRecordColumns holds the columns for table diamond_turntable_record.
var diamondTurntableRecordColumns = DiamondTurntableRecordColumns{
	Id:             "id",
	UserId:         "user_id",
	Cid:            "cid",
	Gid:            "gid",
	Num:            "num",
	TurntableCount: "turntable_count",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewDiamondTurntableRecordDao creates and returns a new DAO object for table data access.
func NewDiamondTurntableRecordDao() *DiamondTurntableRecordDao {
	return &DiamondTurntableRecordDao{
		group:   "default",
		table:   "diamond_turntable_record",
		columns: diamondTurntableRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DiamondTurntableRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DiamondTurntableRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DiamondTurntableRecordDao) Columns() DiamondTurntableRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DiamondTurntableRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DiamondTurntableRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DiamondTurntableRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
