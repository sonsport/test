// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayoutRecordDao is the data access object for table payout_record.
type PayoutRecordDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns PayoutRecordColumns // columns contains all the column names of Table for convenient usage.
}

// PayoutRecordColumns defines and stores column names for table payout_record.
type PayoutRecordColumns struct {
	Id          string //
	PayoutId    string // 代付id
	OrderNo     string // 系统订单号
	TradeNo     string // 交易订单号
	Status      string // 状态
	FailureCode string // 失败状态
	CreateTime  string // 创建时间
	UpdateTime  string // 修改时间
}

// payoutRecordColumns holds the columns for table payout_record.
var payoutRecordColumns = PayoutRecordColumns{
	Id:          "id",
	PayoutId:    "payout_id",
	OrderNo:     "order_no",
	TradeNo:     "trade_no",
	Status:      "status",
	FailureCode: "failure_code",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewPayoutRecordDao creates and returns a new DAO object for table data access.
func NewPayoutRecordDao() *PayoutRecordDao {
	return &PayoutRecordDao{
		group:   "default",
		table:   "payout_record",
		columns: payoutRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PayoutRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PayoutRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PayoutRecordDao) Columns() PayoutRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PayoutRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PayoutRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PayoutRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
