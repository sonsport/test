// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceInfoDao is the data access object for table balance_info.
type BalanceInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns BalanceInfoColumns // columns contains all the column names of Table for convenient usage.
}

// BalanceInfoColumns defines and stores column names for table balance_info.
type BalanceInfoColumns struct {
	UserId              string // 用户Id
	TotalFee            string // 总花费、只有付费购买才累计到这里 印尼盾 单位：分
	TotalPayoutDiamonds string // 总流水，凡是用户花费的钻石都记录到这里 累计
	TotalPayDiamonds    string // 总购买钻石 只累计购买的
	RemainDiamonds      string // 剩余钻石数
	CreateTime          string //
	UpdateTime          string //
}

// balanceInfoColumns holds the columns for table balance_info.
var balanceInfoColumns = BalanceInfoColumns{
	UserId:              "user_id",
	TotalFee:            "total_fee",
	TotalPayoutDiamonds: "total_payout_diamonds",
	TotalPayDiamonds:    "total_pay_diamonds",
	RemainDiamonds:      "remain_diamonds",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
}

// NewBalanceInfoDao creates and returns a new DAO object for table data access.
func NewBalanceInfoDao() *BalanceInfoDao {
	return &BalanceInfoDao{
		group:   "default",
		table:   "balance_info",
		columns: balanceInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BalanceInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BalanceInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BalanceInfoDao) Columns() BalanceInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BalanceInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BalanceInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BalanceInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
