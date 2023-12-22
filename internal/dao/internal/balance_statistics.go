// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceStatisticsDao is the data access object for table balance_statistics.
type BalanceStatisticsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns BalanceStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// BalanceStatisticsColumns defines and stores column names for table balance_statistics.
type BalanceStatisticsColumns struct {
	Id            string //
	DepletionType string // 消费类型
	Diamonds      string // 钻石变动数额
	BeginTime     string // 开始时间
	EndTime       string // 结束时间
	TimeLevel     string // 时间维度：1.每小时，2.每天
	BigType       string // 消费类型 1支出 2收入
	CreateTime    string //
	UpdateTime    string //
}

// balanceStatisticsColumns holds the columns for table balance_statistics.
var balanceStatisticsColumns = BalanceStatisticsColumns{
	Id:            "id",
	DepletionType: "depletion_type",
	Diamonds:      "diamonds",
	BeginTime:     "begin_time",
	EndTime:       "end_time",
	TimeLevel:     "time_level",
	BigType:       "big_type",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewBalanceStatisticsDao creates and returns a new DAO object for table data access.
func NewBalanceStatisticsDao() *BalanceStatisticsDao {
	return &BalanceStatisticsDao{
		group:   "default",
		table:   "balance_statistics",
		columns: balanceStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BalanceStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BalanceStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BalanceStatisticsDao) Columns() BalanceStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BalanceStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BalanceStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BalanceStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
