// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceStatisticsUserDao is the data access object for table balance_statistics_user.
type BalanceStatisticsUserDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns BalanceStatisticsUserColumns // columns contains all the column names of Table for convenient usage.
}

// BalanceStatisticsUserColumns defines and stores column names for table balance_statistics_user.
type BalanceStatisticsUserColumns struct {
	Id            string //
	UserId        string //
	Date          string // 日期
	DepletionType string // 消费类型
	Type          string // 消费类型 1支出 2收入
	Diamonds      string // 钻石变动数额
	CreatedAt     string //
	UpdatedAt     string //
}

// balanceStatisticsUserColumns holds the columns for table balance_statistics_user.
var balanceStatisticsUserColumns = BalanceStatisticsUserColumns{
	Id:            "id",
	UserId:        "user_id",
	Date:          "date",
	DepletionType: "depletion_type",
	Type:          "type",
	Diamonds:      "diamonds",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewBalanceStatisticsUserDao creates and returns a new DAO object for table data access.
func NewBalanceStatisticsUserDao() *BalanceStatisticsUserDao {
	return &BalanceStatisticsUserDao{
		group:   "default",
		table:   "balance_statistics_user",
		columns: balanceStatisticsUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BalanceStatisticsUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BalanceStatisticsUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BalanceStatisticsUserDao) Columns() BalanceStatisticsUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BalanceStatisticsUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BalanceStatisticsUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BalanceStatisticsUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
