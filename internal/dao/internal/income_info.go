// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IncomeInfoDao is the data access object for table income_info.
type IncomeInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns IncomeInfoColumns // columns contains all the column names of Table for convenient usage.
}

// IncomeInfoColumns defines and stores column names for table income_info.
type IncomeInfoColumns struct {
	Id                         string // id
	UserId                     string // 用户id
	TotalIncome                string // 总收益 单位分
	BalancesIncome             string // 收益余额 单位分
	TotalCommissionDiamonds    string //
	BalancesCommissionDiamonds string //
	CreateTime                 string // 总佣金收入 扩大10000
	UpdateTime                 string // 佣金收入 扩大10000
}

// incomeInfoColumns holds the columns for table income_info.
var incomeInfoColumns = IncomeInfoColumns{
	Id:                         "id",
	UserId:                     "user_id",
	TotalIncome:                "total_income",
	BalancesIncome:             "balances_income",
	TotalCommissionDiamonds:    "total_commission_diamonds",
	BalancesCommissionDiamonds: "balances_commission_diamonds",
	CreateTime:                 "create_time",
	UpdateTime:                 "update_time",
}

// NewIncomeInfoDao creates and returns a new DAO object for table data access.
func NewIncomeInfoDao() *IncomeInfoDao {
	return &IncomeInfoDao{
		group:   "default",
		table:   "income_info",
		columns: incomeInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *IncomeInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *IncomeInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *IncomeInfoDao) Columns() IncomeInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *IncomeInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *IncomeInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *IncomeInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
