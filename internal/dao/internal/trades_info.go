// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TradesInfoDao is the data access object for table trades_info.
type TradesInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TradesInfoColumns // columns contains all the column names of Table for convenient usage.
}

// TradesInfoColumns defines and stores column names for table trades_info.
type TradesInfoColumns struct {
	OrderId    string // 订单Id
	Tid        string // 票据Id
	PayType    string // 支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay
	TotalFee   string // 支付费用，三方支付通知有，没有则填充订单的对应货币金额 单位分
	Currency   string // 实际支付货币，三方通知有说明，比如：IDR，没有则填充对应货币单位
	State      string // 1, 交易成功0, 未知的交易状态 2,   退款 3,  续订  4, 等待支付
	CreateTime string //
	UpdateTime string //
}

// tradesInfoColumns holds the columns for table trades_info.
var tradesInfoColumns = TradesInfoColumns{
	OrderId:    "orderId",
	Tid:        "tid",
	PayType:    "pay_type",
	TotalFee:   "total_fee",
	Currency:   "currency",
	State:      "state",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewTradesInfoDao creates and returns a new DAO object for table data access.
func NewTradesInfoDao() *TradesInfoDao {
	return &TradesInfoDao{
		group:   "default",
		table:   "trades_info",
		columns: tradesInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TradesInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TradesInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TradesInfoDao) Columns() TradesInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TradesInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TradesInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TradesInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
