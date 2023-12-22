// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PrizeCashbackConfigDao is the data access object for table prize_cashback_config.
type PrizeCashbackConfigDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns PrizeCashbackConfigColumns // columns contains all the column names of Table for convenient usage.
}

// PrizeCashbackConfigColumns defines and stores column names for table prize_cashback_config.
type PrizeCashbackConfigColumns struct {
	Id                     string // 主键id
	Name                   string // 名称
	CashbackRatio          string // 返现比例
	FiveTimesNumber        string // 5倍数量
	TenTimesNumber         string // 10倍数量
	FiftyTimesNumber       string // 50倍数量
	OneHundredTimesNumber  string // 100倍数量
	FiveHundredTimesNumber string // 500倍
	CreateTime             string // 创建时间
	UpdateTime             string // 更新时间
}

// prizeCashbackConfigColumns holds the columns for table prize_cashback_config.
var prizeCashbackConfigColumns = PrizeCashbackConfigColumns{
	Id:                     "id",
	Name:                   "name",
	CashbackRatio:          "cashback_ratio",
	FiveTimesNumber:        "five_times_number",
	TenTimesNumber:         "ten_times_number",
	FiftyTimesNumber:       "fifty_times_number",
	OneHundredTimesNumber:  "one_hundred_times_number",
	FiveHundredTimesNumber: "five_hundred_times_number",
	CreateTime:             "create_time",
	UpdateTime:             "update_time",
}

// NewPrizeCashbackConfigDao creates and returns a new DAO object for table data access.
func NewPrizeCashbackConfigDao() *PrizeCashbackConfigDao {
	return &PrizeCashbackConfigDao{
		group:   "default",
		table:   "prize_cashback_config",
		columns: prizeCashbackConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PrizeCashbackConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PrizeCashbackConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PrizeCashbackConfigDao) Columns() PrizeCashbackConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PrizeCashbackConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PrizeCashbackConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PrizeCashbackConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
