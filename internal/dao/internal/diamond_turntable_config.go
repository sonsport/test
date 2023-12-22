// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DiamondTurntableConfigDao is the data access object for table diamond_turntable_config.
type DiamondTurntableConfigDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns DiamondTurntableConfigColumns // columns contains all the column names of Table for convenient usage.
}

// DiamondTurntableConfigColumns defines and stores column names for table diamond_turntable_config.
type DiamondTurntableConfigColumns struct {
	Id         string //
	Gid        string // 商品表id
	Gtype      string // 商品类型 1装扮 2钻石
	Num        string // 数量
	Rate       string // 概率百分比  *100的整数
	Status     string // 商品状态 0失效 1生效
	CreateTime string //
	UpdateTime string //
}

// diamondTurntableConfigColumns holds the columns for table diamond_turntable_config.
var diamondTurntableConfigColumns = DiamondTurntableConfigColumns{
	Id:         "id",
	Gid:        "gid",
	Gtype:      "gtype",
	Num:        "num",
	Rate:       "rate",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewDiamondTurntableConfigDao creates and returns a new DAO object for table data access.
func NewDiamondTurntableConfigDao() *DiamondTurntableConfigDao {
	return &DiamondTurntableConfigDao{
		group:   "default",
		table:   "diamond_turntable_config",
		columns: diamondTurntableConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DiamondTurntableConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DiamondTurntableConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DiamondTurntableConfigDao) Columns() DiamondTurntableConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DiamondTurntableConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DiamondTurntableConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DiamondTurntableConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
